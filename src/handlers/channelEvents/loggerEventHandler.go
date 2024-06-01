package channelEventsHandler

import (
	sharedRuntime "github.com/RazvanBerbece/AzteMarket/src/shared/runtime"
)

func HandleLoggerEvents() {
	for logEvent := range sharedRuntime.LogEventsChannel {
		switch logEvent.Type {
		case "INFO":
			go logEvent.Logger.LogInfo(logEvent.Msg)
		case "WARN":
			go logEvent.Logger.LogWarn(logEvent.Msg)
		case "ERROR":
			go logEvent.Logger.LogError(logEvent.Msg)
		}
	}
}
