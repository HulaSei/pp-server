package user

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/identity"
	"github.com/perfect-panel/server/pkg/result"
)

// CurrentUserHandler documents Current user.
//
// @Summary Current user
// @Tags admin
// @Produce json
// @Security BearerAuth
// @Success 200 {object} result.ResponseSuccessBean{data=dto.User}
// @Router /v1/admin/user/current [get]
func CurrentUserHandler(service identity.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		resp, err := service.CurrentUser(ctx)
		result.HttpResult(c, resp, err)
	}
}
