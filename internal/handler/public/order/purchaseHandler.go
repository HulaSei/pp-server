package order

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/model/dto"
	"github.com/perfect-panel/server/internal/module/billing"
	"github.com/perfect-panel/server/internal/validation"
	"github.com/perfect-panel/server/pkg/httpx"
	"github.com/perfect-panel/server/pkg/result"
)

// PurchaseHandler documents purchase Subscription.
//
// @Summary purchase Subscription
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.PurchaseOrderRequest true "Request parameters"
// @Success 200 {object} result.ResponseSuccessBean{data=dto.PurchaseOrderResponse}
// @Router /v1/public/order/purchase [post]
func PurchaseHandler(service billing.Service) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		var req dto.PurchaseOrderRequest
		if err := httpx.ShouldBind(ctx, &req); err != nil {
			result.ParamErrorResult(ctx, err)
			return
		}
		validateErr := validation.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(ctx, validateErr)
			return
		}

		resp, err := service.Purchase(c, &req)
		result.HttpResult(ctx, resp, err)
	}
}
