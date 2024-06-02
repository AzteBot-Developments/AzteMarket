package remoteOnReadyEvent

import (
	"fmt"

	channelEventsHandler "github.com/RazvanBerbece/AzteMarket/src/handlers/channelEvents"
	logUtils "github.com/RazvanBerbece/AzteMarket/src/libs/services/logger/utils"
	sharedConfig "github.com/RazvanBerbece/AzteMarket/src/shared/config"
	sharedRuntime "github.com/RazvanBerbece/AzteMarket/src/shared/runtime"
	"github.com/bwmarrin/discordgo"
)

func DefaultHandler(s *discordgo.Session, event *discordgo.Ready) {

	// Log ready event
	log := fmt.Sprintf("`%s` is now online", event.User.Username)
	go logUtils.PublishDiscordLogInfoEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, log)

	// Start gochannel event handlers which depend on the app session
	go channelEventsHandler.HandleComplexResponseEvents(s, sharedConfig.EmbedPageSize)

}
