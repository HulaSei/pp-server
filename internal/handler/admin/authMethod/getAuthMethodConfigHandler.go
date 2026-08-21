package authMethod

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/model/dto"
	"github.com/perfect-panel/server/internal/module/identity"
	"github.com/perfect-panel/server/internal/validation"
	"github.com/perfect-panel/server/pkg/httpx"
	"github.com/perfect-panel/server/pkg/result"
)

// GetAuthMethodConfigHandler documents Get auth method config.
//
// @Summary Get auth method config
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request query dto.GetAuthMethodConfigRequest false "Request parameters"
// @Success 200 {object} result.ResponseSuccessBean{data=dto.AuthMethodConfig}
// @Router /v1/admin/auth-method/config [get]
func GetAuthMethodConfigHandler(service identity.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req dto.GetAuthMethodConfigRequest
		if err := httpx.ShouldBind(c, &req); err != nil {
			result.ParamErrorResult(c, err)
			return
		}
		validateErr := validation.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		resp, err := service.GetAuthMethodConfig(ctx, &req)
		result.HttpResult(c, resp, err)
	}
}
