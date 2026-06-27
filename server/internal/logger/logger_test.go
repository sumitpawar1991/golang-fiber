package logger

import (
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
