package user

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/identity"
	"github.com/perfect-panel/server/pkg/result"
)

// QueryUserInfoHandler documents returns the current user profile..
//
// @Summary returns the current user profile.
// @Tags user
// @Produce json
// @Security BearerAuth
// @Success 200 {object} result.ResponseSuccessBean{data=dto.User}
// @Router /v1/public/user/info [get]
func QueryUserInfoHandler(service identity.Service) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {

		resp, err := service.QueryUserInfo(c)
		result.HttpResult(ctx, resp, err)
	}
}
