package initialize

import (
	"context"

	"github.com/perfect-panel/server/internal/config"
	"github.com/perfect-panel/server/pkg/logger"
)

func Subscribe(svc *Dependencies) {
	logger.Debug("Subscribe config initialization")
	configs, err := svc.Store.System().GetSubscribeConfig(context.Background())
	if err != nil {
		logger.Error("[Init Subscribe Config] Get Subscribe Config Error: ", logger.Field("error", err.Error()))
		return
	}

	var subscribeConfig config.SubscribeConfig
	config.SystemConfigSliceReflectToStruct(configs, &subscribeConfig)
	svc.updateConfig(func(current *config.Config) { current.Subscribe = subscribeConfig })
}
