package system

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/platform"
	"github.com/perfect-panel/server/pkg/result"
)

// GetInviteConfigHandler documents Get invite config.
//
// @Summary Get invite config
// @Tags admin
// @Produce json
// @Security BearerAuth
// @Success 200 {object} result.ResponseSuccessBean{data=dto.InviteConfig}
// @Router /v1/admin/system/invite_config [get]
func GetInviteConfigHandler(service platform.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {

		resp, err := service.GetInviteConfig(ctx)
		result.HttpResult(c, resp, err)
	}
}
