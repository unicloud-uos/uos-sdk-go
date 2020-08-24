package log

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

type LogLevel int

const (
	LogOffLevel   LogLevel = 0 // Off log
	LogErrorLevel LogLevel = 1 // Errors should be properly handled
	LogWarnLevel  LogLevel = 2 // Info could be ignored; messages might need noticed
	LogInfoLevel  LogLevel = 3 // Informational messages
	LogDebugLevel LogLevel = 4 // Debug messages
)

type Logger struct {
	level  LogLevel
	logger *log.Logger
}

var logFlags = log.Ldate | log.Ltime | log.Lmicroseconds

func NewLogger(logLevel LogLevel) *Logger {
	l := Logger{
		level:  logLevel,
		logger: log.New(os.Stdout, "", logFlags),
	}
	return &l
}

func getCaller(skipCallDepth int) string {
	_, fullPath, line, ok := runtime.Caller(skipCallDepth)
	if !ok {
		return ""
	}
	fileParts := strings.Split(fullPath, "/")
	file := fileParts[len(fileParts)-1]
	return fmt.Sprintf("%s:%d", file, line)
}


func (l Logger) prefixArray() []interface{} {
	array := make([]interface{}, 0, 3)
	array = append(array, getCaller(3))
	return array
}

func (l Logger) Info(args ...interface{}) {
	if l.level < LogInfoLevel {
		return
	}
	prefixArray := l.prefixArray()
	prefixArray = append(prefixArray, "[INFO]")
	args = append(prefixArray, args...)
	l.logger.Println(args...)
}

func (l Logger) Warn(args ...interface{}) {
	if l.level < LogWarnLevel {
		return
	}
	prefixArray := l.prefixArray()
	prefixArray = append(prefixArray, "[WARN]")
	args = append(prefixArray, args...)
	l.logger.Println(args...)
}

func (l Logger) Error(args ...interface{}) {
	if l.level < LogErrorLevel {
		return
	}
	prefixArray := l.prefixArray()
	prefixArray = append(prefixArray, "[ERROR]")
	args = append(prefixArray, args...)
	l.logger.Println(args...)
}

func (l Logger) Debug(args ...interface{}) {
	if l.level < LogDebugLevel {
		return
	}
	prefixArray := l.prefixArray()
	prefixArray = append(prefixArray, "[DEBUG]")
	args = append(prefixArray, args...)
	l.logger.Println(args...)
}
