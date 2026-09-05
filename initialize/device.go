package initialize

import (
	"context"

	"github.com/perfect-panel/server/internal/config"
	"github.com/perfect-panel/server/internal/infra/mapping"
	"github.com/perfect-panel/server/internal/module/identity/entity/auth"
	"github.com/perfect-panel/server/pkg/logger"
)

func Device(ctx *Dependencies) {
	logger.Debug("device config initialization")
	method, err := ctx.Store.Auth().FindOneByMethod(context.Background(), "device")
	if err != nil {
		panic(err)
	}
	var cfg config.DeviceConfig
	var deviceConfig auth.DeviceConfig
	deviceConfig.Unmarshal(method.Config)
	mapping.DeepCopy(&cfg, deviceConfig)
	cfg.Enable = *method.Enabled
	ctx.updateConfig(func(current *config.Config) { current.Device = cfg })
}
