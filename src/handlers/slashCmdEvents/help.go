package slashCmdEvents

import (
	"fmt"

	"github.com/RazvanBerbece/AzteMarket/pkg/dm"
	"github.com/RazvanBerbece/AzteMarket/pkg/embed"
	"github.com/RazvanBerbece/AzteMarket/pkg/interaction"
	"github.com/RazvanBerbece/AzteMarket/pkg/utils"
	sharedConfig "github.com/RazvanBerbece/AzteMarket/src/shared/config"
	sharedRuntime "github.com/RazvanBerbece/AzteMarket/src/shared/runtime"
	"github.com/bwmarrin/discordgo"
)

// TODO: Refactor the slashCmdEvents folder structure and move this outside of this parent scope.
func HandleSlashAzteMarketHelp(s *discordgo.Session, i *discordgo.InteractionCreate) {

	authorUserId := i.Member.User.ID
	authorGuildId := i.GuildID

	go sendHelpGuideToUser(s, i.Interaction, authorUserId, authorGuildId)
}

func sendHelpGuideToUser(s *discordgo.Session, i *discordgo.Interaction, userId string, guildId string) {

	embedToSend := embed.NewEmbed().
		SetAuthor("AzteMarket", "https://i.postimg.cc/262tK7VW/148c9120-e0f0-4ed5-8965-eaa7c59cc9f2-2.jpg").
		SetTitle("🤖   Command Guide").
		SetThumbnail("https://i.postimg.cc/262tK7VW/148c9120-e0f0-4ed5-8965-eaa7c59cc9f2-2.jpg").
		SetColor(sharedConfig.EmbedColorCode)

	// Build guide message from all available and registered commands
	for _, cmd := range DefinedSlashCommands {
		title := fmt.Sprintf("`/%s`", cmd.Name)
		if len(cmd.Options) > 0 {
			for _, param := range cmd.Options {
				var required string
				if param.Required {
					required = "required"
				} else {
					required = "optional"
				}
				title += fmt.Sprintf(" `[%s (%s)]`", param.Name, required)
			}
		}

		if utils.StringInSlice(cmd.Name, sharedConfig.HigherStaffCommands) {
			// don't show restricted higher staff commands
			if sharedRuntime.UserService.UserIsOfStaffType(s, guildId, userId, sharedConfig.HigherStaffRoles) {
				// unless a member of higher staff ran the /help handler
				enrichedTitle := fmt.Sprintf("%s *(higher staff command)*", title)
				embedToSend.AddField(enrichedTitle, cmd.Description, false)
			} else {
				continue
			}
		} else if utils.StringInSlice(cmd.Name, sharedConfig.StaffCommands) {
			// don't show restricted lower staff commands
			if sharedRuntime.UserService.UserIsOfStaffType(s, guildId, userId, sharedConfig.StaffRoles) {
				// unless a member of lower staff ran the /help handler
				enrichedTitle := fmt.Sprintf("%s *(staff command)*", title)
				embedToSend.AddField(enrichedTitle, cmd.Description, false)
			} else {
				continue
			}
		} else {
			embedToSend.AddField(title, cmd.Description, false)
		}
	}

	// Send paginated help DM to user
	paginationRow := embed.GetPaginationActionRowForEmbed(sharedRuntime.PreviousPageOnEmbedEventId, sharedRuntime.NextPageOnEmbedEventId)
	go dm.SendDirectComplexEmbedToMember(s, userId, *embedToSend, paginationRow, sharedConfig.EmbedPageSize, sharedRuntime.EmbedsToPaginate)

	// Response on channel to satisfy Discord protocol
	interaction.SendSimpleEmbedSlashResponse(s, i, fmt.Sprintf("You should have received a help guide for the <@%s> in your DMs !", sharedConfig.DiscordBotAppId))

}
