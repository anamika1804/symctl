package logger

import "log"

var Debug = false

func Debugf(format string, args ...interface{}) {
	if Debug {
		log.Printf(format, args...)
	}
}
