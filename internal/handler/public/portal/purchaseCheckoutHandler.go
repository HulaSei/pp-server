package portal

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/model/dto"
	"github.com/perfect-panel/server/internal/module/billing"
	"github.com/perfect-panel/server/internal/validation"
	"github.com/perfect-panel/server/pkg/httpx"
	"github.com/perfect-panel/server/pkg/result"
)

// PurchaseCheckoutHandler documents Purchase Checkout.
//
// @Summary Purchase Checkout
// @Tags user
// @Accept json
// @Produce json
// @Param request body dto.CheckoutOrderRequest true "Request parameters"
// @Success 200 {object} result.ResponseSuccessBean{data=dto.CheckoutOrderResponse}
// @Router /v1/public/portal/order/checkout [post]
func PurchaseCheckoutHandler(service billing.Service) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		var req dto.CheckoutOrderRequest
		if err := httpx.ShouldBind(ctx, &req); err != nil {
			result.ParamErrorResult(ctx, err)
			return
		}
		validateErr := validation.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(ctx, validateErr)
			return
		}

		resp, err := service.PortalCheckout(c, &req)
		result.HttpResult(ctx, resp, err)
	}
}
