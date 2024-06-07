package slashCmdMarketHandlers

import (
	"fmt"

	"github.com/RazvanBerbece/AzteMarket/pkg/embed"
	"github.com/RazvanBerbece/AzteMarket/pkg/interaction"
	"github.com/RazvanBerbece/AzteMarket/src/libs/models/events"
	sharedConfig "github.com/RazvanBerbece/AzteMarket/src/shared/config"
	sharedRuntime "github.com/RazvanBerbece/AzteMarket/src/shared/runtime"
	"github.com/bwmarrin/discordgo"
)

func HandleSlashViewMarket(s *discordgo.Session, i *discordgo.InteractionCreate) {

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: embed.SimpleEmbed("🤖   Slash Command Confirmation", "Processing `/market` command...", sharedConfig.EmbedColorCode),
		},
	})

	items, err := sharedRuntime.MarketplaceService.GetAllItemsOnMarket()
	if err != nil {
		interaction.ErrorEmbedResponseEdit(s, i.Interaction, err.Error())
		return
	}

	// Only count a benefit as available when it's in stock
	var availableBenefitsCount int = 0
	for _, item := range items {
		if item.NumAvailable > 0 {
			availableBenefitsCount += 1
		}
	}

	// Retrieve the current command id for the buy interaction
	// so AzteMarket can make it into a clickable command in the response embed
	cmdId, err := interaction.GetCommandId(s, sharedConfig.DiscordBotAppId, sharedConfig.DiscordMainGuildId, "market-buy-item")
	if err != nil {
		interaction.ErrorEmbedResponseEdit(s, i.Interaction, err.Error())
		return
	}

	embedToSend := embed.NewEmbed().
		SetAuthor("AzteMarket", "https://i.postimg.cc/262tK7VW/148c9120-e0f0-4ed5-8965-eaa7c59cc9f2-2.jpg").
		SetDescription(fmt.Sprintf("<@%s> is an exchange which offers up various benefits for members to buy via AzteCoins.\nThe items are bought using slash commands and their associated IDs.", sharedConfig.DiscordBotAppId)).
		// SetThumbnail("https://i.postimg.cc/262tK7VW/148c9120-e0f0-4ed5-8965-eaa7c59cc9f2-2.jpg").
		SetColor(sharedConfig.EmbedColorCode).
		DecorateWithTimestampFooter("Mon, 02 Jan 2006 15:04:05 MST").
		AddField(fmt.Sprintf("Currently, there are `%d` benefits available to purchase on the AzteMarket.", availableBenefitsCount), "", false)

	for idx, item := range items {
		embedToSend.AddField("", fmt.Sprintf("%d. `%s` - `🪙 %.2f` AzteCoins (Available: `%d`)\nAdditional details: `%s`\nTo buy: </market-buy-item:%s> `%s`", idx+1, item.DisplayName, item.Cost, item.NumAvailable, item.Details, cmdId, item.Id), false)
	}

	paginationRow := embed.GetPaginationActionRowForEmbed(sharedRuntime.PreviousPageOnEmbedEventId, sharedRuntime.NextPageOnEmbedEventId)
	sharedRuntime.ComplexResponsesChannel <- events.ComplexResponseEvent{
		Interaction:   i.Interaction,
		Embed:         embedToSend,
		PaginationRow: &paginationRow,
	}

}
