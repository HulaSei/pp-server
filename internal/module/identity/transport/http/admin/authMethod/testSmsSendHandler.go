package authMethod

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/identity"
	dto "github.com/perfect-panel/server/internal/module/identity/contract"
	"github.com/perfect-panel/server/internal/validation"
	"github.com/perfect-panel/server/pkg/httpx"
	"github.com/perfect-panel/server/pkg/result"
)

// TestSmsSendHandler documents Test sms send.
//
// @Summary Test sms send
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.TestSmsSendRequest true "Request parameters"
// @Success 200 {object} result.ResponseSuccessBean
// @Router /v1/admin/auth-method/test_sms_send [post]
func TestSmsSendHandler(service identity.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req dto.TestSmsSendRequest
		if err := httpx.ShouldBind(c, &req); err != nil {
			result.ParamErrorResult(c, err)
			return
		}
		validateErr := validation.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		err := service.TestSmsSend(ctx, &req)
		result.HttpResult(c, nil, err)
	}
}
