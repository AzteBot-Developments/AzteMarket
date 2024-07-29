package slashCmdMarketHandlers

import (
	"fmt"

	"github.com/RazvanBerbece/AzteMarket/pkg/interaction"
	logUtils "github.com/RazvanBerbece/AzteMarket/src/libs/services/logger/utils"
	sharedConfig "github.com/RazvanBerbece/AzteMarket/src/shared/config"
	sharedRuntime "github.com/RazvanBerbece/AzteMarket/src/shared/runtime"
	"github.com/bwmarrin/discordgo"
)

func HandleSlashBuyItem(s *discordgo.Session, i *discordgo.InteractionCreate) {

	buyerId := i.Member.User.ID
	buyerTag := fmt.Sprintf("<@%s>", buyerId)

	// Retrieve input args from command
	itemId := i.ApplicationCommandData().Options[0].StringValue()

	// Input validation
	if len(itemId) <= 0 || len(itemId) > 40 {
		interaction.SendErrorEmbedResponse(s, i.Interaction, fmt.Sprintf("Invalid input argument (term: `%s`)", i.ApplicationCommandData().Options[0].Name))
		return
	}

	item, err := sharedRuntime.MarketplaceService.GetItemFromMarket(itemId)
	if err != nil {
		interaction.SendErrorEmbedResponse(s, i.Interaction, err.Error())
		return
	}

	buyerWallet, err := sharedRuntime.WalletService.GetWalletForUser(buyerId)
	if err != nil {
		interaction.SendErrorEmbedResponse(s, i.Interaction, err.Error())
		return
	}

	err = sharedRuntime.MarketplaceService.BuyItem(buyerId, itemId)
	if err != nil {
		interaction.SendErrorEmbedResponse(s, i.Interaction, err.Error())
		return
	}

	// Audit market update
	marketLog := fmt.Sprintf("`%s` was bought on the AzteMarket by `%s` (`%s`) for `🪙 %.2f` AzteCoins", item.DisplayName, i.Member.User.Username, buyerId, item.Cost)
	go logUtils.PublishDiscordLogInfoEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, marketLog)

	// Audit ledger update for the coin transfer
	ledgerLog := fmt.Sprintf("`%s` bought item `%s` [`%s`] for `🪙 %.2f` AzteCoins", buyerWallet.Id, item.DisplayName, item.Id, item.Cost)
	go logUtils.PublishDiscordLogInfoEvent(sharedRuntime.LogEventsChannel, s, "Ledger", sharedConfig.DiscordChannelTopicPairs, ledgerLog)

	// Audit item purchase in designated staff channel
	purchaseLog := fmt.Sprintf("%s bought item `%s` [`%s`] for `🪙 %.2f` AzteCoins.\n\nPlease ensure that they create a ticket and that their benefit is delivered !\n\n@everyone", buyerTag, item.DisplayName, item.Id, item.Cost)
	go logUtils.PublishDiscordLogInfoEvent(sharedRuntime.LogEventsChannel, s, "PurchaseAudit", sharedConfig.DiscordChannelTopicPairs, purchaseLog)

	// Final response to interaction
	interaction.SendSimpleEmbedSlashResponse(s, i.Interaction, fmt.Sprintf("Bought 1 x `%s` (ID: `%s`)", item.DisplayName, itemId))

}
