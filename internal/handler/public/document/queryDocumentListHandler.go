package document

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/support"
	"github.com/perfect-panel/server/pkg/result"
)

// QueryDocumentListHandler documents Get document list.
//
// @Summary Get document list
// @Tags user
// @Produce json
// @Security BearerAuth
// @Success 200 {object} result.ResponseSuccessBean{data=dto.QueryDocumentListResponse}
// @Router /v1/public/document/list [get]
func QueryDocumentListHandler(service support.Service) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {

		resp, err := service.QueryDocumentList(c)
		result.HttpResult(ctx, resp, err)
	}
}
