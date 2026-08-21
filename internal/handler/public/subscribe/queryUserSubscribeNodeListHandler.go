package subscribe

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/subscription"
	"github.com/perfect-panel/server/pkg/result"
)

// QueryUserSubscribeNodeListHandler documents Get user subscribe node info.
//
// @Summary Get user subscribe node info
// @Tags user
// @Produce json
// @Security BearerAuth
// @Success 200 {object} result.ResponseSuccessBean{data=dto.QueryUserSubscribeNodeListResponse}
// @Router /v1/public/subscribe/node/list [get]
func QueryUserSubscribeNodeListHandler(service subscription.Service) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {

		resp, err := service.QueryUserSubscribeNodeList(c)
		result.HttpResult(ctx, resp, err)
	}
}
