package console

import (
	"context"
	dto "github.com/perfect-panel/server/internal/module/platform/contract"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/platform"
	"github.com/perfect-panel/server/pkg/result"
)

var _ dto.RevenueStatisticsResponse

// QueryRevenueStatisticsHandler documents Query revenue statistics.
//
// @Summary Query revenue statistics
// @Tags admin
// @Produce json
// @Security BearerAuth
// @Success 200 {object} result.ResponseSuccessBean{data=dto.RevenueStatisticsResponse}
// @Router /v1/admin/console/revenue [get]
func QueryRevenueStatisticsHandler(service platform.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {

		resp, err := service.QueryRevenueStatistics(ctx)
		result.HttpResult(c, resp, err)
	}
}
