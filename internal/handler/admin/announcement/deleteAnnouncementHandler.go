package announcement

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/perfect-panel/server/internal/model/dto"
	"github.com/perfect-panel/server/internal/module/support"
	"github.com/perfect-panel/server/internal/validation"
	"github.com/perfect-panel/server/pkg/httpx"
	"github.com/perfect-panel/server/pkg/result"
)

// DeleteAnnouncementHandler documents Delete announcement.
//
// @Summary Delete announcement
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.DeleteAnnouncementRequest true "Request parameters"
// @Success 200 {object} result.ResponseSuccessBean
// @Router /v1/admin/announcement/ [delete]
func DeleteAnnouncementHandler(service support.Service) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req dto.DeleteAnnouncementRequest
		if err := httpx.ShouldBind(c, &req); err != nil {
			result.ParamErrorResult(c, err)
			return
		}
		validateErr := validation.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		err := service.DeleteAnnouncement(ctx, &req)
		result.HttpResult(c, nil, err)
	}
}
