package payment

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/billing"
	dto "github.com/perfect-panel/server/internal/module/billing/contract"
	"github.com/perfect-panel/server/internal/validation"
	"github.com/perfect-panel/server/pkg/httpx"
	"github.com/perfect-panel/server/pkg/result"
)

// UpdatePaymentMethodHandler documents Update Payment Method.
//
// @Summary Update Payment Method
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.UpdatePaymentMethodRequest true "Request parameters"
// @Success 200 {object} result.ResponseSuccessBean{data=dto.PaymentConfig}
// @Router /v1/admin/payment/ [put]
func UpdatePaymentMethodHandler(service billing.Service) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		var req dto.UpdatePaymentMethodRequest
		if err := httpx.ShouldBind(ctx, &req); err != nil {
			result.ParamErrorResult(ctx, err)
			return
		}
		validateErr := validation.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(ctx, validateErr)
			return
		}

		resp, err := service.UpdatePaymentMethod(c, &req)
		result.HttpResult(ctx, resp, err)
	}
}
