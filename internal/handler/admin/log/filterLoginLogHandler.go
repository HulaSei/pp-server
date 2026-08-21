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

// FilterLoginLogHandler documents Filter login log.
//
// @Summary Filter login log
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request query dto.FilterLoginLogRequest false "Request parameters"
// @Success 200 {object} result.ResponseSuccessBean{data=dto.FilterLoginLogResponse}
// @Router /v1/admin/log/login/list [get]
func FilterLoginLogHandler(service platform.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req dto.FilterLoginLogRequest
		if err := httpx.ShouldBind(c, &req); err != nil {
			result.ParamErrorResult(c, err)
			return
		}
		validateErr := validation.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		resp, err := service.FilterLoginLog(ctx, &req)
		result.HttpResult(c, resp, err)
	}
}
