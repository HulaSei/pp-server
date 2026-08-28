package svc

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/oschwald/geoip2-golang"
	"github.com/perfect-panel/server/pkg/logger"
)

const GeoIPDBURL = "https://raw.githubusercontent.com/adysec/IP_database/main/geolite/GeoLite2-City.mmdb"

type IPLocation struct {
	Path string
	DB   *geoip2.Reader
}

func NewIPLocation(path string) (*IPLocation, error) {

	// 检查文件是否存在
	if _, err := os.Stat(path); os.IsNotExist(err) {
		logger.Infof("[GeoIP] Database not found, downloading from %s", GeoIPDBURL)
		// 文件不存在，下载数据库
		err := DownloadGeoIPDatabase(GeoIPDBURL, path)
		if err != nil {
			logger.Errorf("[GeoIP] Failed to download database: %v", err.Error())
			return nil, err
		}
		logger.Infof("[GeoIP] Database downloaded successfully")
	}

	db, err := geoip2.Open(path)
	if err != nil {
		return nil, err
	}
	return &IPLocation{
		Path: path,
		DB:   db,
	}, nil
}

func (ipLoc *IPLocation) Close() error {
	return ipLoc.DB.Close()
}

func DownloadGeoIPDatabase(url, path string) error {

	// 创建路径, 确保目录存在
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		logger.Errorf("[GeoIP] Failed to create directory: %v", err.Error())
		return err
	}

	// Write into a sibling temporary file so a timeout or truncated response
	// never leaves a corrupt database that blocks the next startup retry.
	out, err := os.CreateTemp(filepath.Dir(path), ".geoip-*.tmp")
	if err != nil {
		return err
	}
	tempPath := out.Name()
	committed := false
	defer func() {
		_ = out.Close()
		if !committed {
			_ = os.Remove(tempPath)
		}
	}()

	// 请求远程文件
	client := &http.Client{Timeout: 2 * time.Minute}
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return fmt.Errorf("download GeoIP database: HTTP %d", resp.StatusCode)
	}

	// 保存文件
	const maxGeoIPDatabaseSize = int64(256 << 20)
	written, err := io.Copy(out, io.LimitReader(resp.Body, maxGeoIPDatabaseSize+1))
	if err != nil {
		return err
	}
	if written > maxGeoIPDatabaseSize {
		return fmt.Errorf("download GeoIP database: response exceeds %d bytes", maxGeoIPDatabaseSize)
	}
	if err := out.Sync(); err != nil {
		return err
	}
	if err := out.Close(); err != nil {
		return err
	}
	if err := os.Rename(tempPath, path); err != nil {
		return err
	}
	committed = true
	return nil
}
