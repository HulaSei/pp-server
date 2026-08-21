package user

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/model/dto"
	"github.com/perfect-panel/server/internal/module/identity"
	"github.com/perfect-panel/server/internal/validation"
	"github.com/perfect-panel/server/pkg/httpx"
	"github.com/perfect-panel/server/pkg/result"
)

// DeleteUserDeviceHandler documents Delete user device.
//
// @Summary Delete user device
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.DeleteUserDeivceRequest true "Request parameters"
// @Success 200 {object} result.ResponseSuccessBean
// @Router /v1/admin/user/device [delete]
func DeleteUserDeviceHandler(service identity.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req dto.DeleteUserDeivceRequest
		if err := httpx.ShouldBind(c, &req); err != nil {
			result.ParamErrorResult(c, err)
			return
		}
		validateErr := validation.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		err := service.DeleteUserDevice(ctx, &req)
		result.HttpResult(c, nil, err)
	}
}
