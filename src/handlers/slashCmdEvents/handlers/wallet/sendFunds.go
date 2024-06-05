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

	receiverWalletId := i.ApplicationCommandData().Options[0].StringValue()
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

	// Block attempts to send funds when the target ID is not a wallet ID
	if !strings.Contains(receiverWalletId, "@OTA") {
		interaction.SendErrorEmbedResponse(s, i.Interaction, fmt.Sprintf("Invalid input argument (term: `%s`)", i.ApplicationCommandData().Options[0].Name))
		return
	}

	updatedFunds, transferError := sharedRuntime.WalletService.SendFunds(s, authorUserId, receiverWalletId, *fFunds)
	if transferError != nil {
		interaction.SendErrorEmbedResponse(s, i.Interaction, transferError.Error())
		go logUtils.PublishDiscordLogErrorEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, transferError.Error())
		return
	}

	// Audit transfer in the ledger
	go logUtils.PublishDiscordLogInfoEvent(sharedRuntime.LogEventsChannel, s, "Ledger", sharedConfig.DiscordChannelTopicPairs, fmt.Sprintf("`%s` sent `🪙 %.2f` AzteCoins to wallet `%s`", authorUserId, *fFunds, receiverWalletId))

	interaction.SendSimpleEmbedSlashResponse(s, i.Interaction, fmt.Sprintf("Successfully transferred `%.2f` AzteCoins to `%s`. Your new balance is `🪙 %.2f` AzteCoins.", *fFunds, receiverWalletId, updatedFunds))
}
