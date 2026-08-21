package system

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/platform"
	"github.com/perfect-panel/server/pkg/result"
)

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
