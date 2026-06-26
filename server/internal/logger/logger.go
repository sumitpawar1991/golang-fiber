package logger

import (
	"log"
	"os"
)

func new() *log.Logger {

	return log.New(
		os.Stdout,
		"[APP]",
		log.Ldate|log.Ltime|log.Lshortfile,
	)
}
