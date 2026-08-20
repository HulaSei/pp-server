package logger

import (
	"fmt"
	"regexp"
	"strings"
)

const RedactedValue = "[REDACTED]"

var (
	jwtPattern              = regexp.MustCompile(`\b[A-Za-z0-9_-]{10,}\.[A-Za-z0-9_-]{10,}\.[A-Za-z0-9_-]{10,}\b`)
	telegramBotTokenPattern = regexp.MustCompile(`\b[0-9]{6,12}:AA[A-Za-z0-9_-]{20,}\b`)
	emailPattern            = regexp.MustCompile(`\b[A-Za-z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[A-Za-z0-9](?:[A-Za-z0-9-]{0,61}[A-Za-z0-9])?(?:\.[A-Za-z0-9](?:[A-Za-z0-9-]{0,61}[A-Za-z0-9])?)+\b`)
	bearerPattern           = regexp.MustCompile(`(?i)\bbearer\s+[A-Za-z0-9._~+/=-]+`)
	sensitiveQueryPattern   = regexp.MustCompile(`(?i)([?&](?:access_?token|auth_?code|code|key|password|secret|signature|token)=)[^&#\s]+`)
)

// sensitiveFieldKey classifies fields that must never reach a log sink. Keep
// this list deliberately broad: losing a diagnostic value is preferable to
// persisting a credential, message body or personal identifier.
func sensitiveFieldKey(key string) bool {
	normalized := strings.ToLower(strings.TrimSpace(key))
	normalized = strings.NewReplacer("-", "_", ".", "_", " ", "_").Replace(normalized)
	compact := strings.ReplaceAll(normalized, "_", "")

	for _, marker := range []string{
		"token", "secret", "password", "passwd", "credential", "authorization",
		"cookie", "signature", "privatekey", "apikey", "accesskey", "session", "uuid",
		"email", "telephone", "phone", "useragent", "clientip", "deviceid",
		"chatid", "openid", "uniqueid", "senderid",
	} {
		if strings.Contains(compact, marker) {
			return true
		}
	}

	switch compact {
	case "body", "requestbody", "responsebody", "request", "response", "payload",
		"content", "query", "params", "form", "config", "headers", "header",
		"sql", "user", "users", "userinfo", "order", "orders", "orderinfo",
		"subscribe", "subscribers", "req", "template", "value", "cachekey", "coupon",
		"email", "telephone", "phone", "ip", "clientip", "useragent", "identifier",
		"openid", "uniqueid", "chatid", "senderid", "recipient", "subject", "text",
		"command", "redirect", "redirecturl", "url", "code", "refercode", "deductionlog":
		return true
	default:
		return false
	}
}

func redactText(value string) string {
	value = jwtPattern.ReplaceAllString(value, RedactedValue)
	value = telegramBotTokenPattern.ReplaceAllString(value, RedactedValue)
	value = emailPattern.ReplaceAllString(value, RedactedValue)
	value = bearerPattern.ReplaceAllString(value, "Bearer "+RedactedValue)
	return sensitiveQueryPattern.ReplaceAllString(value, `${1}`+RedactedValue)
}

func redactField(field LogField) LogField {
	if sensitiveFieldKey(field.Key) {
		field.Value = RedactedValue
		return field
	}
	field.Value = redactValue(field.Value)
	return field
}

func redactFields(fields []LogField) []LogField {
	if len(fields) == 0 {
		return fields
	}
	redacted := make([]LogField, len(fields))
	for i, field := range fields {
		redacted[i] = redactField(field)
	}
	return redacted
}

func redactValue(value any) any {
	switch v := value.(type) {
	case string:
		return redactText(v)
	case error:
		return redactText(v.Error())
	case fmt.Stringer:
		return redactText(encodeStringer(v))
	case []string:
		redacted := make([]string, len(v))
		for i, item := range v {
			redacted[i] = redactText(item)
		}
		return redacted
	default:
		return value
	}
}

// splitLogArgs preserves the convenient logger.Info("message", Field(...))
// call style used by older code while keeping the fields structured and
// therefore subject to the same redaction policy as InfoW/ErrorW calls.
func splitLogArgs(args []any) (string, []LogField) {
	message := make([]any, 0, len(args))
	fields := make([]LogField, 0, len(args))
	for _, arg := range args {
		switch v := arg.(type) {
		case LogField:
			fields = append(fields, redactField(v))
		case []LogField:
			fields = append(fields, redactFields(v)...)
		default:
			message = append(message, arg)
		}
	}
	return redactText(fmt.Sprint(message...)), fields
}
