package botService

import (
	"fmt"
	"time"

	"github.com/RazvanBerbece/AzteMarket/pkg/embed"
	"github.com/RazvanBerbece/AzteMarket/src/libs/models/domain"
	sharedRuntime "github.com/RazvanBerbece/AzteMarket/src/shared/runtime"
	"github.com/bwmarrin/discordgo"
)

func ReplyComplexToInteraction(s *discordgo.Session, i *discordgo.Interaction, embed embed.Embed, actionsRow discordgo.ActionsRow, pageSize int) error {

	originalAllFields := make([]*discordgo.MessageEmbedField, len(embed.Fields))
	copy(originalAllFields, embed.Fields)

	pages := (len(originalAllFields) + pageSize - 1) / pageSize
	if len(embed.Fields) < pageSize {
		_, err := s.ChannelMessageSendEmbed(i.ChannelID, embed.MessageEmbed)
		if err != nil {
			return err
		}
	} else {
		embed.Fields = embed.Fields[0:pageSize]
		embed.Footer = &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("Page 1 / %d", pages),
		}

		editContent := ""
		editWebhook := discordgo.WebhookEdit{
			Content:    &editContent,
			Embeds:     &[]*discordgo.MessageEmbed{embed.MessageEmbed},
			Components: &[]discordgo.MessageComponent{actionsRow},
		}
		msg, err := s.InteractionResponseEdit(i, &editWebhook)
		if err != nil {
			fmt.Printf("Error sending complex response to interaction: %v\n", err)
			return err
		}

		// Keep paginated embeds in-memory to enable handling on button presses
		sharedRuntime.EmbedsToPaginate[msg.ID] = domain.EmbedData{
			ChannelId:   msg.ChannelID,
			FieldData:   &originalAllFields, // all fields
			CurrentPage: 1,                  // same for all complex paginated embeds
			Timestamp:   float64(time.Now().Unix()),
		}
	}

	return nil
}
