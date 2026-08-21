package common

import (
	"context"
	dto "github.com/perfect-panel/server/internal/module/platform/contract"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/platform"
	"github.com/perfect-panel/server/pkg/result"
)

var _ dto.GetGlobalConfigResponse

// GetGlobalConfigHandler documents Get global config.
//
// @Summary Get global config
// @Tags common
// @Produce json
// @Success 200 {object} result.ResponseSuccessBean{data=dto.GetGlobalConfigResponse}
// @Router /v1/common/site/config [get]
func GetGlobalConfigHandler(service platform.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		resp, err := service.GetGlobalConfig(ctx)
		result.HttpResult(c, resp, err)
	}
}
