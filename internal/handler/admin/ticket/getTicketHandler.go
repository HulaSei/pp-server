package ticket

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/model/dto"
	"github.com/perfect-panel/server/internal/module/support"
	"github.com/perfect-panel/server/internal/validation"
	"github.com/perfect-panel/server/pkg/httpx"
	"github.com/perfect-panel/server/pkg/result"
)

// GetTicketHandler documents Get ticket detail.
//
// @Summary Get ticket detail
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request query dto.GetTicketRequest false "Request parameters"
// @Success 200 {object} result.ResponseSuccessBean{data=dto.Ticket}
// @Router /v1/admin/ticket/detail [get]
func GetTicketHandler(service support.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req dto.GetTicketRequest
		if err := httpx.ShouldBind(c, &req); err != nil {
			result.ParamErrorResult(c, err)
			return
		}
		validateErr := validation.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		resp, err := service.GetTicket(ctx, &req)
		result.HttpResult(c, resp, err)
	}
}
