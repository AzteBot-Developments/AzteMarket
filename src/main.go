package main

import (
	globalConfig "github.com/RazvanBerbece/AzteMarket/src/globals/config"
	globalRuntime "github.com/RazvanBerbece/AzteMarket/src/globals/runtime"
	channelEventsHandler "github.com/RazvanBerbece/AzteMarket/src/libs/handlers/channelEvents"
	"github.com/RazvanBerbece/AzteMarket/src/libs/handlers/remoteEvents"
	botService "github.com/RazvanBerbece/AzteMarket/src/libs/services/bot"
)

func main() {

	var bot botService.DiscordBotApplication

	// Start handlers for gochannel events
	go channelEventsHandler.HandleLoggerEvents()

	// Configure and launch session with provided bot token from the Discord Developer Portal
	bot.Configure(botService.Context{
		GatewayAuthToken: globalConfig.DiscordBotToken,
	}, globalRuntime.ConsoleLogger)

	// Set intents, permissions and state tracking
	bot.SetBotPermissions()
	bot.SetStateTracking()

	// Add event handling
	bot.AddEventHandlers(globalRuntime.ConsoleLogger, remoteEvents.RemoteEventHandlersAsList())

	// Connect to Discord Gateway and keep the connection alive
	// in order to handle receiving and sending remote events
	bot.Connect(globalRuntime.ConsoleLogger)

	bot.Disconnect()

}
