package slashCmdMarketHandlers

import (
	"fmt"

	"github.com/RazvanBerbece/AzteMarket/pkg/interaction"
	logUtils "github.com/RazvanBerbece/AzteMarket/src/libs/services/logger/utils"
	sharedConfig "github.com/RazvanBerbece/AzteMarket/src/shared/config"
	sharedRuntime "github.com/RazvanBerbece/AzteMarket/src/shared/runtime"
	"github.com/bwmarrin/discordgo"
)

func HandleSlashRemoveAvailableUnitsForItem(s *discordgo.Session, i *discordgo.InteractionCreate) {

	itemId := i.ApplicationCommandData().Options[0].StringValue()
	multiplier := i.ApplicationCommandData().Options[1].IntValue()

	err := sharedRuntime.MarketplaceService.RemoveStockUnitsForItem(itemId, int(multiplier))
	if err != nil {
		interaction.SendErrorEmbedResponse(s, i.Interaction, err.Error())
		go logUtils.PublishDiscordLogErrorEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, err.Error())
		return
	}

	marketLog := fmt.Sprintf("`%d` units were removed from sale of item with ID `%s`", multiplier, itemId)
	go logUtils.PublishDiscordLogInfoEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, marketLog)

	interaction.SendSimpleEmbedSlashResponse(s, i.Interaction, fmt.Sprintf("Removed `%d` units for sale from item with ID `%s`.", multiplier, itemId))

}
