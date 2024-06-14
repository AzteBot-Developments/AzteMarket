package slashCmdWalletHandlers

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/RazvanBerbece/AzteMarket/pkg/embed"
	"github.com/RazvanBerbece/AzteMarket/pkg/interaction"
	"github.com/RazvanBerbece/AzteMarket/pkg/utils"
	"github.com/RazvanBerbece/AzteMarket/src/libs/models/dax"
	logUtils "github.com/RazvanBerbece/AzteMarket/src/libs/services/logger/utils"
	sharedConfig "github.com/RazvanBerbece/AzteMarket/src/shared/config"
	sharedRuntime "github.com/RazvanBerbece/AzteMarket/src/shared/runtime"
	"github.com/bwmarrin/discordgo"
)

func HandleSlashViewWallet(s *discordgo.Session, i *discordgo.InteractionCreate) {

	targetId := i.ApplicationCommandData().Options[0].StringValue()

	authorUsername := i.Member.User.Username

	var wallet *dax.Wallet
	var targetUser *dax.User

	if strings.Contains(targetId, "@OTA") {
		// Possibly a wallet ID
		var walletErr error
		wallet, walletErr = sharedRuntime.WalletService.GetWallet(targetId)
		if walletErr != nil {

			// wallet doesn't exist, so customise the message
			if walletErr == sql.ErrNoRows {
				interaction.SendErrorEmbedResponse(s, i.Interaction, "No wallet was found for this wallet ID. They can create a new wallet by using the `/wallet-create` slash command.")
				go logUtils.PublishDiscordLogErrorEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, fmt.Sprintf("User `%s` tried to retrieve a non existing wallet entry (`%s`).", authorUsername, targetId))
				return
			}

			interaction.SendErrorEmbedResponse(s, i.Interaction, walletErr.Error())
			go logUtils.PublishDiscordLogErrorEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, walletErr.Error())
			return
		}

		var userErr error
		targetUser, userErr = sharedRuntime.UserService.GetUser(wallet.UserId)
		if userErr != nil {
			interaction.SendErrorEmbedResponse(s, i.Interaction, userErr.Error())
			go logUtils.PublishDiscordLogErrorEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, userErr.Error())
			return
		}
	} else {
		// Possibly a user ID
		sanitisedUserId := utils.GetDiscordIdFromMentionFormat(targetId)
		var walletErr error
		wallet, walletErr = sharedRuntime.WalletService.GetWalletForUser(sanitisedUserId)
		if walletErr != nil {

			// wallet doesn't exist, so customise the message
			if walletErr == sql.ErrNoRows {
				interaction.SendErrorEmbedResponse(s, i.Interaction, "No wallet was found for this user ID. They can create a new wallet by using the `/wallet-create` slash command.")
				go logUtils.PublishDiscordLogErrorEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, fmt.Sprintf("User `%s` tried to retrieve a non existing wallet entry (`%s`).", authorUsername, sanitisedUserId))
				return
			}

			interaction.SendErrorEmbedResponse(s, i.Interaction, walletErr.Error())
			go logUtils.PublishDiscordLogErrorEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, walletErr.Error())
			return
		}

		var userErr error
		targetUser, userErr = sharedRuntime.UserService.GetUser(sanitisedUserId)
		if userErr != nil {
			interaction.SendErrorEmbedResponse(s, i.Interaction, userErr.Error())
			go logUtils.PublishDiscordLogErrorEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, userErr.Error())
			return
		}
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
		SetTitle(fmt.Sprintf("💳	`%s`'s Wallet", targetUser.DiscordTag)).
		SetColor(sharedConfig.EmbedColorCode).
		DecorateWithTimestampFooter("Mon, 02 Jan 2006 15:04:05 MST").
		AddLineBreakField().
		AddField("🧾 ID", fmt.Sprintf("`%s`", wallet.Id), false).
		AddField("💰 Available Funds", fmt.Sprintf("`🪙 %.2f` AzteCoins", wallet.Funds), false).
		AddLineBreakField().
		AddField("🛍️ Inventory", inventoryDisplay, false)

	interaction.SendEmbedSlashResponse(s, i.Interaction, *embedToSend, true)
}
