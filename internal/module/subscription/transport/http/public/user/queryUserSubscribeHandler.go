package user

import (
	"context"
	dto "github.com/perfect-panel/server/internal/module/subscription/contract"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/subscription"
	"github.com/perfect-panel/server/pkg/result"
)

var _ dto.QueryUserSubscribeListResponse

// QueryUserSubscribeHandler documents Query User Subscribe.
//
// @Summary Query User Subscribe
// @Tags user
// @Produce json
// @Security BearerAuth
// @Success 200 {object} result.ResponseSuccessBean{data=dto.QueryUserSubscribeListResponse}
// @Router /v1/public/user/subscribe [get]
func QueryUserSubscribeHandler(service subscription.Service) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {

		resp, err := service.QueryUserSubscribe(c)
		result.HttpResult(ctx, resp, err)
	}
}
