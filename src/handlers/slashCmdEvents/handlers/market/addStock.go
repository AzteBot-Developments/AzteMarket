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
	availableItems := i.ApplicationCommandData().Options[3].StringValue()

	// Input validation
	fCost, err := utils.StringToFloat64(stockCost)
	if err != nil {
		interaction.SendErrorEmbedResponse(s, i.Interaction, err.Error())
		return
	}
	if len(stockName) <= 0 || len(stockName) > 255 {
		interaction.SendErrorEmbedResponse(s, i.Interaction, fmt.Sprintf("Invalid input argument (term: `%s`)", i.ApplicationCommandData().Options[0].Name))
		return
	}
	if len(stockDetails) <= 0 || len(stockDetails) > 255 {
		interaction.SendErrorEmbedResponse(s, i.Interaction, fmt.Sprintf("Invalid input argument (term: `%s`)", i.ApplicationCommandData().Options[1].Name))
		return
	}
	if *fCost < 0 || *fCost > +1.7e+308 {
		interaction.SendErrorEmbedResponse(s, i.Interaction, fmt.Sprintf("Invalid input argument (term: `%s`)", i.ApplicationCommandData().Options[2].Name))
		return
	}
	availableItemsCount, err := utils.StringToInt(availableItems)
	if err != nil {
		interaction.SendErrorEmbedResponse(s, i.Interaction, err.Error())
		return
	}
	if *availableItemsCount < 0 || *availableItemsCount > 500000 {
		interaction.SendErrorEmbedResponse(s, i.Interaction, fmt.Sprintf("Invalid input argument (term: `%s`)", i.ApplicationCommandData().Options[3].Name))
		return
	}

	itemId, err := sharedRuntime.MarketplaceService.AddItemForSaleOnMarket(stockName, stockDetails, *fCost, *availableItemsCount)
	if err != nil {
		interaction.SendErrorEmbedResponse(s, i.Interaction, err.Error())
		go logUtils.PublishDiscordLogErrorEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, err.Error())
		return
	}

	log := fmt.Sprintf("A new item [id: `%s`]\n(Name: `%s` | Cost: `%s` | Available Items: `%d`)\nhas been added to the OTA marketplace", *itemId, stockName, stockCost, *availableItemsCount)
	go logUtils.PublishDiscordLogInfoEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, log)

	// Final response to interaction
	interaction.SendSimpleEmbedSlashResponse(s, i.Interaction, log)

}
