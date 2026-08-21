package authMethod

import (
	"context"
	dto "github.com/perfect-panel/server/internal/module/identity/contract"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/identity"
	"github.com/perfect-panel/server/pkg/result"
)

var _ dto.AuthPlatformResponse

// GetSmsPlatformHandler documents Get sms support platform.
//
// @Summary Get sms support platform
// @Tags admin
// @Produce json
// @Security BearerAuth
// @Success 200 {object} result.ResponseSuccessBean{data=dto.AuthPlatformResponse}
// @Router /v1/admin/auth-method/sms_platform [get]
func GetSmsPlatformHandler(service identity.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {

		resp, err := service.GetSmsPlatform(ctx)
		result.HttpResult(c, resp, err)
	}
}
