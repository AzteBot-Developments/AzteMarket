package slashCmdWalletHandlers

import (
	"fmt"

	"github.com/RazvanBerbece/AzteMarket/pkg/interaction"
	logUtils "github.com/RazvanBerbece/AzteMarket/src/libs/services/logger/utils"
	sharedConfig "github.com/RazvanBerbece/AzteMarket/src/shared/config"
	sharedRuntime "github.com/RazvanBerbece/AzteMarket/src/shared/runtime"
	"github.com/bwmarrin/discordgo"
)

func HandleSlashDeleteWallet(s *discordgo.Session, i *discordgo.InteractionCreate) {

	authorUserId := i.Member.User.ID
	rowsAffected, err := sharedRuntime.WalletService.DeleteWalletForUser(authorUserId)
	if err != nil {
		interaction.SendErrorEmbedResponse(s, i.Interaction, err.Error())
		go logUtils.PublishDiscordLogErrorEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, err.Error())
		return
	}

	switch rowsAffected {
	case 0:
		interaction.SendSimpleEmbedSlashResponse(s, i.Interaction, "No Wallet was deleted as part of this slash command execution. It may be because the Wallet to delete does not exist anymore.")
	case 1:
		log := fmt.Sprintf("A wallet [userId: `%s`] was deleted", authorUserId)
		go logUtils.PublishDiscordLogInfoEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, log)
		interaction.SendSimpleEmbedSlashResponse(s, i.Interaction, "The Wallet deletion operation completed successfully. Any leftover funds and aquired items have been cleared out.")
	default:
		interaction.SendSimpleEmbedSlashResponse(s, i.Interaction, "The Wallet deletion operation completed successfully, but raised some warnings.")
		go logUtils.PublishDiscordLogWarnEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, "More than 1 Wallet with the same ID states to have been deleted")
	}

}
