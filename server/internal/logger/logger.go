package logger

import (
	"os"
)

// func new() *log.Logger {

// 	return log.New(
// 		os.Stdout,
// 		"[APP]",
// 		log.Ldate|log.Ltime|log.Lshortfile,
// 	)
// }

var log = log.New(
	os.Stdout,
	"[APP] ",
	log.Ldate|log.Ltime|log.Lshortfile,
)

func new() *log.Logger {
	return log
}
