package log

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/platform"
	dto "github.com/perfect-panel/server/internal/module/platform/contract"
	"github.com/perfect-panel/server/internal/validation"
	"github.com/perfect-panel/server/pkg/httpx"
	"github.com/perfect-panel/server/pkg/result"
)

// FilterEmailLogHandler documents Filter email log.
//
// @Summary Filter email log
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request query dto.FilterLogParams false "Request parameters"
// @Success 200 {object} result.ResponseSuccessBean{data=dto.FilterEmailLogResponse}
// @Router /v1/admin/log/email/list [get]
func FilterEmailLogHandler(service platform.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req dto.FilterLogParams
		if err := httpx.ShouldBind(c, &req); err != nil {
			result.ParamErrorResult(c, err)
			return
		}
		validateErr := validation.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		resp, err := service.FilterEmailLog(ctx, &req)
		result.HttpResult(c, resp, err)
	}
}
