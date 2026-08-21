package subscribe

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/subscription"
	dto "github.com/perfect-panel/server/internal/module/subscription/contract"
	"github.com/perfect-panel/server/internal/validation"
	"github.com/perfect-panel/server/pkg/httpx"
	"github.com/perfect-panel/server/pkg/result"
)

// BatchDeleteSubscribeHandler documents Batch delete subscribe.
//
// @Summary Batch delete subscribe
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.BatchDeleteSubscribeRequest true "Request parameters"
// @Success 200 {object} result.ResponseSuccessBean
// @Router /v1/admin/subscribe/batch [delete]
func BatchDeleteSubscribeHandler(service subscription.Service) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		var req dto.BatchDeleteSubscribeRequest
		if err := httpx.ShouldBind(ctx, &req); err != nil {
			result.ParamErrorResult(ctx, err)
			return
		}
		validateErr := validation.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(ctx, validateErr)
			return
		}

		err := service.BatchDeleteSubscribe(c, &req)
		result.HttpResult(ctx, nil, err)
	}
}
