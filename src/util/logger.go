package util

import "fmt"

func Debug(format string, a ...interface{}) {
	fmt.Printf("[DEBUG] "+format, a...)
	fmt.Println()
}
