import (
	"errors"
	"log"
	"os"
)

const (
	LogNone = iota
	LogInfo
	LogWarning
	LogError
	LogDebug
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

func (opsgenieFileLogger *OpsgenieFileLogger) Log(level int, msg string) error {
	if opsgenieFileLogger.Logger == nil {
		return errors.New("FileLogger is not initialized correctly")
	}
	if level >= opsgenieFileLogger.LogLevel {
		opsgenieFileLogger.Logger.SetPrefix(levelPrefix[level])
		opsgenieFileLogger.Logger.Println(msg)
	}
	return nil
}