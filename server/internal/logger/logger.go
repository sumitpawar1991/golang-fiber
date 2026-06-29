package logger

import (
	"io"
	"log"
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
	Level string
}

type logger struct {
	mu     sync.Mutex
	level  Level
	output io.Writer
	fields map[string]interface{}
}

func New(cfg Config) *log.Logger {

	// return log.New(
	// 	os.Stdout,
	// 	"[APP]",
	// 	log.Ldate|log.Ltime|log.Lshortfile,
	// )

	level := parseLevel(cfg.Level)
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
