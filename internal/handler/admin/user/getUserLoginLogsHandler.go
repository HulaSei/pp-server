package user

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/model/dto"
	"github.com/perfect-panel/server/internal/module/identity"
	"github.com/perfect-panel/server/internal/validation"
	"github.com/perfect-panel/server/pkg/httpx"
	"github.com/perfect-panel/server/pkg/result"
)

// GetUserLoginLogsHandler documents Get user login logs.
//
// @Summary Get user login logs
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request query dto.GetUserLoginLogsRequest false "Request parameters"
// @Success 200 {object} result.ResponseSuccessBean{data=dto.GetUserLoginLogsResponse}
// @Router /v1/admin/user/login/logs [get]
func GetUserLoginLogsHandler(service identity.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req dto.GetUserLoginLogsRequest
		if err := httpx.ShouldBind(c, &req); err != nil {
			result.ParamErrorResult(c, err)
			return
		}
		validateErr := validation.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		resp, err := service.GetUserLoginLogs(ctx, &req)
		result.HttpResult(c, resp, err)
	}
}
