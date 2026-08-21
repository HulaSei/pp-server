package system

import (
	"context"
	dto "github.com/perfect-panel/server/internal/module/platform/contract"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/platform"
	"github.com/perfect-panel/server/pkg/result"
)

var _ dto.VerifyCodeConfig

// GetVerifyCodeConfigHandler documents Get Verify Code Config.
//
// @Summary Get Verify Code Config
// @Tags admin
// @Produce json
// @Security BearerAuth
// @Success 200 {object} result.ResponseSuccessBean{data=dto.VerifyCodeConfig}
// @Router /v1/admin/system/verify_code_config [get]
func GetVerifyCodeConfigHandler(service platform.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {

		resp, err := service.GetVerifyCodeConfig(ctx)
		result.HttpResult(c, resp, err)
	}
}
