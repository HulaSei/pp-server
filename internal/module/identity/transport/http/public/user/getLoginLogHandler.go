package user

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/identity"
	dto "github.com/perfect-panel/server/internal/module/identity/contract"
	"github.com/perfect-panel/server/internal/validation"
	"github.com/perfect-panel/server/pkg/httpx"
	"github.com/perfect-panel/server/pkg/result"
)

// GetLoginLogHandler documents Get Login Log.
//
// @Summary Get Login Log
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request query dto.GetLoginLogRequest false "Request parameters"
// @Success 200 {object} result.ResponseSuccessBean{data=dto.GetLoginLogResponse}
// @Router /v1/public/user/login_log [get]
func GetLoginLogHandler(service identity.Service) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		var req dto.GetLoginLogRequest
		if err := httpx.ShouldBind(ctx, &req); err != nil {
			result.ParamErrorResult(ctx, err)
			return
		}
		validateErr := validation.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(ctx, validateErr)
			return
		}

		resp, err := service.GetLoginLog(c, &req)
		result.HttpResult(ctx, resp, err)
	}
}
