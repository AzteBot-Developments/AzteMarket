package botService

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/RazvanBerbece/AzteMarket/pkg/logging"
	"github.com/bwmarrin/discordgo"
)

type DiscordBotApplication struct {
	Session *discordgo.Session
}

func (b *DiscordBotApplication) Configure(ctx Context, logger logging.Logger) {
	session, err := discordgo.New("Bot " + ctx.GatewayAuthToken)
	if err != nil {
		logger.LogError(fmt.Sprintf("Could not create a Discord bot app session: : %v", err))
		return
	}
	b.Session = session
}

func (b *DiscordBotApplication) AddEventHandlers(logger logging.Logger, remoteEventHandlers []interface{}) {

	// onMessage, onReady, onUpdate, etc..
	logger.LogInfo(fmt.Sprintf("Registering %d remote event handlers...\n", len(remoteEventHandlers)))
	for _, handler := range remoteEventHandlers {
		b.Session.AddHandler(handler)
	}

	// /buy, /trade, /wallet, etc..
	// TODO

}

func (b *DiscordBotApplication) SetIntents() {
	b.Session.Identify.Intents = discordgo.IntentsGuildMessages
}

func (b *DiscordBotApplication) SetPermissions() {
	b.Session.Identify.Intents = discordgo.PermissionManageMessages |
		discordgo.PermissionReadMessageHistory |
		discordgo.PermissionManageServer |
		discordgo.PermissionManageRoles |
		discordgo.PermissionManageChannels
}

func (b *DiscordBotApplication) SetStateTracking() {
	b.Session.StateEnabled = true
	b.Session.State.TrackChannels = true
	b.Session.State.MaxMessageCount = 100
}

func (b *DiscordBotApplication) Connect(logger logging.Logger) {

	err := b.Session.Open()
	if err != nil {
		logger.LogError(fmt.Sprintf("Could not connect the bot to the Discord Gateway: %v", err))
		return
	}

	// wait here until CTRL-C or anther term signal is received
	go logger.LogInfo("Discord bot session is now connected !")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

}

func (b *DiscordBotApplication) Disconnect() {
	b.Session.Close()
}
