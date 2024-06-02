package botService

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/RazvanBerbece/AzteMarket/pkg/logging"
	"github.com/RazvanBerbece/AzteMarket/src/handlers/slashCmdEvents"
	logUtils "github.com/RazvanBerbece/AzteMarket/src/libs/services/logger/utils"
	sharedConfig "github.com/RazvanBerbece/AzteMarket/src/shared/config"
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
	if sharedConfig.DiscordMainGuildId != "" {
		// Register slash commands only for main guild
		err := slashCmdEvents.CreateAndRegisterSlashEventHandlers(b.Session, true, slashCmdEvents.DefinedSlashCommands)
		if err != nil {
			log.Fatal("Error registering slash commands for AzteBot: ", err)
		}
	} else {
		// Register slash commands for all guilds
		err := slashCmdEvents.CreateAndRegisterSlashEventHandlers(b.Session, false, slashCmdEvents.DefinedSlashCommands)
		if err != nil {
			log.Fatal("Error registering slash commands for AzteBot: ", err)
		}
	}

}

func (b *DiscordBotApplication) SetBotPermissions() {
	b.Session.Identify.Intents = discordgo.IntentsGuildMessages |
		discordgo.IntentsMessageContent
}

func (b *DiscordBotApplication) SetStateTracking() {
	b.Session.StateEnabled = true
	b.Session.State.TrackRoles = true
	b.Session.State.TrackMembers = true
	b.Session.State.TrackPresences = true
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
