package stdlogtoapex

import (
	"io"
	"log"
	"regexp"
	"testing"

	alog "github.com/apex/log"
	"github.com/apex/log/handlers/memory"
)

func TestNewWriter(t *testing.T) {
	var _ io.Writer = NewWriter(nil)
}

func TestSetOutputToWriter(t *testing.T) {
	handler := memory.New()
	alog.SetHandler(handler)
	writer := NewWriter(handler)
	log.SetOutput(writer)
	log.Print("Hello!")
	if len(handler.Entries) != 1 {
		t.Fatal("No log message recorder in handler")
	}
	rx := "^[0-9]{4}/[0-9]{2}/[0-9]{2} [0-2][0-9]:[0-5][0-9]:[0-5][0-9] Hello!\n$"
	entry := handler.Entries[0]
	matched, err := regexp.MatchString(rx, entry.Message)
	if err != nil {
		t.Fatalf("Error matching regexp: %s", err)
	}
	if !matched {
		t.Errorf("Expected log message %q to match regexp %q", entry.Message, rx)
	}
	if entry.Level != alog.InfoLevel {
		t.Errorf("Expected to log at the default level (%s), but logged at %s.", alog.InfoLevel, entry.Level)
	}
}
