package marketing

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/model/dto"
	"github.com/perfect-panel/server/internal/module/support"
	"github.com/perfect-panel/server/internal/validation"
	"github.com/perfect-panel/server/pkg/httpx"
	"github.com/perfect-panel/server/pkg/result"
)

// QueryQuotaTaskListHandler documents Query quota task list.
//
// @Summary Query quota task list
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request query dto.QueryQuotaTaskListRequest false "Request parameters"
// @Success 200 {object} result.ResponseSuccessBean{data=dto.QueryQuotaTaskListResponse}
// @Router /v1/admin/marketing/quota/list [get]
func QueryQuotaTaskListHandler(service support.Service) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		var req dto.QueryQuotaTaskListRequest
		if err := httpx.ShouldBind(ctx, &req); err != nil {
			result.ParamErrorResult(ctx, err)
			return
		}
		validateErr := validation.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(ctx, validateErr)
			return
		}

		resp, err := service.QueryQuotaTaskList(c, &req)
		result.HttpResult(ctx, resp, err)
	}
}
