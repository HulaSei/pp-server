package user

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/billing"
	"github.com/perfect-panel/server/pkg/result"
)

// QueryUserAffiliateHandler documents Query User Affiliate Count.
//
// @Summary Query User Affiliate Count
// @Tags user
// @Produce json
// @Security BearerAuth
// @Success 200 {object} result.ResponseSuccessBean{data=dto.QueryUserAffiliateCountResponse}
// @Router /v1/public/user/affiliate/count [get]
func QueryUserAffiliateHandler(service billing.Service) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {

		resp, err := service.QueryUserAffiliate(c)
		result.HttpResult(ctx, resp, err)
	}
}
