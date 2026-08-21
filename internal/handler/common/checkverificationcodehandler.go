package common

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/model/dto"
	"github.com/perfect-panel/server/internal/module/identity"
	"github.com/perfect-panel/server/internal/validation"
	"github.com/perfect-panel/server/pkg/httpx"
	"github.com/perfect-panel/server/pkg/result"
)

// CheckVerificationCodeHandler documents Check verification code.
//
// @Summary Check verification code
// @Tags common
// @Accept json
// @Produce json
// @Param request body dto.CheckVerificationCodeRequest true "Request parameters"
// @Success 200 {object} result.ResponseSuccessBean{data=dto.CheckVerificationCodeRespone}
// @Router /v1/common/check_verification_code [post]
func CheckVerificationCodeHandler(service identity.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req dto.CheckVerificationCodeRequest
		if err := httpx.ShouldBind(c, &req); err != nil {
			result.ParamErrorResult(c, err)
			return
		}
		validateErr := validation.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		resp, err := service.CheckVerificationCode(ctx, &req)
		result.HttpResult(c, resp, err)
	}
}
