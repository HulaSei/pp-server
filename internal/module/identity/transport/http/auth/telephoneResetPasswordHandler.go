package auth

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/config"
	"github.com/perfect-panel/server/internal/module/identity"
	dto "github.com/perfect-panel/server/internal/module/identity/contract"
	"github.com/perfect-panel/server/internal/validation"
	"github.com/perfect-panel/server/pkg/httpx"
	"github.com/perfect-panel/server/pkg/result"
	"github.com/perfect-panel/server/pkg/turnstile"
	"github.com/perfect-panel/server/pkg/xerr"
	"github.com/pkg/errors"
)

// TelephoneResetPasswordHandler documents Reset password.
//
// @Summary Reset password
// @Tags common
// @Accept json
// @Produce json
// @Param request body dto.TelephoneResetPasswordRequest true "Request parameters"
// @Success 200 {object} result.ResponseSuccessBean{data=dto.LoginResponse}
// @Router /v1/auth/reset/telephone [post]
func TelephoneResetPasswordHandler(service identity.Service, verifyConfig func() config.Verify) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req dto.TelephoneResetPasswordRequest
		if err := httpx.ShouldBind(c, &req); err != nil {
			result.ParamErrorResult(c, err)
			return
		}
		validateErr := validation.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}
		// get client ip
		req.IP = c.ClientIP()
		req.UserAgent = string(c.UserAgent())
		verify := verifyConfig()
		if verify.ResetPasswordVerify {
			verifyTurns := turnstile.New(turnstile.Config{
				Secret:  verify.TurnstileSecret,
				Timeout: 3 * time.Second,
			})
			if verify, err := verifyTurns.Verify(ctx, req.CfToken, req.IP); err != nil || !verify {
				err = errors.Wrapf(xerr.NewErrCode(xerr.TooManyRequests), "error: %v, verify: %v", err, verify)
				result.HttpResult(c, nil, err)
				return
			}
		}
		resp, err := service.TelephoneResetPassword(ctx, &req)
		result.HttpResult(c, resp, err)
	}
}
