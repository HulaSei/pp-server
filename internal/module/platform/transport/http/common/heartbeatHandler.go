package common

import (
	"context"
	dto "github.com/perfect-panel/server/internal/module/platform/contract"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/platform"
	"github.com/perfect-panel/server/pkg/result"
)

var _ dto.HeartbeatResponse

// HeartbeatHandler documents Heartbeat.
//
// @Summary Heartbeat
// @Tags common
// @Produce json
// @Success 200 {object} result.ResponseSuccessBean{data=dto.HeartbeatResponse}
// @Router /v1/common/heartbeat [get]
func HeartbeatHandler(service platform.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {

		resp, err := service.Heartbeat(ctx)
		result.HttpResult(c, resp, err)
	}
}
