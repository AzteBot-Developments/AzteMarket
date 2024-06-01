package botService

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/RazvanBerbece/AzteMarket/pkg/logging"
	logUtils "github.com/RazvanBerbece/AzteMarket/src/libs/services/logger/utils"
	sharedRuntime "github.com/RazvanBerbece/AzteMarket/src/shared/runtime"
	"github.com/bwmarrin/discordgo"
)

type DiscordBotApplication struct {
	Session *discordgo.Session
}

func (b *DiscordBotApplication) Configure(ctx Context, logger logging.Logger) {
	session, err := discordgo.New("Bot " + ctx.GatewayAuthToken)
	if err != nil {
		go logUtils.PublishConsoleLogErrorEvent(sharedRuntime.LogEventsChannel, fmt.Sprintf("Could not create a Discord bot app session: %v", err))
		return
	}
	b.Session = session
}

func (b *DiscordBotApplication) AddEventHandlers(logger logging.Logger, remoteEventHandlers []interface{}) {

	go logUtils.PublishConsoleLogInfoEvent(sharedRuntime.LogEventsChannel, fmt.Sprintf("Registering %d remote event handlers...", len(remoteEventHandlers)))

	// onMessage, onReady, onUpdate, etc..
	for _, handler := range remoteEventHandlers {
		b.Session.AddHandler(handler)
	}

	// /buy, /trade, /wallet, etc..
	// TODO

}

func (b *DiscordBotApplication) SetBotPermissions() {
	b.Session.Identify.Intents = discordgo.IntentsGuildMessages |
		discordgo.IntentsMessageContent
}

func (b *DiscordBotApplication) SetStateTracking() {
	b.Session.StateEnabled = true
	b.Session.State.TrackChannels = true
	b.Session.State.MaxMessageCount = 100
}

func (b *DiscordBotApplication) Connect(logger logging.Logger) {

	err := b.Session.Open()
	if err != nil {
		go logUtils.PublishConsoleLogErrorEvent(sharedRuntime.LogEventsChannel, fmt.Sprintf("Could not connect the bot to the Discord Gateway: %v", err))
		return
	}

	go logUtils.PublishConsoleLogInfoEvent(sharedRuntime.LogEventsChannel, "Discord bot session is now connected !")

	// wait here until CTRL-C or anther term signal is received
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

}

func (b *DiscordBotApplication) Disconnect() {
	b.Session.Close()
}
