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

// SubscribeSortHandler documents Subscribe sort.
//
// @Summary Subscribe sort
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.SubscribeSortRequest true "Request parameters"
// @Success 200 {object} result.ResponseSuccessBean
// @Router /v1/admin/subscribe/sort [post]
func SubscribeSortHandler(service subscription.Service) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		var req dto.SubscribeSortRequest
		if err := httpx.ShouldBind(ctx, &req); err != nil {
			result.ParamErrorResult(ctx, err)
			return
		}
		validateErr := validation.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(ctx, validateErr)
			return
		}

		err := service.SubscribeSort(c, &req)
		result.HttpResult(ctx, nil, err)
	}
}
