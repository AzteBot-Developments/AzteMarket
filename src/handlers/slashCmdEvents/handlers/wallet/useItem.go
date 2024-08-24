package slashCmdWalletHandlers

import (
	"fmt"
	"strings"

	"github.com/RazvanBerbece/AzteMarket/pkg/interaction"
	"github.com/RazvanBerbece/AzteMarket/pkg/utils"
	"github.com/RazvanBerbece/AzteMarket/src/libs/models/dax"
	logUtils "github.com/RazvanBerbece/AzteMarket/src/libs/services/logger/utils"
	sharedConfig "github.com/RazvanBerbece/AzteMarket/src/shared/config"
	sharedRuntime "github.com/RazvanBerbece/AzteMarket/src/shared/runtime"
	"github.com/bwmarrin/discordgo"
)

func HandleSlashUseItemFromWallet(s *discordgo.Session, i *discordgo.InteractionCreate) {

	receiverId := i.ApplicationCommandData().Options[0].StringValue()
	itemName := i.ApplicationCommandData().Options[1].StringValue()

	var targetWallet *dax.Wallet

	var receiverUserId string = ""
	var receiverWalletId string = ""
	if strings.Contains(receiverId, "@OTA") {
		// Possibly a wallet ID
		receiverWalletId = receiverId
	} else {
		// Possibly a user ID
		receiverUserId = utils.GetDiscordIdFromMentionFormat(receiverId)
	}

	var err error
	if receiverUserId != "" && receiverWalletId == "" {
		targetWallet, err = sharedRuntime.WalletService.GetWalletForUser(receiverUserId)
		if err != nil {
			interaction.SendErrorEmbedResponse(s, i.Interaction, err.Error())
			go logUtils.PublishDiscordLogErrorEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, err.Error())
			return
		}
	} else if receiverUserId == "" && receiverWalletId != "" {
		targetWallet, err = sharedRuntime.WalletService.GetWallet(receiverWalletId)
		if err != nil {
			interaction.SendErrorEmbedResponse(s, i.Interaction, err.Error())
			go logUtils.PublishDiscordLogErrorEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, err.Error())
			return
		}
	}

	targetUser, err := sharedRuntime.UserService.GetUser(targetWallet.UserId)
	if err != nil {
		interaction.SendErrorEmbedResponse(s, i.Interaction, err.Error())
		go logUtils.PublishDiscordLogErrorEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, err.Error())
		return
	}

	targetItem, err := sharedRuntime.MarketplaceService.GetItemFromMarketByName(itemName)
	if err != nil {
		interaction.SendErrorEmbedResponse(s, i.Interaction, err.Error())
		go logUtils.PublishDiscordLogErrorEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, err.Error())
		return
	}

	err = sharedRuntime.WalletService.ConsumeItemForUser(targetUser.DiscordTag, targetWallet.Id, targetItem.Id)
	if err != nil {
		interaction.SendErrorEmbedResponse(s, i.Interaction, err.Error())
		go logUtils.PublishDiscordLogErrorEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, err.Error())
		return
	}

	var useCount int = 1 // 1 for now, could be dynamic if using multiple items at a time becomes a featrue
	go logUtils.PublishDiscordLogInfoEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, fmt.Sprintf("`%s` used an item (`%d` x `%s`)", targetUser.DiscordTag, useCount, targetItem.DisplayName))

	interaction.SendSimpleEmbedSlashResponse(s, i.Interaction, fmt.Sprintf("Successfully used `%d` x `%s` in `%s`'s wallet [`%s`]. Make sure that a member of staff is aware of this action so they can deliver the benefits.", useCount, targetItem.DisplayName, targetUser.DiscordTag, targetWallet.Id))
}
