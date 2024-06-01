package remoteEvents

import (
	remoteOnMessageEvent "github.com/RazvanBerbece/AzteMarket/src/handlers/remoteEvents/messageEvent"
	remoteOnReadyEvent "github.com/RazvanBerbece/AzteMarket/src/handlers/remoteEvents/readyEvent"
)

func RemoteEventHandlersAsList() []interface{} {
	return []interface{}{
		// <---- On Ready ---->
		remoteOnReadyEvent.DefaultHandler,
		// <---- On Message Create ---->
		remoteOnMessageEvent.Ping, remoteOnMessageEvent.DbPing,
	}
}
