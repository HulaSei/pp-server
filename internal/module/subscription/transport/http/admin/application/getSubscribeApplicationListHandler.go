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

// GetSubscribeApplicationListHandler documents Get subscribe application list.
//
// @Summary Get subscribe application list
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request query dto.GetSubscribeApplicationListRequest false "Request parameters"
// @Success 200 {object} result.ResponseSuccessBean{data=dto.GetSubscribeApplicationListResponse}
// @Router /v1/admin/application/subscribe_application_list [get]
func GetSubscribeApplicationListHandler(service subscription.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req dto.GetSubscribeApplicationListRequest
		if err := httpx.ShouldBind(c, &req); err != nil {
			result.ParamErrorResult(c, err)
			return
		}
		validateErr := validation.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		resp, err := service.GetSubscribeApplicationList(ctx, &req)
		result.HttpResult(c, resp, err)
	}
}
