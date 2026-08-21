package system

import (
	"context"
	dto "github.com/perfect-panel/server/internal/module/platform/contract"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/platform"
	"github.com/perfect-panel/server/pkg/result"
)

var _ dto.RegisterConfig

// GetRegisterConfigHandler documents Get register config.
//
// @Summary Get register config
// @Tags admin
// @Produce json
// @Security BearerAuth
// @Success 200 {object} result.ResponseSuccessBean{data=dto.RegisterConfig}
// @Router /v1/admin/system/register_config [get]
func GetRegisterConfigHandler(service platform.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {

		resp, err := service.GetRegisterConfig(ctx)
		result.HttpResult(c, resp, err)
	}
}
