package user

import (
	"context"
	dto "github.com/perfect-panel/server/internal/module/identity/contract"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/identity"
	"github.com/perfect-panel/server/pkg/result"
)

var _ dto.GetOAuthMethodsResponse

// GetOAuthMethodsHandler documents Get OAuth Methods.
//
// @Summary Get OAuth Methods
// @Tags user
// @Produce json
// @Security BearerAuth
// @Success 200 {object} result.ResponseSuccessBean{data=dto.GetOAuthMethodsResponse}
// @Router /v1/public/user/oauth_methods [get]
func GetOAuthMethodsHandler(service identity.Service) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {

		resp, err := service.GetOAuthMethods(c)
		result.HttpResult(ctx, resp, err)
	}
}
