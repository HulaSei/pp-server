package ticket

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/support"
	dto "github.com/perfect-panel/server/internal/module/support/contract"
	"github.com/perfect-panel/server/internal/validation"
	"github.com/perfect-panel/server/pkg/httpx"
	"github.com/perfect-panel/server/pkg/result"
)

// GetUserTicketListHandler documents Get ticket list.
//
// @Summary Get ticket list
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request query dto.GetUserTicketListRequest false "Request parameters"
// @Success 200 {object} result.ResponseSuccessBean{data=dto.GetUserTicketListResponse}
// @Router /v1/public/ticket/list [get]
func GetUserTicketListHandler(service support.Service) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		var req dto.GetUserTicketListRequest
		if err := httpx.ShouldBind(ctx, &req); err != nil {
			result.ParamErrorResult(ctx, err)
			return
		}
		validateErr := validation.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(ctx, validateErr)
			return
		}

		resp, err := service.GetUserTicketList(c, &req)
		result.HttpResult(ctx, resp, err)
	}
}
