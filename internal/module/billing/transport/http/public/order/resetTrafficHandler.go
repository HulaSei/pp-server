package order

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/billing"
	dto "github.com/perfect-panel/server/internal/module/billing/contract"
	"github.com/perfect-panel/server/internal/validation"
	"github.com/perfect-panel/server/pkg/httpx"
	"github.com/perfect-panel/server/pkg/result"
)

// ResetTrafficHandler documents Reset traffic.
//
// @Summary Reset traffic
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.ResetTrafficOrderRequest true "Request parameters"
// @Success 200 {object} result.ResponseSuccessBean{data=dto.ResetTrafficOrderResponse}
// @Router /v1/public/order/reset [post]
func ResetTrafficHandler(service billing.Service) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		var req dto.ResetTrafficOrderRequest
		if err := httpx.ShouldBind(ctx, &req); err != nil {
			result.ParamErrorResult(ctx, err)
			return
		}
		validateErr := validation.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(ctx, validateErr)
			return
		}

		resp, err := service.ResetTraffic(c, &req)
		result.HttpResult(ctx, resp, err)
	}
}
