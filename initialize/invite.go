package initialize

import (
	"context"

	"github.com/perfect-panel/server/internal/config"
	"github.com/perfect-panel/server/pkg/logger"
)

func Invite(ctx *Dependencies) {
	// Initialize the system configuration
	logger.Debug("Register config initialization")
	configs, err := ctx.Store.System().GetInviteConfig(context.Background())
	if err != nil {
		logger.Error("[Init Invite Config] Get Invite Config Error: ", logger.Field("error", err.Error()))
		return
	}
	var inviteConfig config.InviteConfig
	config.SystemConfigSliceReflectToStruct(configs, &inviteConfig)
	ctx.updateConfig(func(current *config.Config) { current.Invite = inviteConfig })
}
