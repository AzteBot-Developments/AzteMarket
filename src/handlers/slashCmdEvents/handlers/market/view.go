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
			Embeds: embed.SimpleEmbed("🤖   Slash Command Confirmation", "Processing `/market` command..."),
		},
	})

	items, err := sharedRuntime.MarketplaceService.GetAllItemsOnMarket()
	if err != nil {
		interaction.ErrorEmbedResponseEdit(s, i.Interaction, err.Error())
		return
	}

	embedToSend := embed.NewEmbed().
		SetAuthor("AzteMarket", "https://i.postimg.cc/262tK7VW/148c9120-e0f0-4ed5-8965-eaa7c59cc9f2-2.jpg").
		SetDescription("The AzteMarket is an exchange which offers up various benefits for members to buy via AzteCoins.").
		SetThumbnail("https://i.postimg.cc/262tK7VW/148c9120-e0f0-4ed5-8965-eaa7c59cc9f2-2.jpg").
		SetColor(sharedConfig.EmbedColorCode).
		AddLineBreakField().
		DecorateWithTimestampFooter("Mon, 02 Jan 2006 15:04:05 MST").
		AddField(fmt.Sprintf("There are `%d` items available to buy on the AzteMarket at the moment.", len(items)), "", false)

	for idx, item := range items {
		embedToSend.AddField("", fmt.Sprintf("%d. `%s` [`%s`] (`%.2f` AzteCoins)\n%s", idx+1, item.DisplayName, item.Id, item.Cost, item.Details), false)
	}

	paginationRow := embed.GetPaginationActionRowForEmbed(sharedRuntime.PreviousPageOnEmbedEventId, sharedRuntime.NextPageOnEmbedEventId)
	sharedRuntime.ComplexResponsesChannel <- events.ComplexResponseEvent{
		Interaction:   i.Interaction,
		Embed:         embedToSend,
		PaginationRow: &paginationRow,
	}

}
