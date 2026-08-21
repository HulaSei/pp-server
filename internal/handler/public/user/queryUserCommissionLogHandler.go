package user

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/model/dto"
	"github.com/perfect-panel/server/internal/module/billing"
	"github.com/perfect-panel/server/internal/validation"
	"github.com/perfect-panel/server/pkg/httpx"
	"github.com/perfect-panel/server/pkg/result"
)

// QueryUserCommissionLogHandler documents Query User Commission Log.
//
// @Summary Query User Commission Log
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request query dto.QueryUserCommissionLogListRequest false "Request parameters"
// @Success 200 {object} result.ResponseSuccessBean{data=dto.QueryUserCommissionLogListResponse}
// @Router /v1/public/user/commission_log [get]
func QueryUserCommissionLogHandler(service billing.Service) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		var req dto.QueryUserCommissionLogListRequest
		if err := httpx.ShouldBind(ctx, &req); err != nil {
			result.ParamErrorResult(ctx, err)
			return
		}
		validateErr := validation.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(ctx, validateErr)
			return
		}

		resp, err := service.QueryUserCommissionLog(c, &req)
		result.HttpResult(ctx, resp, err)
	}
}
