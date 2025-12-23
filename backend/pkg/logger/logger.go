package logger

import (
	"log"
)

func Info(msg string, args ...any) {
	log.Printf("[INFO] "+msg, args...)
}

func Error(err error, msg string, args ...any) {
	log.Printf("[ERROR] "+msg+": %v", append(args, err)...)
}
