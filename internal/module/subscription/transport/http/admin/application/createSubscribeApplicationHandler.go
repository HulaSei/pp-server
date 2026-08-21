package application

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/subscription"
	dto "github.com/perfect-panel/server/internal/module/subscription/contract"
	"github.com/perfect-panel/server/internal/validation"
	"github.com/perfect-panel/server/pkg/httpx"
	"github.com/perfect-panel/server/pkg/result"
)

// CreateSubscribeApplicationHandler documents Create subscribe application.
//
// @Summary Create subscribe application
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateSubscribeApplicationRequest true "Request parameters"
// @Success 200 {object} result.ResponseSuccessBean{data=dto.SubscribeApplication}
// @Router /v1/admin/application/ [post]
func CreateSubscribeApplicationHandler(service subscription.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req dto.CreateSubscribeApplicationRequest
		if err := httpx.ShouldBind(c, &req); err != nil {
			result.ParamErrorResult(c, err)
			return
		}
		validateErr := validation.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		resp, err := service.CreateSubscribeApplication(ctx, &req)
		result.HttpResult(c, resp, err)
	}
}
