package user

import (
	"context"
	dto "github.com/perfect-panel/server/internal/module/billing/contract"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/billing"
	"github.com/perfect-panel/server/pkg/result"
)

var _ dto.QueryUserBalanceLogListResponse

// QueryUserBalanceLogHandler documents Query User Balance Log.
//
// @Summary Query User Balance Log
// @Tags user
// @Produce json
// @Security BearerAuth
// @Success 200 {object} result.ResponseSuccessBean{data=dto.QueryUserBalanceLogListResponse}
// @Router /v1/public/user/balance_log [get]
func QueryUserBalanceLogHandler(service billing.Service) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {

		resp, err := service.QueryUserBalanceLog(c)
		result.HttpResult(ctx, resp, err)
	}
}
