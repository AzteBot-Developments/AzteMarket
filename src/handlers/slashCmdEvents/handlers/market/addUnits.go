package slashCmdMarketHandlers

import (
	"fmt"

	"github.com/RazvanBerbece/AzteMarket/pkg/interaction"
	logUtils "github.com/RazvanBerbece/AzteMarket/src/libs/services/logger/utils"
	sharedConfig "github.com/RazvanBerbece/AzteMarket/src/shared/config"
	sharedRuntime "github.com/RazvanBerbece/AzteMarket/src/shared/runtime"
	"github.com/bwmarrin/discordgo"
)

func HandleSlashAddAvailableUnitsForItem(s *discordgo.Session, i *discordgo.InteractionCreate) {

	itemId := i.ApplicationCommandData().Options[0].StringValue()
	multiplier := i.ApplicationCommandData().Options[1].IntValue()

	err := sharedRuntime.MarketplaceService.AddStockUnitsForItem(itemId, int(multiplier))
	if err != nil {
		interaction.SendErrorEmbedResponse(s, i.Interaction, err.Error())
		go logUtils.PublishDiscordLogErrorEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, err.Error())
		return
	}

	marketLog := fmt.Sprintf("`%d` units were added as available for sale to item with ID `%s`", multiplier, itemId)
	go logUtils.PublishDiscordLogInfoEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, marketLog)

	interaction.SendSimpleEmbedSlashResponse(s, i.Interaction, fmt.Sprintf("Added `%d` new units for sale to item with ID `%s` !", multiplier, itemId))

}
