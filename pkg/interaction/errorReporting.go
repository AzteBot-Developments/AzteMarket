package interaction

import (
	"github.com/RazvanBerbece/AzteMarket/pkg/embed"
	"github.com/bwmarrin/discordgo"
)

func SendErrorEmbedResponse(s *discordgo.Session, i *discordgo.Interaction, errorMessage string) {

	embed := embed.NewEmbed().
		SetTitle("🤖❌   An Error Ocurred").
		SetThumbnail("https://i.postimg.cc/262tK7VW/148c9120-e0f0-4ed5-8965-eaa7c59cc9f2-2.jpg").
		SetColor(000000).
		AddField("Error Report", errorMessage, false)

	s.InteractionRespond(i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed.MessageEmbed},
		},
	})
}
