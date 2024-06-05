package slashCmdMarketHandlers

import (
	"fmt"

	"github.com/RazvanBerbece/AzteMarket/pkg/embed"
	"github.com/RazvanBerbece/AzteMarket/pkg/interaction"
	logUtils "github.com/RazvanBerbece/AzteMarket/src/libs/services/logger/utils"
	sharedConfig "github.com/RazvanBerbece/AzteMarket/src/shared/config"
	sharedRuntime "github.com/RazvanBerbece/AzteMarket/src/shared/runtime"
	"github.com/bwmarrin/discordgo"
)

func HandleSlashViewItemOnMarket(s *discordgo.Session, i *discordgo.InteractionCreate) {

	itemId := i.ApplicationCommandData().Options[0].StringValue()
	if len(itemId) < 1 {
		interaction.SendErrorEmbedResponse(s, i.Interaction, fmt.Sprintf("Argument `%s` invalid (term: `%s`)", i.ApplicationCommandData().Options[0].Name, itemId))
		return
	}

	item, err := sharedRuntime.MarketplaceService.GetItemFromMarket(itemId)
	if err != nil {
		interaction.SendErrorEmbedResponse(s, i.Interaction, err.Error())
		go logUtils.PublishDiscordLogErrorEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, err.Error())
		return
	}

	embedToSend := embed.NewEmbed().
		SetAuthor("AzteMarket", "https://i.postimg.cc/262tK7VW/148c9120-e0f0-4ed5-8965-eaa7c59cc9f2-2.jpg").
		SetTitle(fmt.Sprintf("💷    `%s` (id: `%s`)", item.DisplayName, item.Id)).
		SetDescription(item.Details).
		SetColor(sharedConfig.EmbedColorCode).
		DecorateWithTimestampFooter("Mon, 02 Jan 2006 15:04:05 MST").
		AddField("Available to Buy", fmt.Sprintf("`%d` units", item.NumAvailable), false).
		AddField("Cost", fmt.Sprintf("`%.2f` AzteCoins", item.Cost), false)

	interaction.SendEmbedSlashResponse(s, i.Interaction, *embedToSend)

}
