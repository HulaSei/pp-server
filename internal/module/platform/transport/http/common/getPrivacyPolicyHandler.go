package common

import (
	"context"
	dto "github.com/perfect-panel/server/internal/module/platform/contract"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/platform"
	"github.com/perfect-panel/server/pkg/result"
)

var _ dto.PrivacyPolicyConfig

// GetPrivacyPolicyHandler documents Get Privacy Policy.
//
// @Summary Get Privacy Policy
// @Tags common
// @Produce json
// @Success 200 {object} result.ResponseSuccessBean{data=dto.PrivacyPolicyConfig}
// @Router /v1/common/site/privacy [get]
func GetPrivacyPolicyHandler(service platform.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {

		resp, err := service.GetPrivacyPolicy(ctx)
		result.HttpResult(c, resp, err)
	}
}
