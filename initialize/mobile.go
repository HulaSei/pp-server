package initialize

import (
	"context"
	"encoding/json"

	"github.com/perfect-panel/server/pkg/logger"

	"github.com/perfect-panel/server/internal/config"
	"github.com/perfect-panel/server/internal/module/identity/entity/auth"
	"github.com/perfect-panel/server/pkg/tool"
)

func Mobile(ctx *Dependencies) {
	logger.Debug("Mobile config initialization")
	method, err := ctx.Store.Auth().FindOneByMethod(context.Background(), "mobile")
	if err != nil {
		panic(err)
	}
	var cfg config.MobileConfig
	var mobileConfig auth.MobileAuthConfig
	mobileConfig.Unmarshal(method.Config)
	tool.DeepCopy(&cfg, mobileConfig)
	cfg.Enable = *method.Enabled
	value, _ := json.Marshal(mobileConfig.PlatformConfig)
	cfg.PlatformConfig = string(value)
	ctx.updateConfig(func(current *config.Config) { current.Mobile = cfg })
}
