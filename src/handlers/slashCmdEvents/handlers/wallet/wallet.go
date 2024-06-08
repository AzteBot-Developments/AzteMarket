package slashCmdWalletHandlers

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/RazvanBerbece/AzteMarket/pkg/embed"
	"github.com/RazvanBerbece/AzteMarket/pkg/interaction"
	logUtils "github.com/RazvanBerbece/AzteMarket/src/libs/services/logger/utils"
	sharedConfig "github.com/RazvanBerbece/AzteMarket/src/shared/config"
	sharedRuntime "github.com/RazvanBerbece/AzteMarket/src/shared/runtime"
	"github.com/bwmarrin/discordgo"
)

func HandleSlashWallet(s *discordgo.Session, i *discordgo.InteractionCreate) {

	authorUserId := i.Member.User.ID
	wallet, err := sharedRuntime.WalletService.GetWalletForUser(authorUserId)
	if err != nil {

		// wallet doesn't exist, so customise the message
		if err == sql.ErrNoRows {
			interaction.SendErrorEmbedResponse(s, i.Interaction, "No wallet was found for your user ID. You can create a new wallet by using the `/wallet-create` slash command.")
			go logUtils.PublishDiscordLogErrorEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, fmt.Sprintf("User `%s` tried to retrieve a non existing wallet entry.", authorUserId))
			return
		}

		interaction.SendErrorEmbedResponse(s, i.Interaction, err.Error())
		go logUtils.PublishDiscordLogErrorEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, err.Error())
		return
	}

	user, err := sharedRuntime.UserService.GetUser(authorUserId)
	if err != nil {
		interaction.SendErrorEmbedResponse(s, i.Interaction, err.Error())
		go logUtils.PublishDiscordLogErrorEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, err.Error())
		return
	}

	// Build the inventory subcomponent to show in the embed
	// Holds all available item names in a bullet point list
	var inventoryDisplay string = ""
	inventoryString := wallet.Inventory
	itemIds := strings.Split(inventoryString, ",")
	for _, id := range itemIds {
		if len(id) == 0 {
			continue
		}
		item, err := sharedRuntime.MarketplaceService.GetItemFromMarket(id)
		if err != nil {
			go logUtils.PublishDiscordLogErrorEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, err.Error())
			continue
		}
		inventoryDisplay += fmt.Sprintf("- `%s` (worth `🪙 %.2f`)\n", item.DisplayName, item.Cost)
	}
	if len(inventoryDisplay) == 0 {
		inventoryDisplay = "_there's no items currently in this wallet._"
	}

	embedToSend := embed.NewEmbed().
		SetTitle(fmt.Sprintf("💳	`%s`'s Wallet", user.DiscordTag)).
		SetColor(sharedConfig.EmbedColorCode).
		DecorateWithTimestampFooter("Mon, 02 Jan 2006 15:04:05 MST").
		AddLineBreakField().
		AddField("🧾 ID", fmt.Sprintf("`%s`", wallet.Id), false).
		AddField("💰 Available Funds", fmt.Sprintf("`🪙 %.2f` AzteCoins", wallet.Funds), false).
		AddLineBreakField().
		AddField("🛍️ Inventory", inventoryDisplay, false)

	interaction.SendEmbedSlashResponse(s, i.Interaction, *embedToSend)
}
