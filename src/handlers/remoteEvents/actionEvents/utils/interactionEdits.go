package actionEventsUtils

import (
	"fmt"

	"github.com/RazvanBerbece/AzteMarket/src/libs/models/domain"
	sharedRuntime "github.com/RazvanBerbece/AzteMarket/src/shared/runtime"
	"github.com/bwmarrin/discordgo"
)

func UpdatePaginatedEmbedPage(s *discordgo.Session, embedData *domain.EmbedData, opCode string, channelId string, messageId string, pageSize int) error {

	// Retrieve original embed (the one with the pagination action row)
	message, err := s.ChannelMessage(channelId, messageId)
	if err != nil {
		return err
	}

	if len(message.Embeds) > 0 {

		pages := (len(*embedData.FieldData) + pageSize - 1) / pageSize

		// Calculate next page (and wrap if necessary)
		currentPage := embedData.CurrentPage
		if opCode == "NEXT" {
			currentPage += 1
		} else if opCode == "PREV" {
			currentPage -= 1
		}
		if currentPage > pages {
			currentPage = 1
		} else if currentPage < 1 {
			currentPage = pages
		}

		// Update map to hold new page number
		// assume that key exists
		sharedRuntime.EmbedsToPaginate[messageId] = domain.EmbedData{
			ChannelId:   embedData.ChannelId,
			CurrentPage: currentPage,
			FieldData:   embedData.FieldData,
			Timestamp:   embedData.Timestamp,
		}

		originalEmbed := message.Embeds[0] // this gets mutated

		// Determine the start and end index of fields to display for the current page
		startIdx := (currentPage - 1) * pageSize
		endIdx := startIdx + pageSize
		if endIdx > len(*embedData.FieldData) {
			endIdx = len(*embedData.FieldData)
		}

		fields := *embedData.FieldData
		paginatedFields := fields[startIdx:endIdx]
		originalEmbed.Fields = paginatedFields
		originalEmbed.Footer = &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("Page %d / %d", currentPage, pages),
		}
		interactionEdit := &discordgo.MessageEdit{
			Channel: channelId,
			ID:      messageId,
			Content: nil,
			Embeds:  &[]*discordgo.MessageEmbed{originalEmbed},
			Components: &[]discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Emoji: &discordgo.ComponentEmoji{
								Name: "⬅️",
							},
							Label:    "Previous",
							Style:    discordgo.PrimaryButton,
							CustomID: sharedRuntime.PreviousPageOnEmbedEventId,
							Disabled: false,
						},
						discordgo.Button{
							Emoji: &discordgo.ComponentEmoji{
								Name: "➡️",
							},
							Label:    "Next",
							Style:    discordgo.PrimaryButton,
							CustomID: sharedRuntime.NextPageOnEmbedEventId,
							Disabled: false,
						},
					},
				},
			},
		}

		_, err = s.ChannelMessageEditComplex(interactionEdit)
		if err != nil {
			// Handle error
			return err
		}
	}

	return nil
}

func DisablePaginatedEmbed(s *discordgo.Session, channelId string, messageId string) error {

	// Retrieve original embed (the one with the pagination action row)
	message, err := s.ChannelMessage(channelId, messageId)
	if err != nil {
		return err
	}

	originalEmbed := message.Embeds[0]

	interactionEdit := &discordgo.MessageEdit{
		Channel: channelId,
		ID:      messageId,
		Content: nil,
		Embeds:  &[]*discordgo.MessageEmbed{originalEmbed},
		Components: &[]discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Emoji: &discordgo.ComponentEmoji{
							Name: "⬅️",
						},
						Label:    "Previous",
						Style:    discordgo.PrimaryButton,
						CustomID: sharedRuntime.PreviousPageOnEmbedEventId,
						Disabled: true,
					},
					discordgo.Button{
						Emoji: &discordgo.ComponentEmoji{
							Name: "➡️",
						},
						Label:    "Next",
						Style:    discordgo.PrimaryButton,
						CustomID: sharedRuntime.NextPageOnEmbedEventId,
						Disabled: true,
					},
				},
			},
		},
	}

	_, err = s.ChannelMessageEditComplex(interactionEdit)
	if err != nil {
		// Handle error
		return err
	}

	return nil
}
