package middleware

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/perfect-panel/server/pkg/logger"
	"github.com/perfect-panel/server/pkg/result"
)

func LoggerMiddleware() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		start := time.Now()
		ctx.Next(c)

		cost := time.Since(start)
		responseStatus := responseStatus(ctx)
		method := string(ctx.Method())
		route := ctx.FullPath()
		if route == "" {
			// Do not fall back to the raw path: unmatched paths are controlled by
			// the caller and may themselves contain credentials or personal data.
			route = "<unmatched>"
		}

		logs := []logger.LogField{
			logger.Field("status", responseStatus),
			logger.Field("method", method),
			logger.Field("route", route),
			logger.Field("request_bytes", len(ctx.Request.Body())),
			logger.Field("response_bytes", len(ctx.Response.Body())),
		}
		if result.ParamErrorFromRequestContext(ctx) != nil {
			// Parameter errors can echo rejected input values. Preserve the fact
			// that validation failed without persisting the submitted value.
			logs = append(logs, logger.Field("parameter_error", true))
		}
		logs = append(logs, logger.Field("duration", cost))
		if responseStatus >= 500 && responseStatus <= 599 {
			logger.WithContext(c).Errorw("HTTP Error", logs...)
		} else {
			logger.WithContext(c).Infow("HTTP Request", logs...)
		}
	}
}

func responseStatus(ctx *app.RequestContext) int {
	status := ctx.Response.StatusCode()
	if status == 0 {
		return consts.StatusOK
	}
	return status
}
