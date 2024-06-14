package slashCmdWalletHandlers

import (
	"fmt"
	"strings"

	"github.com/RazvanBerbece/AzteMarket/pkg/interaction"
	"github.com/RazvanBerbece/AzteMarket/pkg/utils"
	logUtils "github.com/RazvanBerbece/AzteMarket/src/libs/services/logger/utils"
	sharedConfig "github.com/RazvanBerbece/AzteMarket/src/shared/config"
	sharedRuntime "github.com/RazvanBerbece/AzteMarket/src/shared/runtime"
	"github.com/bwmarrin/discordgo"
)

func HandleSlashSendFundsFromWallet(s *discordgo.Session, i *discordgo.InteractionCreate) {

	receiverId := i.ApplicationCommandData().Options[0].StringValue()
	funds := i.ApplicationCommandData().Options[1].StringValue()

	var ref string = ""
	if len(i.ApplicationCommandData().Options) > 2 {
		ref = i.ApplicationCommandData().Options[2].StringValue()
	}

	fFunds, err := utils.StringToFloat64(funds)
	if err != nil {
		interaction.SendErrorEmbedResponse(s, i.Interaction, err.Error())
		return
	}
	if *fFunds <= 0 || *fFunds > 50000 {
		interaction.SendErrorEmbedResponse(s, i.Interaction, fmt.Sprintf("Invalid input argument (term: `%s`)", i.ApplicationCommandData().Options[1].Name))
		return
	}

	authorUserId := i.Member.User.ID // sender

	var receiverUserId string = ""
	var receiverWalletId string = ""
	if strings.Contains(receiverId, "@OTA") {
		// Possibly a wallet ID
		receiverWalletId = receiverId
	} else {
		// Possibly a user ID
		receiverUserId = utils.GetDiscordIdFromMentionFormat(receiverId)
	}

	var updatedFunds float64

	if receiverUserId != "" && receiverWalletId == "" {
		// Send by user ID
		updatedFunds, err = sharedRuntime.WalletService.SendFundsToUser(s, authorUserId, receiverUserId, *fFunds, ref)
		if err != nil {
			interaction.SendErrorEmbedResponse(s, i.Interaction, err.Error())
			go logUtils.PublishDiscordLogErrorEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, err.Error())
			return
		}
	} else if receiverWalletId != "" && receiverUserId == "" {
		// send by wallet ID
		updatedFunds, err = sharedRuntime.WalletService.SendFunds(s, authorUserId, receiverWalletId, *fFunds, ref)
		if err != nil {
			interaction.SendErrorEmbedResponse(s, i.Interaction, err.Error())
			go logUtils.PublishDiscordLogErrorEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, err.Error())
			return
		}
	}

	senderWallet, err := sharedRuntime.WalletService.GetWalletForUser(authorUserId)
	if err != nil {
		interaction.SendErrorEmbedResponse(s, i.Interaction, err.Error())
		go logUtils.PublishDiscordLogErrorEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, err.Error())
		return
	}

	// Audit transfer in the ledger
	go logUtils.PublishDiscordLogInfoEvent(sharedRuntime.LogEventsChannel, s, "Ledger", sharedConfig.DiscordChannelTopicPairs, fmt.Sprintf("`%s` sent `🪙 %.2f` AzteCoins to wallet `%s`", senderWallet.Id, *fFunds, receiverWalletId))

	if receiverUserId != "" && receiverWalletId == "" {
		interaction.SendSimpleEmbedSlashResponse(s, i.Interaction, fmt.Sprintf("Successfully transferred `%.2f` AzteCoins to `%s`. Your new balance is `🪙 %.2f` AzteCoins.", *fFunds, receiverUserId, updatedFunds))
	} else if receiverWalletId != "" && receiverUserId == "" {
		interaction.SendSimpleEmbedSlashResponse(s, i.Interaction, fmt.Sprintf("Successfully transferred `%.2f` AzteCoins to `%s`. Your new balance is `🪙 %.2f` AzteCoins.", *fFunds, receiverWalletId, updatedFunds))
	}
}
