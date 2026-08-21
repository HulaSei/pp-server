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

// GetServerProtocolsHandler documents Get Server Protocols.
//
// @Summary Get Server Protocols
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request query dto.GetServerProtocolsRequest false "Request parameters"
// @Success 200 {object} result.ResponseSuccessBean{data=dto.GetServerProtocolsResponse}
// @Router /v1/admin/server/protocols [get]
func GetServerProtocolsHandler(service network.Service) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		var req dto.GetServerProtocolsRequest
		if err := httpx.ShouldBind(ctx, &req); err != nil {
			result.ParamErrorResult(ctx, err)
			return
		}
		validateErr := validation.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(ctx, validateErr)
			return
		}

		resp, err := service.GetServerProtocols(c, &req)
		result.HttpResult(ctx, resp, err)
	}
}
