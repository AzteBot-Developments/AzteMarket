package events

import "github.com/RazvanBerbece/AzteMarket/pkg/logging"

type LogEvent struct {
	Logger logging.Logger
	Type   string
	Msg    string
}
