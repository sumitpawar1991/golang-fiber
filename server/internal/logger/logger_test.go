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

func TestLogger_Info(t *testing.T) {

	t.Run("should log info message", func(t *testing.T) {

		//	var buf bytes.Buffer

		log := New(Config{
			Level: "debug",
		})

		if log == nil {
			t.Error("expected logger to be created, got nil")
		}
	})
}
