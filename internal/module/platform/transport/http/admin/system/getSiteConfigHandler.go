package system

import (
	"context"
	dto "github.com/perfect-panel/server/internal/module/platform/contract"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/platform"
	"github.com/perfect-panel/server/pkg/result"
)

var _ dto.SiteConfig

// GetSiteConfigHandler documents Get site config.
//
// @Summary Get site config
// @Tags admin
// @Produce json
// @Security BearerAuth
// @Success 200 {object} result.ResponseSuccessBean{data=dto.SiteConfig}
// @Router /v1/admin/system/site_config [get]
func GetSiteConfigHandler(service platform.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {

		resp, err := service.GetSiteConfig(ctx)
		result.HttpResult(c, resp, err)
	}
}
