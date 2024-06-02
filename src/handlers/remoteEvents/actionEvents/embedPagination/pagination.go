package actionEventEmbedPagination

import (
	"fmt"
	"log"

	"github.com/RazvanBerbece/AzteMarket/pkg/interaction"
	actionEventsUtils "github.com/RazvanBerbece/AzteMarket/src/handlers/remoteEvents/actionEvents/utils"
	sharedConfig "github.com/RazvanBerbece/AzteMarket/src/shared/config"
	sharedRuntime "github.com/RazvanBerbece/AzteMarket/src/shared/runtime"
	"github.com/bwmarrin/discordgo"
)

func HandlePaginateNextOnEmbed(s *discordgo.Session, i *discordgo.InteractionCreate) {

	originalPaginatedEmbedId := i.Message.ID
	originalPaginatedEmbedChannelId := i.Message.ChannelID

	// Ensure that embed pagination can only be used by its creator
	// TODO

	// Respond to the button press (on help embed interaction source)
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Paginating...",
			Flags:   1 << 6, // ephemeral response
		},
	})
	if err != nil {
		log.Printf("Error responding to interaction: %v\n", err)
		interaction.SendErrorEmbedResponse(s, i.Interaction, err.Error())
	}
	interaction.DeleteInteractionResponse(s, i.Interaction, 0)

	// Get original interaction if it can be found in the in-memory map
	embedData, exists := sharedRuntime.EmbedsToPaginate[originalPaginatedEmbedId]
	if !exists {
		fmt.Println("Failed to paginate (Next)")
		return
	} else {
		go actionEventsUtils.UpdatePaginatedEmbedPage(s, &embedData, "NEXT", originalPaginatedEmbedChannelId, originalPaginatedEmbedId, sharedConfig.EmbedPageSize)
	}

}

func HandlePaginatePreviousOnEmbed(s *discordgo.Session, i *discordgo.InteractionCreate) {

	originalPaginatedEmbedId := i.Message.ID
	originalPaginatedEmbedChannelId := i.Message.ChannelID

	// Ensure that embed pagination can only be used by its creator
	// TODO

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Paginating...",
			Flags:   1 << 6,
		},
	})
	if err != nil {
		log.Printf("Error sending ACK message: %v", err)
		return
	}
	interaction.DeleteInteractionResponse(s, i.Interaction, 0)

	// Get original interaction if it can be found in the in-memory map
	embedData, exists := sharedRuntime.EmbedsToPaginate[originalPaginatedEmbedId]
	if !exists {
		fmt.Println("Failed to paginate (Previous)")
		return
	} else {
		go actionEventsUtils.UpdatePaginatedEmbedPage(s, &embedData, "PREV", originalPaginatedEmbedChannelId, originalPaginatedEmbedId, sharedConfig.EmbedPageSize)
	}

}
