package remoteOnReadyEvent

import (
	"fmt"

	"github.com/RazvanBerbece/AzteMarket/pkg/config"
	globalConfig "github.com/RazvanBerbece/AzteMarket/src/globals/config"
	globalRuntime "github.com/RazvanBerbece/AzteMarket/src/globals/runtime"
	"github.com/RazvanBerbece/AzteMarket/src/libs/models/events"
	loggerService "github.com/RazvanBerbece/AzteMarket/src/libs/services/logger"
	"github.com/bwmarrin/discordgo"
)

func DefaultHandler(s *discordgo.Session, event *discordgo.Ready) {

	log := fmt.Sprintf("`%s` is now online", event.User.Username)
	globalRuntime.LogEventsChannel <- events.LogEvent{
		Logger: loggerService.NewDiscordChannelLogger(s,
			config.GetDiscordChannelIdForTopic("Debug", globalConfig.DiscordChannelTopicPairs), true),
		Msg:  log,
		Type: "INFO",
	}

}
