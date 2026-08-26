package user

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/identity"
	dto "github.com/perfect-panel/server/internal/module/identity/contract"
	"github.com/perfect-panel/server/internal/validation"
	"github.com/perfect-panel/server/pkg/httpx"
	"github.com/perfect-panel/server/pkg/result"
)

// VerifyEmailHandler documents Verify Email.
//
// @Summary Verify Email
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.VerifyEmailRequest true "Request parameters"
// @Success 200 {object} result.ResponseSuccessBean
// @Router /v1/public/user/verify_email [post]
func VerifyEmailHandler(service identity.Service) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		var req dto.VerifyEmailRequest
		if err := httpx.ShouldBind(ctx, &req); err != nil {
			result.ParamErrorResult(ctx, err)
			return
		}
		validateErr := validation.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(ctx, validateErr)
			return
		}

		err := service.VerifyEmail(c, &req)
		result.HttpResult(ctx, nil, err)
	}
}
