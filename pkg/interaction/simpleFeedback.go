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

func SendEmbedSlashResponse(s *discordgo.Session, i *discordgo.Interaction, embed embed.Embed, ephemeral bool) {
	if ephemeral {
		s.InteractionRespond(i, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embed.MessageEmbed},
				Flags:  64,
			},
		})

		return
	}
	s.InteractionRespond(i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed.MessageEmbed},
		},
	})
}
