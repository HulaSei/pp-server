package user

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/subscription"
	dto "github.com/perfect-panel/server/internal/module/subscription/contract"
	"github.com/perfect-panel/server/internal/validation"
	"github.com/perfect-panel/server/pkg/httpx"
	"github.com/perfect-panel/server/pkg/result"
)

// GetSubscribeLogHandler documents Get Subscribe Log.
//
// @Summary Get Subscribe Log
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request query dto.GetSubscribeLogRequest false "Request parameters"
// @Success 200 {object} result.ResponseSuccessBean{data=dto.GetSubscribeLogResponse}
// @Router /v1/public/user/subscribe_log [get]
func GetSubscribeLogHandler(service subscription.Service) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		var req dto.GetSubscribeLogRequest
		if err := httpx.ShouldBind(ctx, &req); err != nil {
			result.ParamErrorResult(ctx, err)
			return
		}
		validateErr := validation.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(ctx, validateErr)
			return
		}

		resp, err := service.GetSubscribeLog(c, &req)
		result.HttpResult(ctx, resp, err)
	}
}
