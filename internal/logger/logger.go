package logger

import (
	"log"
	"os"
	"time"
)

// Custom loggers for different log levels
var (
	infoLogger  = log.New(os.Stdout, "INFO: ", 0)
	warnLogger  = log.New(os.Stdout, "WARN: ", 0)
	fatalLogger = log.New(os.Stderr, "FATAL: ", 0)
)

// Info logs a general informational message.
func Info(message string) {
	infoLogger.Printf("%s %s\n", time.Now().Format(time.RFC3339), message)
}

// Warn logs a warning message with additional context.
func Warn(err error, context string) {
	if err != nil {
		warnLogger.Printf("%s [%s]: %v\n", time.Now().Format(time.RFC3339), context, err)
	}
}

// Fatal logs a critical error with additional context, then terminates the program.
func Fatal(err error, context string) {
	if err != nil {
		fatalLogger.Fatalf("%s [%s]: %v\n", time.Now().Format(time.RFC3339), context, err)
		// os.Exit(1) is implicit in log.Fatalf
	}
}
