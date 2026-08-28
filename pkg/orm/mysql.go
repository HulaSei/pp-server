package orm

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/perfect-panel/server/pkg/logger"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	DriverMySQL     = "mysql"
	DriverPostgres  = "postgres"
	DriverPostgres2 = "postgresql"

	DefaultMySQLConfig             = "charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
	DefaultPostgresConfig          = "sslmode=disable&TimeZone=Asia%2FShanghai&application_name=perfect-panel"
	DefaultSlowThresholdMs         = 1000
	DefaultConnMaxLifetimeSeconds  = 1800
	DefaultConnMaxIdleTimeSeconds  = 300
	defaultPostgresApplicationName = "perfect-panel"
)

type Config struct {
	Driver          string `yaml:"Driver" default:"mysql"`
	Addr            string `yaml:"Addr"`
	Username        string `yaml:"Username"`
	Password        string `yaml:"Password"`
	Dbname          string `yaml:"Dbname"`
	Config          string `yaml:"Config" default:"charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"`
	MaxIdleConns    int    `yaml:"MaxIdleConns" default:"10"`
	MaxOpenConns    int    `yaml:"MaxOpenConns" default:"10"`
	ConnMaxLifetime int64  `yaml:"ConnMaxLifetime" default:"1800"`
	ConnMaxIdleTime int64  `yaml:"ConnMaxIdleTime" default:"300"`
	SlowThreshold   int64  `yaml:"SlowThreshold" default:"1000"`
}

type Mysql struct {
	Config Config
}

func NormalizeDriver(driver string) string {
	switch strings.ToLower(strings.TrimSpace(driver)) {
	case "", DriverMySQL:
		return DriverMySQL
	case DriverPostgres, DriverPostgres2, "pgsql":
		return DriverPostgres
	default:
		return strings.ToLower(strings.TrimSpace(driver))
	}
}

func (m Mysql) Driver() string {
	return NormalizeDriver(m.Config.Driver)
}

func (m Mysql) Dsn() string {
	switch m.Driver() {
	case DriverPostgres:
		return m.postgresDsn()
	default:
		return m.mysqlDsn()
	}
}

func (m Mysql) MigrationDsn() string {
	return m.Dsn()
}

func (m Mysql) mysqlDsn() string {
	query := m.Config.Config
	if query == "" {
		query = DefaultMySQLConfig
	}
	return m.Config.Username + ":" + m.Config.Password + "@tcp(" + m.Config.Addr + ")/" + m.Config.Dbname + "?" + query
}

func (m Mysql) postgresDsn() string {
	query := m.Config.Config
	if query == "" || query == DefaultMySQLConfig {
		query = DefaultPostgresConfig
	}
	u := url.URL{
		Scheme: DriverPostgres,
		Host:   m.Config.Addr,
		Path:   "/" + m.Config.Dbname,
	}
	if m.Config.Username != "" {
		u.User = url.UserPassword(m.Config.Username, m.Config.Password)
	}
	params, err := url.ParseQuery(query)
	if err != nil {
		u.RawQuery = query
	} else {
		if params.Get("application_name") == "" {
			params.Set("application_name", defaultPostgresApplicationName)
		}
		u.RawQuery = params.Encode()
	}
	return u.String()
}

func (m *Mysql) GetSlowThreshold() time.Duration {
	return time.Duration(m.Config.SlowThreshold) * time.Millisecond
}
func (m *Mysql) GetColorful() bool {
	return true
}

func ConnectMysql(m Mysql) (*gorm.DB, error) {
	return ConnectDatabase(m)
}

func ConnectDatabase(m Mysql) (*gorm.DB, error) {
	if m.Config.Dbname == "" {
		return nil, errors.New("database name is empty")
	}
	var dialector gorm.Dialector
	switch m.Driver() {
	case DriverMySQL:
		dialector = mysql.New(mysql.Config{DSN: m.Dsn()})
	case DriverPostgres:
		dialector = postgres.Open(m.Dsn())
	default:
		return nil, fmt.Errorf("unsupported database driver: %s", m.Config.Driver)
	}
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: &logger.GormLogger{SlowThreshold: m.GetSlowThreshold()},
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}
	sqldb, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqldb.SetMaxIdleConns(m.Config.MaxIdleConns)
	sqldb.SetMaxOpenConns(m.Config.MaxOpenConns)
	if m.Config.ConnMaxLifetime > 0 {
		sqldb.SetConnMaxLifetime(time.Duration(m.Config.ConnMaxLifetime) * time.Second)
	}
	if m.Config.ConnMaxIdleTime > 0 {
		sqldb.SetConnMaxIdleTime(time.Duration(m.Config.ConnMaxIdleTime) * time.Second)
	}
	return db, nil
}

func Ping(dsn string) bool {
	return PingDatabase(DriverMySQL, dsn)
}

func PingDatabase(driver, dsn string) bool {
	var dialector gorm.Dialector
	switch NormalizeDriver(driver) {
	case DriverMySQL:
		dialector = mysql.Open(dsn)
	case DriverPostgres:
		dialector = postgres.Open(dsn)
	default:
		fmt.Printf("unsupported database driver: %s\n", driver)
		return false
	}
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		fmt.Printf("connect database failed, err: %v\n", err.Error())
		return false
	}
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Printf("get database connection failed, err: %v\n", err)
		return false
	}
	defer sqlDB.Close()
	return sqlDB.Ping() == nil
}
