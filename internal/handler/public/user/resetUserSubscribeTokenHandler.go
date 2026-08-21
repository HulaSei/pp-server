package user

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/model/dto"
	"github.com/perfect-panel/server/internal/module/subscription"
	"github.com/perfect-panel/server/internal/validation"
	"github.com/perfect-panel/server/pkg/httpx"
	"github.com/perfect-panel/server/pkg/result"
)

// ResetUserSubscribeTokenHandler documents Reset User Subscribe Token.
//
// @Summary Reset User Subscribe Token
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.ResetUserSubscribeTokenRequest true "Request parameters"
// @Success 200 {object} result.ResponseSuccessBean
// @Router /v1/public/user/subscribe_token [put]
func ResetUserSubscribeTokenHandler(service subscription.Service) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		var req dto.ResetUserSubscribeTokenRequest
		if err := httpx.ShouldBind(ctx, &req); err != nil {
			result.ParamErrorResult(ctx, err)
			return
		}
		validateErr := validation.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(ctx, validateErr)
			return
		}

		err := service.ResetOwnSubscribeToken(c, &req)
		result.HttpResult(ctx, nil, err)
	}
}
