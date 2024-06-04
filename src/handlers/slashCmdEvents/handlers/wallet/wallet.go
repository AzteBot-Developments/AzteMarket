package slashCmdWalletHandlers

import (
	"database/sql"
	"fmt"

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

	embedToSend := embed.NewEmbed().
		SetTitle(fmt.Sprintf("💳	`%s`'s Wallet", user.DiscordTag)).
		SetColor(sharedConfig.EmbedColorCode).
		DecorateWithTimestampFooter("Mon, 02 Jan 2006 15:04:05 MST").
		AddField("🧾 ID", fmt.Sprintf("`%s`", wallet.Id), false).
		AddField("🪙 Available Funds", fmt.Sprintf("`%d` AzteCoins", wallet.Funds), false).
		AddField("🛍️ Inventory", wallet.Inventory, false)

	interaction.SendEmbedSlashResponse(s, i.Interaction, *embedToSend)
}
