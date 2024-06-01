package sharedRuntime

import "github.com/RazvanBerbece/AzteMarket/src/libs/models/events"

var LogEventsChannel = make(chan events.LogEvent)
