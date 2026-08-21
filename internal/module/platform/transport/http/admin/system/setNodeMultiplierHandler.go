package system

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/platform"
	dto "github.com/perfect-panel/server/internal/module/platform/contract"
	"github.com/perfect-panel/server/internal/validation"
	"github.com/perfect-panel/server/pkg/httpx"
	"github.com/perfect-panel/server/pkg/result"
)

// SetNodeMultiplierHandler documents Set Node Multiplier.
//
// @Summary Set Node Multiplier
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.SetNodeMultiplierRequest true "Request parameters"
// @Success 200 {object} result.ResponseSuccessBean
// @Router /v1/admin/system/set_node_multiplier [post]
func SetNodeMultiplierHandler(service platform.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req dto.SetNodeMultiplierRequest
		if err := httpx.ShouldBind(c, &req); err != nil {
			result.ParamErrorResult(c, err)
			return
		}
		validateErr := validation.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		err := service.SetNodeMultiplier(ctx, &req)
		result.HttpResult(c, nil, err)
	}
}
