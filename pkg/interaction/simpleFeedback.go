package interaction

import (
	"github.com/RazvanBerbece/AzteMarket/pkg/embed"
	"github.com/bwmarrin/discordgo"
)

func SendSimpleEmbedSlashResponse(s *discordgo.Session, i *discordgo.Interaction, msg string) {

	embed := embed.NewEmbed().
		SetTitle("🤖   Slash Command Result").
		SetColor(000000).
		SetDescription(msg)

	s.InteractionRespond(i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed.MessageEmbed},
		},
	})
}
