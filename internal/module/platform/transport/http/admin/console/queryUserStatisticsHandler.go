package console

import (
	"context"
	dto "github.com/perfect-panel/server/internal/module/platform/contract"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/platform"
	"github.com/perfect-panel/server/pkg/result"
)

var _ dto.UserStatisticsResponse

// QueryUserStatisticsHandler documents Query user statistics.
//
// @Summary Query user statistics
// @Tags admin
// @Produce json
// @Security BearerAuth
// @Success 200 {object} result.ResponseSuccessBean{data=dto.UserStatisticsResponse}
// @Router /v1/admin/console/user [get]
func QueryUserStatisticsHandler(service platform.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {

		resp, err := service.QueryUserStatistics(ctx)
		result.HttpResult(c, resp, err)
	}
}
