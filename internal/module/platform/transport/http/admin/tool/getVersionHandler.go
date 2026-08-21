package tool

import (
	"context"
	dto "github.com/perfect-panel/server/internal/module/platform/contract"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/platform"
	"github.com/perfect-panel/server/pkg/result"
)

var _ dto.VersionResponse

// GetVersionHandler documents Get Version.
//
// @Summary Get Version
// @Tags admin
// @Produce json
// @Security BearerAuth
// @Success 200 {object} result.ResponseSuccessBean{data=dto.VersionResponse}
// @Router /v1/admin/tool/version [get]
func GetVersionHandler(service platform.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {

		resp, err := service.GetVersion(ctx)
		result.HttpResult(c, resp, err)
	}
}
