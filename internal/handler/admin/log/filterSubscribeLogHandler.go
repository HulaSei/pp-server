package log

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/model/dto"
	"github.com/perfect-panel/server/internal/module/platform"
	"github.com/perfect-panel/server/internal/validation"
	"github.com/perfect-panel/server/pkg/httpx"
	"github.com/perfect-panel/server/pkg/result"
)

// FilterSubscribeLogHandler documents Filter subscribe log.
//
// @Summary Filter subscribe log
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request query dto.FilterSubscribeLogRequest false "Request parameters"
// @Success 200 {object} result.ResponseSuccessBean{data=dto.FilterSubscribeLogResponse}
// @Router /v1/admin/log/subscribe/list [get]
func FilterSubscribeLogHandler(service platform.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req dto.FilterSubscribeLogRequest
		if err := httpx.ShouldBind(c, &req); err != nil {
			result.ParamErrorResult(c, err)
			return
		}
		validateErr := validation.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		resp, err := service.FilterSubscribeLog(ctx, &req)
		result.HttpResult(c, resp, err)
	}
}
