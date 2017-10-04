package stdlogtoapex

import (
	"io"
	"log"
	"testing"

	alog "github.com/apex/log"
	"github.com/apex/log/handlers/memory"
)

func TestNewWriter(t *testing.T) {
	var _ io.Writer = NewWriter()
}

func TestSetOutputToWriter(t *testing.T) {
	handler := memory.New()
	alog.SetHandler(handler)
	writer := NewWriter()
	log.SetOutput(writer)
	log.Print("Hello!")
	if len(handler.Entries) != 1 {
		t.Fatal("No log message recorder in handler")
	}

	entry := handler.Entries[0]
	expected := "Hello!\n"
	if entry.Message != expected {
		t.Errorf("Expected log message %q, got %q", expected, entry.Message)
	}
	if entry.Level != alog.InfoLevel {
		t.Errorf("Expected to log at the default level (%s), but logged at %s.", alog.InfoLevel, entry.Level)
	}
}
