package user

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/identity"
	"github.com/perfect-panel/server/pkg/result"
)

// GetDeviceListHandler documents Get Device List.
//
// @Summary Get Device List
// @Tags user
// @Produce json
// @Security BearerAuth
// @Success 200 {object} result.ResponseSuccessBean{data=dto.GetDeviceListResponse}
// @Router /v1/public/user/devices [get]
func GetDeviceListHandler(service identity.Service) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {

		resp, err := service.GetDeviceList(c)
		result.HttpResult(ctx, resp, err)
	}
}
