package system

import (
	"context"
	dto "github.com/perfect-panel/server/internal/module/platform/contract"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/platform"
	"github.com/perfect-panel/server/pkg/result"
)

var _ dto.GetNodeMultiplierResponse

// GetNodeMultiplierHandler documents Get Node Multiplier.
//
// @Summary Get Node Multiplier
// @Tags admin
// @Produce json
// @Security BearerAuth
// @Success 200 {object} result.ResponseSuccessBean{data=dto.GetNodeMultiplierResponse}
// @Router /v1/admin/system/get_node_multiplier [get]
func GetNodeMultiplierHandler(service platform.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {

		resp, err := service.GetNodeMultiplier(ctx)
		result.HttpResult(c, resp, err)
	}
}
