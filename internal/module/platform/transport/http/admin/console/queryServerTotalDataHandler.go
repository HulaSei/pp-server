package console

import (
	"context"
	dto "github.com/perfect-panel/server/internal/module/platform/contract"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/platform"
	"github.com/perfect-panel/server/pkg/result"
)

var _ dto.ServerTotalDataResponse

// QueryServerTotalDataHandler documents Query server total data.
//
// @Summary Query server total data
// @Tags admin
// @Produce json
// @Security BearerAuth
// @Success 200 {object} result.ResponseSuccessBean{data=dto.ServerTotalDataResponse}
// @Router /v1/admin/console/server [get]
func QueryServerTotalDataHandler(service platform.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {

		resp, err := service.QueryServerTotalData(ctx)
		result.HttpResult(c, resp, err)
	}
}
