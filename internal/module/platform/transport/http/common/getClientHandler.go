package common

import (
	"context"
	dto "github.com/perfect-panel/server/internal/module/platform/contract"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/platform"
	"github.com/perfect-panel/server/pkg/result"
)

var _ dto.GetSubscribeClientResponse

// GetClientHandler documents Get Client.
//
// @Summary Get Client
// @Tags common
// @Produce json
// @Success 200 {object} result.ResponseSuccessBean{data=dto.GetSubscribeClientResponse}
// @Router /v1/common/client [get]
func GetClientHandler(service platform.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		resp, err := service.GetClient(ctx)
		result.HttpResult(c, resp, err)
	}
}
