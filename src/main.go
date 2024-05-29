package main

import (
	globalConfig "github.com/RazvanBerbece/AzteMarket/src/globals/config"
	channelEventsHandler "github.com/RazvanBerbece/AzteMarket/src/libs/handlers/channelEvents"
	"github.com/RazvanBerbece/AzteMarket/src/libs/handlers/remoteEvents"
	botService "github.com/RazvanBerbece/AzteMarket/src/libs/services/bot"
	loggerService "github.com/RazvanBerbece/AzteMarket/src/libs/services/logger"
)

func main() {

	var bot botService.DiscordBotApplication

	// Start handlers for gochannel events
	go channelEventsHandler.HandleLoggerEvents()

	// Configure and launch session with provided bot token from the Discord Developer Portal
	bot.Configure(botService.Context{
		GatewayAuthToken: globalConfig.DiscordBotToken,
	}, loggerService.NewConsoleLogger())

	// Set intents, permissions and state tracking
	bot.SetBotPermissions()
	bot.SetStateTracking()

	// Add event handling
	bot.AddEventHandlers(loggerService.NewConsoleLogger(), remoteEvents.RemoteEventHandlersAsList())

	// Connect to Discord Gateway and keep the connection alive
	// in order to handle receiving and sending remote events
	bot.Connect(loggerService.NewConsoleLogger())

	bot.Disconnect()

}
