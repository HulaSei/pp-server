package system

import (
	"context"
	dto "github.com/perfect-panel/server/internal/module/platform/contract"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/platform"
	"github.com/perfect-panel/server/pkg/result"
)

var _ dto.SubscribeConfig

// GetSubscribeConfigHandler documents Get subscribe config.
//
// @Summary Get subscribe config
// @Tags admin
// @Produce json
// @Security BearerAuth
// @Success 200 {object} result.ResponseSuccessBean{data=dto.SubscribeConfig}
// @Router /v1/admin/system/subscribe_config [get]
func GetSubscribeConfigHandler(service platform.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {

		resp, err := service.GetSubscribeConfig(ctx)
		result.HttpResult(c, resp, err)
	}
}
