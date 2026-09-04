package initialize

import (
	"context"

	"github.com/perfect-panel/server/internal/config"
	"github.com/perfect-panel/server/pkg/logger"
)

func Site(ctx *Dependencies) {
	logger.Debug("initialize site config")
	configs, err := ctx.Store.System().GetSiteConfig(context.Background())
	if err != nil {
		panic(err)
	}
	var siteConfig config.SiteConfig
	config.SystemConfigSliceReflectToStruct(configs, &siteConfig)
	ctx.updateConfig(func(current *config.Config) { current.Site = siteConfig })
}
