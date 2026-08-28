package log

import (
	"strings"
	"testing"

	"github.com/perfect-panel/server/pkg/logger"
)

func TestSecurityLogMarshalRedactsPersonalDataAndCredentials(t *testing.T) {
	tests := []struct {
		name    string
		marshal func() ([]byte, error)
		secrets []string
	}{
		{
			name: "message",
			marshal: func() ([]byte, error) {
				return (&Message{
					To:       "person@example.com",
					Subject:  "Hello person@example.com",
					Content:  map[string]interface{}{"code": "123456", "body": "private body"},
					Template: "private template",
					Status:   1,
				}).Marshal()
			},
			secrets: []string{"person@example.com", "123456", "private body", "private template"},
		},
		{
			name: "login",
			marshal: func() ([]byte, error) {
				return (&Login{LoginIP: "192.0.2.10", UserAgent: "private-agent", Success: true}).Marshal()
			},
			secrets: []string{"192.0.2.10", "private-agent"},
		},
		{
			name: "registration",
			marshal: func() ([]byte, error) {
				return (&Register{Identifier: "person@example.com", RegisterIP: "192.0.2.11", UserAgent: "private-agent"}).Marshal()
			},
			secrets: []string{"person@example.com", "192.0.2.11", "private-agent"},
		},
		{
			name: "subscription",
			marshal: func() ([]byte, error) {
				return (&Subscribe{Token: "bearer-token", ClientIP: "192.0.2.12", UserAgent: "private-agent"}).Marshal()
			},
			secrets: []string{"bearer-token", "192.0.2.12", "private-agent"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			data, err := tc.marshal()
			if err != nil {
				t.Fatal(err)
			}
			text := string(data)
			for _, secret := range tc.secrets {
				if strings.Contains(text, secret) {
					t.Fatalf("serialized audit log contains %q: %s", secret, text)
				}
			}
			if !strings.Contains(text, logger.RedactedValue) {
				t.Fatalf("serialized audit log is not marked redacted: %s", text)
			}
		})
	}
}

func TestExpirableTypesNeverIncludesFinancialLedgers(t *testing.T) {
	financial := map[int]bool{
		int(TypeBalance):    true,
		int(TypeCommission): true,
		int(TypeGift):       true,
	}
	for _, typ := range ExpirableTypes() {
		if financial[typ] {
			t.Fatalf("financial log type %d must not be expirable", typ)
		}
	}
}

func TestSecurityLogUnmarshalRedactsLegacyRows(t *testing.T) {
	var message Message
	if err := message.Unmarshal([]byte(`{"to":"person@example.com","subject":"Hello","content":{"code":"123456"},"template":"secret","platform":"smtp","status":1}`)); err != nil {
		t.Fatal(err)
	}
	if message.To != logger.RedactedValue || message.Subject != logger.RedactedValue || message.Template != logger.RedactedValue || message.Content["redacted"] != true || message.Platform != "smtp" || message.Status != 1 {
		t.Fatalf("legacy message audit was not safely decoded: %+v", message)
	}

	var subscribe Subscribe
	if err := subscribe.Unmarshal([]byte(`{"token":"legacy-token","user_agent":"legacy-agent","client_ip":"192.0.2.1","user_subscribe_id":7}`)); err != nil {
		t.Fatal(err)
	}
	if subscribe.Token != logger.RedactedValue || subscribe.UserAgent != logger.RedactedValue || subscribe.ClientIP != logger.RedactedValue || subscribe.UserSubscribeId != 7 {
		t.Fatalf("legacy subscription audit was not safely decoded: %+v", subscribe)
	}

	var login Login
	if err := login.Unmarshal([]byte(`{"method":"email","login_ip":"192.0.2.2","user_agent":"legacy-agent","success":true}`)); err != nil {
		t.Fatal(err)
	}
	if login.LoginIP != logger.RedactedValue || login.UserAgent != logger.RedactedValue || !login.Success {
		t.Fatalf("legacy login audit was not safely decoded: %+v", login)
	}
}
