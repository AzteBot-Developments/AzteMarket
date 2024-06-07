package slashCmdMarketHandlers

import (
	"fmt"

	"github.com/RazvanBerbece/AzteMarket/pkg/interaction"
	logUtils "github.com/RazvanBerbece/AzteMarket/src/libs/services/logger/utils"
	sharedConfig "github.com/RazvanBerbece/AzteMarket/src/shared/config"
	sharedRuntime "github.com/RazvanBerbece/AzteMarket/src/shared/runtime"
	"github.com/bwmarrin/discordgo"
)

func HandleSlashRemoveItemFromMarket(s *discordgo.Session, i *discordgo.InteractionCreate) {

	// Retrieve input args from command
	itemId := i.ApplicationCommandData().Options[0].StringValue()

	itemToDelete, err := sharedRuntime.MarketplaceService.GetItemFromMarket(itemId)
	if err != nil {
		interaction.SendErrorEmbedResponse(s, i.Interaction, err.Error())
		go logUtils.PublishDiscordLogErrorEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, err.Error())
		return
	}

	err = sharedRuntime.MarketplaceService.RemoveItemFromMarket(itemId)
	if err != nil {
		interaction.SendErrorEmbedResponse(s, i.Interaction, err.Error())
		go logUtils.PublishDiscordLogErrorEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, err.Error())
		return
	}

	// Audit market update
	marketLog := fmt.Sprintf("An item (`%s` [`%s`]) was removed from the AzteMarket by `%s`", itemToDelete.DisplayName, itemToDelete.Id, i.Member.User.Username)
	go logUtils.PublishDiscordLogInfoEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, marketLog)

	// Final response to interaction
	interaction.SendSimpleEmbedSlashResponse(s, i.Interaction, fmt.Sprintf("Removed item with ID `%s` (`%s`) from the AzteMarket!", itemId, itemToDelete.DisplayName))

}
