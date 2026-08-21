package application

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/model/dto"
	"github.com/perfect-panel/server/internal/module/subscription"
	"github.com/perfect-panel/server/internal/validation"
	"github.com/perfect-panel/server/pkg/httpx"
	"github.com/perfect-panel/server/pkg/result"
)

// UpdateSubscribeApplicationHandler documents Update subscribe application.
//
// @Summary Update subscribe application
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.UpdateSubscribeApplicationRequest true "Request parameters"
// @Success 200 {object} result.ResponseSuccessBean{data=dto.SubscribeApplication}
// @Router /v1/admin/application/subscribe_application [put]
func UpdateSubscribeApplicationHandler(service subscription.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req dto.UpdateSubscribeApplicationRequest
		if err := httpx.ShouldBind(c, &req); err != nil {
			result.ParamErrorResult(c, err)
			return
		}
		validateErr := validation.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		resp, err := service.UpdateSubscribeApplication(ctx, &req)
		result.HttpResult(c, resp, err)
	}
}
