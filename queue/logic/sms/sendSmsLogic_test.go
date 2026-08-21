package smslogic

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/perfect-panel/server/pkg/logger"
)

func TestSMSMessageLogDoesNotPersistRecipientOrCode(t *testing.T) {
	message := newSMSMessageLog("provider", 1)
	encoded, err := json.Marshal(message)
	if err != nil {
		t.Fatalf("marshal message log: %v", err)
	}

	got := string(encoded)
	if message.To != logger.RedactedValue {
		t.Fatalf("recipient was not redacted: %q", message.To)
	}
	for _, sensitive := range []string{"+6591234567", "654321"} {
		if strings.Contains(got, sensitive) {
			t.Fatalf("message log contains sensitive value %q: %s", sensitive, got)
		}
	}
	if redacted, ok := message.Content["redacted"].(bool); !ok || !redacted {
		t.Fatalf("message content is not marked redacted: %#v", message.Content)
	}
}
