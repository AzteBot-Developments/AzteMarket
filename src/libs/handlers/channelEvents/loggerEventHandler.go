package channelEventsHandler

import (
	globalRuntime "github.com/RazvanBerbece/AzteMarket/src/globals/runtime"
)

func HandleLoggerEvents() {
	for logEvent := range globalRuntime.LogEventsChannel {
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
