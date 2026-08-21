package common

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/identity"
	dto "github.com/perfect-panel/server/internal/module/identity/contract"
	"github.com/perfect-panel/server/internal/validation"
	"github.com/perfect-panel/server/pkg/httpx"
	"github.com/perfect-panel/server/pkg/result"
)

// SendSmsCodeHandler documents Get sms verification code.
//
// @Summary Get sms verification code
// @Tags common
// @Accept json
// @Produce json
// @Param request body dto.SendSmsCodeRequest true "Request parameters"
// @Success 200 {object} result.ResponseSuccessBean{data=dto.SendCodeResponse}
// @Router /v1/common/send_sms_code [post]
func SendSmsCodeHandler(service identity.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req dto.SendSmsCodeRequest
		if err := httpx.ShouldBind(c, &req); err != nil {
			result.ParamErrorResult(c, err)
			return
		}
		validateErr := validation.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		resp, err := service.SendSmsCode(ctx, &req)
		result.HttpResult(c, resp, err)
	}
}
