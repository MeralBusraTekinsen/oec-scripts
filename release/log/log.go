package main

import (
	"fmt"
	"log"
	"os"
)

const (
	LogNone = iota
	LogDebug
	LogInfo
	LogWarning
	LogError
)

var levelPrefix = map[int]string{LogInfo: "INFO ", LogDebug: "DEBUG ", LogWarning: "WARNING ", LogError: "ERROR "}

type OpsgenieFileLogger struct {
	Logger   *log.Logger
	LogFile  *os.File
	LogLevel int
}

func NewFileLogger() *OpsgenieFileLogger {
	return &OpsgenieFileLogger{
		Logger:   nil,
		LogFile:  nil,
		LogLevel: LogNone,
	}
}

func (opsgenieFileLogger *OpsgenieFileLogger) setOutput(file *os.File) {
	opsgenieFileLogger.Logger = log.New(file, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lmsgprefix)
	opsgenieFileLogger.LogFile = file
}

func (opsgenieFileLogger *OpsgenieFileLogger) log(level int, msg string) {
	if opsgenieFileLogger.Logger != nil {
		if level >= opsgenieFileLogger.LogLevel {
			opsgenieFileLogger.Logger.SetPrefix(levelPrefix[level])
			opsgenieFileLogger.Logger.Println(msg)
		}
	} else {
		fmt.Println("FileLogger is not initialized correctly")
	}
}

func (opsgenieFileLogger *OpsgenieFileLogger) Error(msg string) {
	opsgenieFileLogger.log(LogError, msg)
}

func (opsgenieFileLogger *OpsgenieFileLogger) Info(msg string) {
	opsgenieFileLogger.log(LogInfo, msg)
}

func (opsgenieFileLogger *OpsgenieFileLogger) Warning(msg string) {
	opsgenieFileLogger.log(LogWarning, msg)
}

func (opsgenieFileLogger *OpsgenieFileLogger) Debug(msg string) {
	opsgenieFileLogger.log(LogDebug, msg)
}