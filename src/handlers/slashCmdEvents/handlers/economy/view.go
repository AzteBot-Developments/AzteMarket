package slashCmdEconomyHandlers

import (
	"fmt"

	"github.com/RazvanBerbece/AzteMarket/pkg/embed"
	"github.com/RazvanBerbece/AzteMarket/pkg/interaction"
	"github.com/RazvanBerbece/AzteMarket/pkg/utils"
	"github.com/RazvanBerbece/AzteMarket/src/libs/models/events"
	sharedConfig "github.com/RazvanBerbece/AzteMarket/src/shared/config"
	sharedRuntime "github.com/RazvanBerbece/AzteMarket/src/shared/runtime"
	"github.com/bwmarrin/discordgo"
)

func HandleSlashViewEconomy(s *discordgo.Session, i *discordgo.InteractionCreate) {

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: embed.SimpleEmbed("🤖   Slash Command Confirmation", "Processing `/economy` command...", sharedConfig.EmbedColorCode),
		},
	})

	guild, err := s.Guild(i.GuildID)
	if err != nil {
		interaction.ErrorEmbedResponseEdit(s, i.Interaction, err.Error())
		return
	}

	economy, err := sharedRuntime.EconomyService.GetCurrencyStateForGuild(i.GuildID)
	if err != nil {
		interaction.ErrorEmbedResponseEdit(s, i.Interaction, err.Error())
		return
	}

	embedToSend := embed.NewEmbed().
		SetAuthor(fmt.Sprintf("Economy Report - %s", economy.CurrencyName), "https://i.postimg.cc/262tK7VW/148c9120-e0f0-4ed5-8965-eaa7c59cc9f2-2.jpg").
		SetDescription(fmt.Sprintf("This is the current economic system state of the `%s` currency in the `%s` server.", economy.CurrencyName, guild.Name)).
		SetColor(sharedConfig.EmbedColorCode).
		DecorateWithTimestampFooter("Mon, 02 Jan 2006 15:04:05 MST").
		AddLineBreakField().
		AddField("Globally Allocated Amount of Funds", fmt.Sprintf("`%.2f` %s", economy.TotalCurrencyAvailable, economy.CurrencyName), false).
		AddField("Amount of Funds In-Flow", fmt.Sprintf("`%.2f` %s", economy.TotalCurrencyInFlow, economy.CurrencyName), false).
		AddField("Timestamp of Last Replenishment", fmt.Sprintf("`%s`", utils.FormatUnixAsString(economy.DateOfLastReplenish, "Mon, 02 Jan 2006 15:04:05 MST")), false)

	sharedRuntime.ComplexResponsesChannel <- events.ComplexResponseEvent{
		Interaction: i.Interaction,
		Embed:       embedToSend,
	}

}
