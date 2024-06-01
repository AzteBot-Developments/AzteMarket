package remoteOnMessageEvent

import (
	"fmt"

	logUtils "github.com/RazvanBerbece/AzteMarket/src/libs/services/logger/utils"
	sharedConfig "github.com/RazvanBerbece/AzteMarket/src/shared/config"
	sharedRuntime "github.com/RazvanBerbece/AzteMarket/src/shared/runtime"
	"github.com/bwmarrin/discordgo"
)

func DbPing(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Retrieves the author user data as a ping to the DB
	if m.Content == "dbping" {

		authorUserId := m.Author.ID

		user, err := sharedRuntime.UserService.GetUser(authorUserId)
		if err != nil {
			go logUtils.PublishDiscordLogErrorEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, err.Error())
			return
		}

		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Read data for user `%s` [`%s`] from the DB", user.DiscordTag, user.UserId))
	}

}
