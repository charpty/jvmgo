package util

import "fmt"

const (
	LOG_LEVEL_DEBUG = 10
	LOG_LEVEL_ERROR = 30
)

var LogLevel int = LOG_LEVEL_ERROR

func Debug(format string, a ...interface{}) {
	if LogLevel <= LOG_LEVEL_DEBUG {
		fmt.Printf("[DEBUG] "+format, a...)
		fmt.Println()
	}
}

func Error(format string, a ...interface{}) {
	fmt.Printf("[ERROR] "+format, a...)
	fmt.Println()
}
