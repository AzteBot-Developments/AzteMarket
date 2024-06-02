package channelEventsHandler

import (
	"fmt"

	botService "github.com/RazvanBerbece/AzteMarket/src/libs/services/bot"
	sharedRuntime "github.com/RazvanBerbece/AzteMarket/src/shared/runtime"
	"github.com/bwmarrin/discordgo"
)

func HandleComplexResponseEvents(s *discordgo.Session, responseMaxPageSize int) {
	for complexResponseEvent := range sharedRuntime.ComplexResponsesChannel {
		if complexResponseEvent.Embed != nil {
			// The event has an embed to passthrough
			if len(complexResponseEvent.Embed.Fields) > responseMaxPageSize && complexResponseEvent.PaginationRow != nil {
				// and the embed needs pagination !
				err := botService.ReplyComplexToInteraction(s, complexResponseEvent.Interaction, *complexResponseEvent.Embed, *complexResponseEvent.PaginationRow, responseMaxPageSize)
				if err != nil {
					fmt.Printf("Failed to process ComplexResponseEvent (Pagination: On): %v\n", err)
				}
			} else {
				editContent := ""
				editWebhook := discordgo.WebhookEdit{
					Content: &editContent,
					Embeds:  &[]*discordgo.MessageEmbed{complexResponseEvent.Embed.MessageEmbed},
				}
				s.InteractionResponseEdit(complexResponseEvent.Interaction, &editWebhook)
			}
		} else {
			if complexResponseEvent.Text != nil && complexResponseEvent.Title != nil {
				complexResponseEvent.Embed.Fields[0].Name = *complexResponseEvent.Title
				complexResponseEvent.Embed.Fields[0].Value = *complexResponseEvent.Text
				editContent := ""
				editWebhook := discordgo.WebhookEdit{
					Content: &editContent,
					Embeds:  &[]*discordgo.MessageEmbed{complexResponseEvent.Embed.MessageEmbed},
				}
				s.InteractionResponseEdit(complexResponseEvent.Interaction, &editWebhook)
			} else {
				fmt.Println("This response event:", complexResponseEvent, "is not valid.")
			}
		}
	}
}
