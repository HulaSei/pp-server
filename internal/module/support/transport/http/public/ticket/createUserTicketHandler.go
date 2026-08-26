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

// CreateUserTicketHandler documents Create ticket.
//
// @Summary Create ticket
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateUserTicketRequest true "Request parameters"
// @Success 200 {object} result.ResponseSuccessBean
// @Router /v1/public/ticket/ [post]
func CreateUserTicketHandler(service support.Service) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		var req dto.CreateUserTicketRequest
		if err := httpx.ShouldBind(ctx, &req); err != nil {
			result.ParamErrorResult(ctx, err)
			return
		}
		validateErr := validation.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(ctx, validateErr)
			return
		}

		err := service.CreateUserTicket(c, &req)
		result.HttpResult(ctx, nil, err)
	}
}
