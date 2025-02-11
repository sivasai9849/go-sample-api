package logger

import (
    "fmt"
    "log"
    "os"
    "time"
)

type Logger struct {
    infoLogger  *log.Logger
    errorLogger *log.Logger
}

func NewLogger() *Logger {
    return &Logger{
        infoLogger:  log.New(os.Stdout, "INFO: ", log.LstdFlags),
        errorLogger: log.New(os.Stderr, "ERROR: ", log.LstdFlags),
    }
}

func (l *Logger) Info(format string, v ...interface{}) {
    l.log(l.infoLogger, format, v...)
}

func (l *Logger) Error(format string, v ...interface{}) {
    l.log(l.errorLogger, format, v...)
}

func (l *Logger) Fatal(format string, v ...interface{}) {
    l.log(l.errorLogger, format, v...)
    os.Exit(1)
}

func (l *Logger) log(logger *log.Logger, format string, v ...interface{}) {
    msg := fmt.Sprintf(format, v...)
    logger.Printf("[%s] %s", time.Now().Format("2006-01-02 15:04:05"), msg)
}