package system

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/platform"
	"github.com/perfect-panel/server/pkg/result"
)

// PreViewNodeMultiplierHandler documents PreView Node Multiplier.
//
// @Summary PreView Node Multiplier
// @Tags admin
// @Produce json
// @Security BearerAuth
// @Success 200 {object} result.ResponseSuccessBean{data=dto.PreViewNodeMultiplierResponse}
// @Router /v1/admin/system/node_multiplier/preview [get]
func PreViewNodeMultiplierHandler(service platform.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {

		resp, err := service.PreViewNodeMultiplier(ctx)
		result.HttpResult(c, resp, err)
	}
}
