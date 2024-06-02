package actionEvent

import (
	actionEventEmbedPagination "github.com/RazvanBerbece/AzteMarket/src/handlers/remoteEvents/actionEvents/embedPagination"
	sharedRuntime "github.com/RazvanBerbece/AzteMarket/src/shared/runtime"
	"github.com/bwmarrin/discordgo"
)

func HandleMessageComponentInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {

	eventCustomId := i.MessageComponentData().CustomID

	// Future button event handlers should be added here by custom ID
	switch eventCustomId {
	case sharedRuntime.PreviousPageOnEmbedEventId:
		actionEventEmbedPagination.HandlePaginatePreviousOnEmbed(s, i)
	case sharedRuntime.NextPageOnEmbedEventId:
		actionEventEmbedPagination.HandlePaginateNextOnEmbed(s, i)
	}
}
