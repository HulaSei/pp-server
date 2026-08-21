package log

import (
	"context"
	dto "github.com/perfect-panel/server/internal/module/platform/contract"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/platform"
	"github.com/perfect-panel/server/pkg/result"
)

var _ dto.LogSetting

// GetLogSettingHandler documents Get log setting.
//
// @Summary Get log setting
// @Tags admin
// @Produce json
// @Security BearerAuth
// @Success 200 {object} result.ResponseSuccessBean{data=dto.LogSetting}
// @Router /v1/admin/log/setting [get]
func GetLogSettingHandler(service platform.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {

		resp, err := service.GetLogSetting(ctx)
		result.HttpResult(c, resp, err)
	}
}
