package slashCmdEconomyHandlers

import (
	"time"

	"github.com/RazvanBerbece/AzteMarket/pkg/interaction"
	"github.com/RazvanBerbece/AzteMarket/pkg/utils"
	logUtils "github.com/RazvanBerbece/AzteMarket/src/libs/services/logger/utils"
	sharedConfig "github.com/RazvanBerbece/AzteMarket/src/shared/config"
	sharedRuntime "github.com/RazvanBerbece/AzteMarket/src/shared/runtime"
	"github.com/bwmarrin/discordgo"
)

func HandleSlashCreateEconomy(s *discordgo.Session, i *discordgo.InteractionCreate) {

	// Retrieve input args from command
	currencyName := i.ApplicationCommandData().Options[0].StringValue()
	totalCurrencyAvailable := i.ApplicationCommandData().Options[1].StringValue()

	fTotalCurrencyAvailable, err := utils.StringToFloat64(totalCurrencyAvailable)
	if err != nil {
		interaction.SendErrorEmbedResponse(s, i.Interaction, err.Error())
		return
	}

	_, err = sharedRuntime.EconomyService.CreateCurrencySystem(i.GuildID, currencyName, *fTotalCurrencyAvailable, 0, time.Now().Unix())
	if err != nil {
		interaction.SendErrorEmbedResponse(s, i.Interaction, err.Error())
		go logUtils.PublishDiscordLogErrorEvent(
			sharedRuntime.LogEventsChannel, s, "Debug",
			sharedConfig.DiscordChannelTopicPairs,
			err.Error(),
		)
		return
	}

	go logUtils.PublishDiscordLogInfoEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, "A new economy was created for this server")

	// Final response to interaction
	interaction.SendSimpleEmbedSlashResponse(s, i.Interaction, "Your server economy has been created !")

}
