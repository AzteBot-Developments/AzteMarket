package remoteEvents

import remoteOnReadyEvent "github.com/RazvanBerbece/AzteMarket/src/libs/handlers/remoteEvents/readyEvent"

func RemoteEventHandlersAsList() []interface{} {
	return []interface{}{
		// <---- On Ready ---->
		remoteOnReadyEvent.DefaultHandler,
	}
}
