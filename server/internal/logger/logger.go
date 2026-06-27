package logger

import (
	"log"
	"os"
)

// config hold logger information
type Config struct {
	Level string
}

func New(cfg Config) *log.Logger {

	return log.New(
		os.Stdout,
		"[APP]",
		log.Ldate|log.Ltime|log.Lshortfile,
	)
}

// var log = log.New(
// 	os.Stdout,
// 	"[APP] ",
// 	log.Ldate|log.Ltime|log.Lshortfile,
// )

// func new() *log.Logger {
// 	return log
// }
