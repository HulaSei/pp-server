package user

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/identity"
	"github.com/perfect-panel/server/pkg/result"
)

// BindTelegramHandler documents Bind Telegram.
//
// @Summary Bind Telegram
// @Tags user
// @Produce json
// @Security BearerAuth
// @Success 200 {object} result.ResponseSuccessBean{data=dto.BindTelegramResponse}
// @Router /v1/public/user/bind_telegram [get]
func BindTelegramHandler(service identity.Service) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {

		resp, err := service.BindTelegram(c)
		result.HttpResult(ctx, resp, err)
	}
}
