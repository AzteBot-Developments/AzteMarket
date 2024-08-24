package slashCmdEconomyHandlers

import (
	"github.com/RazvanBerbece/AzteMarket/pkg/interaction"
	"github.com/RazvanBerbece/AzteMarket/pkg/utils"
	logUtils "github.com/RazvanBerbece/AzteMarket/src/libs/services/logger/utils"
	sharedConfig "github.com/RazvanBerbece/AzteMarket/src/shared/config"
	sharedRuntime "github.com/RazvanBerbece/AzteMarket/src/shared/runtime"
	"github.com/bwmarrin/discordgo"
)

func HandleSlashReplenishEconomy(s *discordgo.Session, i *discordgo.InteractionCreate) {

	// Retrieve input args from command
	currencyAmount := i.ApplicationCommandData().Options[0].StringValue()

	fCurrencyAmount, err := utils.StringToFloat64(currencyAmount)
	if err != nil {
		interaction.SendErrorEmbedResponse(s, i.Interaction, err.Error())
		return
	}

	err = sharedRuntime.EconomyService.ReplenishCurrencyForGuild(i.GuildID, *fCurrencyAmount)
	if err != nil {
		interaction.SendErrorEmbedResponse(s, i.Interaction, err.Error())
		go logUtils.PublishDiscordLogErrorEvent(
			sharedRuntime.LogEventsChannel, s, "Debug",
			sharedConfig.DiscordChannelTopicPairs,
			err.Error(),
		)
	}

	go logUtils.PublishDiscordLogInfoEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, "A new economy was created for this server")

	// Final response to interaction
	interaction.SendSimpleEmbedSlashResponse(s, i.Interaction, "Your server economy has been created !")

}
