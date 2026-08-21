package subscribe

import (
	"context"
	dto "github.com/perfect-panel/server/internal/module/subscription/contract"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/subscription"
	"github.com/perfect-panel/server/pkg/result"
)

var _ dto.GetSubscribeGroupListResponse

// GetSubscribeGroupListHandler documents Get subscribe group list.
//
// @Summary Get subscribe group list
// @Tags admin
// @Produce json
// @Security BearerAuth
// @Success 200 {object} result.ResponseSuccessBean{data=dto.GetSubscribeGroupListResponse}
// @Router /v1/admin/subscribe/group/list [get]
func GetSubscribeGroupListHandler(service subscription.Service) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {

		resp, err := service.GetSubscribeGroupList(c)
		result.HttpResult(ctx, resp, err)
	}
}
