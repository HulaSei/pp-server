package middleware

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/perfect-panel/server/pkg/logger"
	"github.com/perfect-panel/server/pkg/result"
)

func TestLoggerMiddlewareOmitsRequestAndResponseData(t *testing.T) {
	const (
		pathSecret     = "alice@example.com"
		querySecret    = "query-secret-value"
		requestSecret  = "request-body-secret"
		responseSecret = "response-body-secret"
		parameterValue = "rejected-personal-value"
	)

	var output bytes.Buffer
	oldWriter := logger.Reset()
	logger.SetWriter(logger.NewWriter(&output))
	t.Cleanup(func() {
		logger.Reset()
		if oldWriter != nil {
			logger.SetWriter(oldWriter)
		}
	})

	ctx := app.NewContext(0)
	ctx.Request.Header.SetMethod(consts.MethodPost)
	ctx.Request.SetRequestURI("/unknown/" + pathSecret + "?token=" + querySecret)
	ctx.Request.SetBodyString(`{"password":"` + requestSecret + `"}`)
	ctx.Response.SetStatusCode(consts.StatusBadRequest)
	ctx.Response.SetBodyString(`{"token":"` + responseSecret + `"}`)
	result.ParamErrorResult(ctx, &sensitiveValidationError{value: parameterValue})

	LoggerMiddleware(nil)(context.Background(), ctx)

	got := output.String()
	for _, secret := range []string{pathSecret, querySecret, requestSecret, responseSecret, parameterValue} {
		if strings.Contains(got, secret) {
			t.Fatalf("access log contains sensitive value %q: %s", secret, got)
		}
	}
	for _, expected := range []string{`"route":"<unmatched>"`, `"method":"POST"`, `"parameter_error":true`, `"request_bytes":`} {
		if !strings.Contains(got, expected) {
			t.Fatalf("access log is missing safe diagnostic field %q: %s", expected, got)
		}
	}
}

type sensitiveValidationError struct {
	value string
}

func (e *sensitiveValidationError) Error() string {
	return e.value
}
