package system

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/model/dto"
	"github.com/perfect-panel/server/internal/module/platform"
	"github.com/perfect-panel/server/internal/validation"
	"github.com/perfect-panel/server/pkg/httpx"
	"github.com/perfect-panel/server/pkg/result"
)

// UpdateVerifyCodeConfigHandler documents Update Verify Code Config.
//
// @Summary Update Verify Code Config
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.VerifyCodeConfig true "Request parameters"
// @Success 200 {object} result.ResponseSuccessBean
// @Router /v1/admin/system/verify_code_config [put]
func UpdateVerifyCodeConfigHandler(service platform.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req dto.VerifyCodeConfig
		if err := httpx.ShouldBind(c, &req); err != nil {
			result.ParamErrorResult(c, err)
			return
		}
		validateErr := validation.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		err := service.UpdateVerifyCodeConfig(ctx, &req)
		result.HttpResult(c, nil, err)
	}
}
