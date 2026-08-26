package system

import (
	"context"
	dto "github.com/perfect-panel/server/internal/module/platform/contract"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/platform"
	"github.com/perfect-panel/server/pkg/result"
)

var _ dto.ModuleConfig

// GetModuleConfigHandler documents Get Module Config.
//
// @Summary Get Module Config
// @Tags admin
// @Produce json
// @Security BearerAuth
// @Success 200 {object} result.ResponseSuccessBean{data=dto.ModuleConfig}
// @Router /v1/admin/system/module [get]
func GetModuleConfigHandler(service platform.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {

		resp, err := service.GetModuleConfig(ctx)
		result.HttpResult(c, resp, err)
	}
}
