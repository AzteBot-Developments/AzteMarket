package interaction

import (
	"fmt"
	"time"

	"github.com/RazvanBerbece/AzteMarket/pkg/embed"
	"github.com/bwmarrin/discordgo"
)

var ErrorColorCode = 15548997

func DeleteInteractionResponse(s *discordgo.Session, i *discordgo.Interaction, msDelay int) {

	time.Sleep(time.Duration(msDelay) * time.Millisecond)

	// Delete the response
	s.InteractionResponseDelete(i)
}

func SendErrorEmbedResponse(s *discordgo.Session, i *discordgo.Interaction, errorMessage string) {

	embed := embed.NewEmbed().
		SetTitle("🤖❌   An Error Ocurred").
		SetThumbnail("https://i.postimg.cc/262tK7VW/148c9120-e0f0-4ed5-8965-eaa7c59cc9f2-2.jpg").
		SetColor(ErrorColorCode).
		AddField("Error Report", errorMessage, false)

	s.InteractionRespond(i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed.MessageEmbed},
		},
	})
}

func ErrorEmbedResponseEdit(s *discordgo.Session, i *discordgo.Interaction, errorMessage string) {

	embed := embed.NewEmbed().
		SetTitle(fmt.Sprintf("🤖❌   `/%s` Command Execution Error", i.ApplicationCommandData().Name)).
		SetThumbnail("https://i.postimg.cc/262tK7VW/148c9120-e0f0-4ed5-8965-eaa7c59cc9f2-2.jpg").
		SetColor(ErrorColorCode).
		AddField("Error Report", errorMessage, false)

	editWebhook := discordgo.WebhookEdit{
		Embeds: &[]*discordgo.MessageEmbed{embed.MessageEmbed},
	}

	s.InteractionResponseEdit(i, &editWebhook)
}
