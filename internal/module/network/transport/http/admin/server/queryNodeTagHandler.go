package server

import (
	"context"
	dto "github.com/perfect-panel/server/internal/module/network/contract"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/network"
	"github.com/perfect-panel/server/pkg/result"
)

var _ dto.QueryNodeTagResponse

// QueryNodeTagHandler documents Query all node tags.
//
// @Summary Query all node tags
// @Tags admin
// @Produce json
// @Security BearerAuth
// @Success 200 {object} result.ResponseSuccessBean{data=dto.QueryNodeTagResponse}
// @Router /v1/admin/server/node/tags [get]
func QueryNodeTagHandler(service network.Service) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {

		resp, err := service.QueryNodeTag(c)
		result.HttpResult(ctx, resp, err)
	}
}
