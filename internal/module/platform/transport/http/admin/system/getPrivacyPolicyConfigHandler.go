package system

import (
	"context"
	dto "github.com/perfect-panel/server/internal/module/platform/contract"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/platform"
	"github.com/perfect-panel/server/pkg/result"
)

var _ dto.PrivacyPolicyConfig

// GetPrivacyPolicyConfigHandler documents get Privacy Policy Config.
//
// @Summary get Privacy Policy Config
// @Tags admin
// @Produce json
// @Security BearerAuth
// @Success 200 {object} result.ResponseSuccessBean{data=dto.PrivacyPolicyConfig}
// @Router /v1/admin/system/privacy [get]
func GetPrivacyPolicyConfigHandler(service platform.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {

		resp, err := service.GetPrivacyPolicyConfig(ctx)
		result.HttpResult(c, resp, err)
	}
}
