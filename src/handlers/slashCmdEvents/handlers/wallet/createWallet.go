package slashCmdWalletHandlers

import (
	"fmt"

	"github.com/RazvanBerbece/AzteMarket/pkg/interaction"
	logUtils "github.com/RazvanBerbece/AzteMarket/src/libs/services/logger/utils"
	sharedConfig "github.com/RazvanBerbece/AzteMarket/src/shared/config"
	sharedRuntime "github.com/RazvanBerbece/AzteMarket/src/shared/runtime"
	"github.com/bwmarrin/discordgo"
	"github.com/go-sql-driver/mysql"
)

func HandleSlashCreateWallet(s *discordgo.Session, i *discordgo.InteractionCreate) {

	authorUserId := i.Member.User.ID
	wallet, err := sharedRuntime.WalletService.CreateWalletForUser(authorUserId)
	if err != nil {
		switch err.(*mysql.MySQLError).Number {
		case 1062:
			interaction.SendErrorEmbedResponse(s, i.Interaction, fmt.Sprintf("A server error ocurred while creating an AzteMarket Wallet for user with ID `%s`: This user already has a Wallet on the AzteMarket, so the slash command failed to process.", authorUserId))
			go logUtils.PublishDiscordLogErrorEvent(
				sharedRuntime.LogEventsChannel, s, "Debug",
				sharedConfig.DiscordChannelTopicPairs,
				err.Error(),
			)
		default:
			interaction.SendErrorEmbedResponse(s, i.Interaction, fmt.Sprintf("A server error ocurred while creating an AzteMarket Wallet for user with ID `%s`", authorUserId))
			go logUtils.PublishDiscordLogErrorEvent(
				sharedRuntime.LogEventsChannel, s, "Debug",
				sharedConfig.DiscordChannelTopicPairs,
				err.Error(),
			)
		}
		return
	}

	log := fmt.Sprintf("A new wallet [id: `%s`] was created", wallet.Id)
	go logUtils.PublishDiscordLogInfoEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, log)

	// Final response to interaction
	interaction.SendSimpleEmbedSlashResponse(s, i.Interaction, fmt.Sprintf("Your AzteMarket wallet has been created !\nYour assigned wallet ID is `%s`", wallet.Id))

}
