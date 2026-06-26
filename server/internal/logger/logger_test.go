package logger

import "testing"

func TestNewLogger(t *testing.T) {

	log := new()

	if log == nil {
		t.Fatal("expected logger instance,got nil")
	}
}
