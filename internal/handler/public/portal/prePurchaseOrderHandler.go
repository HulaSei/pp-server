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

// PrePurchaseOrderHandler documents Pre Purchase Order.
//
// @Summary Pre Purchase Order
// @Tags user
// @Accept json
// @Produce json
// @Param request body dto.PrePurchaseOrderRequest true "Request parameters"
// @Success 200 {object} result.ResponseSuccessBean{data=dto.PrePurchaseOrderResponse}
// @Router /v1/public/portal/pre [post]
func PrePurchaseOrderHandler(service billing.Service) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		var req dto.PrePurchaseOrderRequest
		if err := httpx.ShouldBind(ctx, &req); err != nil {
			result.ParamErrorResult(ctx, err)
			return
		}
		validateErr := validation.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(ctx, validateErr)
			return
		}

		resp, err := service.PortalPrePurchase(c, &req)
		result.HttpResult(ctx, resp, err)
	}
}
