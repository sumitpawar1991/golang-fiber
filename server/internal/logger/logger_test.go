package logger

import (
	"bytes"
	"strings"
	"testing"
)

func TestNewLogger(t *testing.T) {

	t.Run("should create a new logger with default config", func(t *testing.T) {
		log := New(Config{
			Level: "info",
		})

		if log == nil {
			t.Error("expected logger to be created, got nil")
		}

	})
}

func TestLogger_Info(t *testing.T) {

	t.Run("should log info message", func(t *testing.T) {

		var buf bytes.Buffer

		log := New(Config{
			Level: "Info",
		})

		log.(*logger).logger = log.(*logger).logger.output(buf)

		log.Info("test message")

		if log == nil {
			t.Error("expected logger to be created, got nil")
		}

		output := buf.String()

		if !strings.Contains(output, "test message") {
			t.Errorf("expected message 'test message' in output , got %s", output)
		}

		if !strings.Contains(output, "info") {
			t.Errorf("expected level info in output,got %s", output)
		}
	})
}
