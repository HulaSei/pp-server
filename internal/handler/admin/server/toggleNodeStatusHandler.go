package server

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/model/dto"
	"github.com/perfect-panel/server/internal/module/network"
	"github.com/perfect-panel/server/internal/validation"
	"github.com/perfect-panel/server/pkg/httpx"
	"github.com/perfect-panel/server/pkg/result"
)

// ToggleNodeStatusHandler documents Toggle Node Status.
//
// @Summary Toggle Node Status
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.ToggleNodeStatusRequest true "Request parameters"
// @Success 200 {object} result.ResponseSuccessBean
// @Router /v1/admin/server/node/status/toggle [post]
func ToggleNodeStatusHandler(service network.Service) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		var req dto.ToggleNodeStatusRequest
		if err := httpx.ShouldBind(ctx, &req); err != nil {
			result.ParamErrorResult(ctx, err)
			return
		}
		validateErr := validation.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(ctx, validateErr)
			return
		}

		err := service.ToggleNodeStatus(c, &req)
		result.HttpResult(ctx, nil, err)
	}
}
