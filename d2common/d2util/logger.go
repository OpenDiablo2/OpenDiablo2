package d2util

import (
	"fmt"
	"io"
	"log"
)

// LogLevel determines how verbose the logging is (higher is more verbose)
type LogLevel int

// Log levels
const (
	LogLevelNone LogLevel = iota
	LogLevelError
	LogLevelWarning
	LogLevelInfo
	LogLevelDebug
)

// Log format strings for log levels
const (
	LogFmtDebug   = "[DEBUG] %s\n\r"
	LogFmtInfo    = "[INFO] %s\n\r"
	LogFmtWarning = "[WARNING] %s\n\r"
	LogFmtError   = "[ERROR] %s\n\r"
)

// Logger is used to write log messages, and can have a log level to determine verbosity
type Logger struct {
	io.Writer
	level LogLevel
}

// SetLevel sets the log level
func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

// Debug logs a debug message
func (l *Logger) Debug(msg string) {
	if l == nil {
		return
	}

	l.print(LogLevelDebug, msg)
}

// Info logs an info message
func (l *Logger) Info(msg string) {
	if l == nil {
		return
	}

	l.print(LogLevelInfo, msg)
}

// Warning logs a warning message
func (l *Logger) Warning(msg string) {
	if l == nil {
		return
	}

	l.print(LogLevelWarning, msg)
}

// Error logs an error message
func (l *Logger) Error(msg string) {
	if l == nil {
		return
	}

	l.print(LogLevelError, msg)
}

func (l *Logger) print(level LogLevel, msg string) {
	if l == nil || l.level < level {
		return
	}

	fmtString := ""

	switch level {
	case LogLevelDebug:
		fmtString = LogFmtDebug
	case LogLevelInfo:
		fmtString = LogFmtInfo
	case LogLevelWarning:
		fmtString = LogFmtWarning
	case LogLevelError:
		fmtString = LogFmtError
	case LogLevelNone:
	default:
		return
	}

	_, err := l.Write(format(fmtString, []byte(msg)))
	if err != nil {
		log.Print(err)
	}
}

func format(fmtStr string, fmtInput []byte) []byte {
	return []byte(fmt.Sprintf(fmtStr, string(fmtInput)))
}
