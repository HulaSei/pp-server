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

// CloseOrderHandler documents Close order.
//
// @Summary Close order
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CloseOrderRequest true "Request parameters"
// @Success 200 {object} result.ResponseSuccessBean
// @Router /v1/public/order/close [post]
func CloseOrderHandler(service billing.Service) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		var req dto.CloseOrderRequest
		if err := httpx.ShouldBind(ctx, &req); err != nil {
			result.ParamErrorResult(ctx, err)
			return
		}
		validateErr := validation.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(ctx, validateErr)
			return
		}

		err := service.CloseOrder(c, &req)
		result.HttpResult(ctx, nil, err)
	}
}
