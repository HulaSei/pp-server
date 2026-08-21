package document

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/module/support"
	dto "github.com/perfect-panel/server/internal/module/support/contract"
	"github.com/perfect-panel/server/internal/validation"
	"github.com/perfect-panel/server/pkg/httpx"
	"github.com/perfect-panel/server/pkg/result"
)

// CreateDocumentHandler documents Create document.
//
// @Summary Create document
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateDocumentRequest true "Request parameters"
// @Success 200 {object} result.ResponseSuccessBean
// @Router /v1/admin/document/ [post]
func CreateDocumentHandler(service support.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req dto.CreateDocumentRequest
		if err := httpx.ShouldBind(c, &req); err != nil {
			result.ParamErrorResult(c, err)
			return
		}
		validateErr := validation.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		err := service.CreateDocument(ctx, &req)
		result.HttpResult(c, nil, err)
	}
}
