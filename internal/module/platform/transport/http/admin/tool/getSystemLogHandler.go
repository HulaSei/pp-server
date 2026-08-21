package tool

import (
	"context"
	dto "github.com/perfect-panel/server/internal/module/platform/contract"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/platform"
	"github.com/perfect-panel/server/pkg/result"
)

var _ dto.LogResponse

// GetSystemLogHandler documents Get System Log.
//
// @Summary Get System Log
// @Tags admin
// @Produce json
// @Security BearerAuth
// @Success 200 {object} result.ResponseSuccessBean{data=dto.LogResponse}
// @Router /v1/admin/tool/log [get]
func GetSystemLogHandler(service platform.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {

		resp, err := service.GetSystemLog(ctx)
		result.HttpResult(c, resp, err)
	}
}
