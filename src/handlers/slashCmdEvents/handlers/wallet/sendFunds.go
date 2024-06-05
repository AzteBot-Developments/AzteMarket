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

	id := i.ApplicationCommandData().Options[0].StringValue()
	funds := i.ApplicationCommandData().Options[1].StringValue()

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

	// This will be logged in the ledger
	var ledgerEntry string = ""
	var response string = ""

	// Figure out if user passed in a wallet or a user ID argument value
	receiverWalletId := ""
	receiverUserId := ""
	if strings.Contains(id, "@OTA") {
		receiverWalletId = id
		ledgerEntry = fmt.Sprintf("`%s` sent `🪙 %.2f` AzteCoins to wallet `%s`", authorUserId, *fFunds, receiverWalletId)
	} else {
		receiverUserId = utils.GetDiscordIdFromMentionFormat(id)
		ledgerEntry = fmt.Sprintf("`%s` sent `🪙 %.2f` AzteCoins to user `%s`", authorUserId, *fFunds, receiverUserId)
	}

	var updatedFunds float64
	var transferError error
	if len(receiverWalletId) > 0 {
		// send funds using the target wallet ID
		updatedFunds, transferError = sharedRuntime.WalletService.SendFunds(authorUserId, receiverWalletId, *fFunds)
		response = fmt.Sprintf("Successfully transferred `%.2f` AzteCoins to `%s`. Your new balance is `🪙 %.2f` AzteCoins.", *fFunds, receiverWalletId, updatedFunds)
	} else {
		// send funds using the target user ID
		updatedFunds, transferError = sharedRuntime.WalletService.SendFundsWithUserIds(authorUserId, receiverUserId, *fFunds)
		response = fmt.Sprintf("Successfully transferred `🪙 %.2f` AzteCoins to `%s`. Your new balance is `🪙 %.2f` AzteCoins.", *fFunds, receiverUserId, updatedFunds)
	}

	if transferError != nil {
		interaction.SendErrorEmbedResponse(s, i.Interaction, transferError.Error())
		go logUtils.PublishDiscordLogErrorEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, transferError.Error())
		return
	}

	// Audit transfer in the ledger
	go logUtils.PublishDiscordLogInfoEvent(sharedRuntime.LogEventsChannel, s, "Ledger", sharedConfig.DiscordChannelTopicPairs, ledgerEntry)

	interaction.SendSimpleEmbedSlashResponse(s, i.Interaction, response)
}
