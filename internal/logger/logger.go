// logger/logger.go
package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Info logs a general informational message with a timestamp.
func Info(message string) {
	fmt.Printf("%s: %s\n", time.Now().Format(time.RFC3339), message)
}

// Warn logs a warning message with additional context and a timestamp.
func Warn(err error, context string) {
	if err != nil {
		fmt.Printf("%s - Warning [%s]: %v\n", time.Now().Format(time.RFC3339), context, err)
	}
}

// Fatal logs a critical error with additional context and a timestamp, then terminates the program.
func Fatal(err error, context string) {
	if err != nil {
		log.Fatalf("%s - Fatal [%s]: %v\n", time.Now().Format(time.RFC3339), context, err)
		// Terminate the program with a non-zero exit code
		os.Exit(1)
	}
}
