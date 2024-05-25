package botService

import "github.com/RazvanBerbece/AzteMarket/pkg/logging"

type BotApplication interface {
	Configure(ctx Context, logger logging.Logger)
	Connect(logger logging.Logger)
	Disconnect()
	AddEventHandlers(remoteEventHandlers []interface{})
	SetBotPermissions()
	SetStateTracking()
}
