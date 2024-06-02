package slashCmdMarketHandlers

import (
	"fmt"

	"github.com/RazvanBerbece/AzteMarket/pkg/interaction"
	"github.com/RazvanBerbece/AzteMarket/pkg/utils"
	logUtils "github.com/RazvanBerbece/AzteMarket/src/libs/services/logger/utils"
	sharedConfig "github.com/RazvanBerbece/AzteMarket/src/shared/config"
	sharedRuntime "github.com/RazvanBerbece/AzteMarket/src/shared/runtime"
	"github.com/bwmarrin/discordgo"
)

func HandleSlashAddStock(s *discordgo.Session, i *discordgo.InteractionCreate) {

	// Retrieve input args from command
	stockName := i.ApplicationCommandData().Options[0].StringValue()
	stockDetails := i.ApplicationCommandData().Options[1].StringValue()
	stockCost := i.ApplicationCommandData().Options[2].StringValue()

	// Input validation
	val, err := utils.StringToFloat64(stockCost)
	if err != nil {
		interaction.SendErrorEmbedResponse(s, i.Interaction, err.Error())
		return
	}

	err = sharedRuntime.MarketplaceService.AddItemForSaleOnMarket(stockName, stockDetails, *val)
	if err != nil {
		interaction.SendErrorEmbedResponse(s, i.Interaction, err.Error())
		go logUtils.PublishDiscordLogErrorEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, err.Error())
		return
	}

	log := fmt.Sprintf("A new item (Name: `%s`, Cost: `%s`) has been added to the OTA marketplace", stockName, stockCost)
	go logUtils.PublishDiscordLogInfoEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, log)

	// Final response to interaction
	interaction.SendSimpleEmbedSlashResponse(s, i.Interaction, log)

}
