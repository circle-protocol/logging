package logging

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"log/slog"
)

// newTestLogger creates a logger for testing that writes to a string builder
// and returns both the builder and the logger
func newTestLogger() (out *strings.Builder, logger *slog.Logger) {
	out = new(strings.Builder)
	handler := slog.NewJSONHandler(out, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}).WithAttrs([]slog.Attr{slog.String("time", "not")})
	return out, slog.New(handler)
}

func TestRequestToAttr(t *testing.T) {
	out, logger := newTestLogger()
	logger.Info("test", RequestToAttr(
		httptest.NewRequest("GET", "/target", nil),
	))

	want := `{
		"level":"INFO",
		"msg":"test",
		"time":"not",
		"request":{
			"method":"GET",
			"url":"/target"
		}
	}`
	got := out.String()
	assert.JSONEq(t, want, got)
}
