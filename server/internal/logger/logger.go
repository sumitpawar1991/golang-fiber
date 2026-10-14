package logger

import (
	"io"
	"os"
	"strings"
	"sync"
)

type Level int

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

var levelNames = map[Level]string{
	DebugLevel: "DEBUG",
	InfoLevel:  "INFO",
	WarnLevel:  "WARN",
	ErrorLevel: "ERROR",
}

// config hold logger information
type Config struct {
	Level  string
	Output io.Writer
}

type logger struct {
	mu     sync.Mutex
	level  Level
	output io.Writer
	fields map[string]interface{}
}

type Logger interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	With(args ...interface{}) Logger
}

func New(cfg Config) Logger {
	// Parse log level
	level := parseLevel(cfg.Level)

	// Set default output
	output := cfg.Output
	if output == nil {
		output = os.Stdout
	}

	return &logger{
		level:  level,
		output: output,
		fields: make(map[string]interface{}),
	}
}

func parseLevel(level string) Level {

	level = strings.ToLower(level)

	switch level {
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	case "warn", "warning":
		return WarnLevel
	case "error":
		return ErrorLevel
	default:
		return InfoLevel
	}
}
