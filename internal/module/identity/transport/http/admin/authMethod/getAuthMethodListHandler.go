package authMethod

import (
	"context"
	dto "github.com/perfect-panel/server/internal/module/identity/contract"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/identity"
	"github.com/perfect-panel/server/pkg/result"
)

var _ dto.GetAuthMethodListResponse

// GetAuthMethodListHandler documents Get auth method list.
//
// @Summary Get auth method list
// @Tags admin
// @Produce json
// @Security BearerAuth
// @Success 200 {object} result.ResponseSuccessBean{data=dto.GetAuthMethodListResponse}
// @Router /v1/admin/auth-method/list [get]
func GetAuthMethodListHandler(service identity.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {

		resp, err := service.GetAuthMethodList(ctx)
		result.HttpResult(c, resp, err)
	}
}
