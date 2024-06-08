package loggerService

import (
	"fmt"

	"github.com/RazvanBerbece/AzteMarket/pkg/embed"
	"github.com/RazvanBerbece/AzteMarket/pkg/interaction"
	"github.com/bwmarrin/discordgo"
)

type DiscordChannelLogger struct {
	session            *discordgo.Session
	topic              string
	addTimestampFooter bool
}

func NewDiscordChannelLogger(s *discordgo.Session, channelId string, addTimestampFooter bool) *DiscordChannelLogger {
	return &DiscordChannelLogger{
		session:            s,
		topic:              channelId,
		addTimestampFooter: addTimestampFooter,
	}
}

func (l DiscordChannelLogger) LogInfo(msg string) {

	log := fmt.Sprintf("ℹ️ INFO: %s", msg)

	embed := embed.NewEmbed().
		SetColor(interaction.InfofEmbedColorCode)

	embed.AddField("", log, false)

	if l.addTimestampFooter {
		embed.DecorateWithTimestampFooter("Mon, 02 Jan 2006 15:04:05 MST")
	}

	_, err := l.session.ChannelMessageSendEmbed(l.topic, embed.MessageEmbed)
	if err != nil {
		fmt.Printf("Error sending log to channel %s: %v", l.topic, err)
	}
}

func (l DiscordChannelLogger) LogWarn(msg string) {

	log := fmt.Sprintf("⚠️ WARN: %s", msg)

	embed := embed.NewEmbed().
		SetColor(interaction.WarnEmbedColorCode)

	embed.AddField("", log, false)

	if l.addTimestampFooter {
		embed.DecorateWithTimestampFooter("Mon, 02 Jan 2006 15:04:05 MST")
	}

	_, err := l.session.ChannelMessageSendEmbed(l.topic, embed.MessageEmbed)
	if err != nil {
		fmt.Printf("Error sending log to channel %s: %v", l.topic, err)
	}
}

func (l DiscordChannelLogger) LogError(msg string) {

	log := fmt.Sprintf("⛔ ERROR: %s", msg)

	embed := embed.NewEmbed().
		SetColor(interaction.ErrorEmbedColorCode)

	embed.AddField("", log, false)

	if l.addTimestampFooter {
		embed.DecorateWithTimestampFooter("Mon, 02 Jan 2006 15:04:05 MST")
	}

	_, err := l.session.ChannelMessageSendEmbed(l.topic, embed.MessageEmbed)
	if err != nil {
		fmt.Printf("Error sending log to channel %s: %v", l.topic, err)
	}
}
