package payment

import (
	"context"
	dto "github.com/perfect-panel/server/internal/module/billing/contract"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/billing"
	"github.com/perfect-panel/server/pkg/result"
)

var _ dto.GetAvailablePaymentMethodsResponse

// GetAvailablePaymentMethodsHandler documents Get available payment methods.
//
// @Summary Get available payment methods
// @Tags user
// @Produce json
// @Security BearerAuth
// @Success 200 {object} result.ResponseSuccessBean{data=dto.GetAvailablePaymentMethodsResponse}
// @Router /v1/public/payment/methods [get]
func GetAvailablePaymentMethodsHandler(service billing.Service) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {

		resp, err := service.GetAvailablePaymentMethods(c)
		result.HttpResult(ctx, resp, err)
	}
}
