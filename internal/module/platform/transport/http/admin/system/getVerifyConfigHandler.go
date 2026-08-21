package system

import (
	"context"
	dto "github.com/perfect-panel/server/internal/module/platform/contract"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/platform"
	"github.com/perfect-panel/server/pkg/result"
)

var _ dto.VerifyConfig

// GetVerifyConfigHandler documents Get verify config.
//
// @Summary Get verify config
// @Tags admin
// @Produce json
// @Security BearerAuth
// @Success 200 {object} result.ResponseSuccessBean{data=dto.VerifyConfig}
// @Router /v1/admin/system/verify_config [get]
func GetVerifyConfigHandler(service platform.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {

		resp, err := service.GetVerifyConfig(ctx)
		result.HttpResult(c, resp, err)
	}
}
