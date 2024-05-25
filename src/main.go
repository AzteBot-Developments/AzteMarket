package main

import (
	globalConfig "github.com/RazvanBerbece/AzteMarket/src/globals/config"
	"github.com/RazvanBerbece/AzteMarket/src/libs/handlers/remoteEvents"
	botService "github.com/RazvanBerbece/AzteMarket/src/libs/services/bot"
	loggerService "github.com/RazvanBerbece/AzteMarket/src/libs/services/logger"
)

func main() {

	// Some startup dependencies
	var consoleLogger loggerService.ConsoleLogger

	var bot botService.DiscordBotApplication

	// Configure and launch session with provided bot token from the Discord Developer Portal
	bot.Configure(botService.Context{
		GatewayAuthToken: globalConfig.DiscordBotToken,
	}, consoleLogger)

	defer bot.Disconnect()

	// Add event handling
	bot.AddEventHandlers(consoleLogger, remoteEvents.RemoteEventHandlersAsList())

	// Set intents, permissions and state tracking
	bot.SetIntents()
	bot.SetPermissions()
	bot.SetStateTracking()

	// Connect to Discord Gateway
	bot.Connect(consoleLogger)

}
