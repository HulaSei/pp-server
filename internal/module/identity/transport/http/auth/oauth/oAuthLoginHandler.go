package oauth

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/identity"
	dto "github.com/perfect-panel/server/internal/module/identity/contract"
	"github.com/perfect-panel/server/internal/validation"
	"github.com/perfect-panel/server/pkg/httpx"
	"github.com/perfect-panel/server/pkg/result"
)

// OAuthLoginHandler documents OAuth login.
//
// @Summary OAuth login
// @Tags common
// @Accept json
// @Produce json
// @Param request body dto.OAthLoginRequest true "Request parameters"
// @Success 200 {object} result.ResponseSuccessBean{data=dto.OAuthLoginResponse}
// @Router /v1/auth/oauth/login [post]
func OAuthLoginHandler(service identity.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req dto.OAthLoginRequest
		if err := httpx.ShouldBind(c, &req); err != nil {
			result.ParamErrorResult(c, err)
			return
		}
		validateErr := validation.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		resp, err := service.OAuthLogin(ctx, &req)
		result.HttpResult(c, resp, err)
	}
}
