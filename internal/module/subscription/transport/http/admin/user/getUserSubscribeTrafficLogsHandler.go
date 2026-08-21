package user

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/subscription"
	dto "github.com/perfect-panel/server/internal/module/subscription/contract"
	"github.com/perfect-panel/server/internal/validation"
	"github.com/perfect-panel/server/pkg/httpx"
	"github.com/perfect-panel/server/pkg/result"
)

// GetUserSubscribeTrafficLogsHandler documents Get user subcribe traffic logs.
//
// @Summary Get user subcribe traffic logs
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request query dto.GetUserSubscribeTrafficLogsRequest false "Request parameters"
// @Success 200 {object} result.ResponseSuccessBean{data=dto.GetUserSubscribeTrafficLogsResponse}
// @Router /v1/admin/user/subscribe/traffic_logs [get]
func GetUserSubscribeTrafficLogsHandler(service subscription.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req dto.GetUserSubscribeTrafficLogsRequest
		if err := httpx.ShouldBind(c, &req); err != nil {
			result.ParamErrorResult(c, err)
			return
		}
		validateErr := validation.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		resp, err := service.GetUserSubscribeTrafficLogs(ctx, &req)
		result.HttpResult(c, resp, err)
	}
}
