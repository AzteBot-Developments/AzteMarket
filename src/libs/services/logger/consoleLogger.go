package loggerService

import (
	"fmt"
)

type ConsoleLogger struct {
}

func NewConsoleLogger() *ConsoleLogger {
	return &ConsoleLogger{}
}

func (l ConsoleLogger) LogInfo(msg string) {
	log := fmt.Sprintf("INFO: %s", msg)
	fmt.Println(log)
}

func (l ConsoleLogger) LogWarn(msg string) {
	log := fmt.Sprintf("WARN: %s", msg)
	fmt.Println(log)
}

func (l ConsoleLogger) LogError(msg string) {
	log := fmt.Sprintf("ERROR: %s", msg)
	fmt.Println(log)
}
