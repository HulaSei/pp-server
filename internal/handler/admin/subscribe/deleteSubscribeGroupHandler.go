package subscribe

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/model/dto"
	"github.com/perfect-panel/server/internal/module/subscription"
	"github.com/perfect-panel/server/internal/validation"
	"github.com/perfect-panel/server/pkg/httpx"
	"github.com/perfect-panel/server/pkg/result"
)

// DeleteSubscribeGroupHandler documents Delete subscribe group.
//
// @Summary Delete subscribe group
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.DeleteSubscribeGroupRequest true "Request parameters"
// @Success 200 {object} result.ResponseSuccessBean
// @Router /v1/admin/subscribe/group [delete]
func DeleteSubscribeGroupHandler(service subscription.Service) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		var req dto.DeleteSubscribeGroupRequest
		if err := httpx.ShouldBind(ctx, &req); err != nil {
			result.ParamErrorResult(ctx, err)
			return
		}
		validateErr := validation.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(ctx, validateErr)
			return
		}

		err := service.DeleteSubscribeGroup(c, &req)
		result.HttpResult(ctx, nil, err)
	}
}
