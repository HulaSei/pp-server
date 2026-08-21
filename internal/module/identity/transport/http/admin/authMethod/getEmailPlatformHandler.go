package authMethod

import (
	"context"
	dto "github.com/perfect-panel/server/internal/module/identity/contract"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/identity"
	"github.com/perfect-panel/server/pkg/result"
)

var _ dto.AuthPlatformResponse

// GetEmailPlatformHandler documents Get email support platform.
//
// @Summary Get email support platform
// @Tags admin
// @Produce json
// @Security BearerAuth
// @Success 200 {object} result.ResponseSuccessBean{data=dto.AuthPlatformResponse}
// @Router /v1/admin/auth-method/email_platform [get]
func GetEmailPlatformHandler(service identity.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {

		resp, err := service.GetEmailPlatform(ctx)
		result.HttpResult(c, resp, err)
	}
}
