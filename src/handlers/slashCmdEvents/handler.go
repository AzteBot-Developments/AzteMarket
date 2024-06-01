package slashCmdEvents

import (
	"fmt"
	"log"

	"github.com/RazvanBerbece/AzteMarket/pkg/utils"
	logUtils "github.com/RazvanBerbece/AzteMarket/src/libs/services/logger/utils"
	sharedConfig "github.com/RazvanBerbece/AzteMarket/src/shared/config"
	sharedRuntime "github.com/RazvanBerbece/AzteMarket/src/shared/runtime"
	"github.com/bwmarrin/discordgo"
)

func HandleSlashEvents(s *discordgo.Session) {

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {

		appData := i.ApplicationCommandData()

		// Permission validation for slash commands
		if i.Type == discordgo.InteractionApplicationCommand {
			if !UserHasEnoughPermissionsForCommand(s,
				i.Interaction,
				appData.Name,
				sharedConfig.HigherStaffRoles,
				sharedConfig.HigherStaffCommands,
				sharedConfig.StaffRoles,
				sharedConfig.StaffCommands,
			) {

				// Inform of missing permissions
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "You do not have the required role to use this command.",
					},
				})

				// Audit execution failure due to missing permissions
				log := fmt.Sprintf("User `%s` failed to run command `%s` due to lack of permissions", i.Member.User.Username, appData.Name)
				go logUtils.PublishDiscordLogWarnEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, log)

				return
			}
		}

		// Actual execution of the handler
		if handlerFunc, ok := RegisteredSlashCommandHandlers[i.ApplicationCommandData().Name]; ok {
			handlerFunc(s, i)
		}
	})
}

func UserHasEnoughPermissionsForCommand(s *discordgo.Session,
	i *discordgo.Interaction,
	commandName string,
	higherStaffRoles []string,
	higherStaffCommands []string,
	staffRoles []string,
	staffCommands []string) bool {

	var hasRequiredPerms bool = false

	isHigherStaffCommand := utils.StringInSlice(commandName, higherStaffCommands)
	isStaffCommand := utils.StringInSlice(commandName, staffCommands)

	// if not a restricted-access command, then allow any
	if !isHigherStaffCommand && !isStaffCommand {
		return true
	}

	if isHigherStaffCommand {
		for _, role := range i.Member.Roles {
			authorDiscordRole, err := s.State.Role(i.GuildID, role)
			if err != nil {
				log.Println("Error getting role:", err)
				return false
			}
			if utils.StringInSlice(authorDiscordRole.Name, higherStaffRoles) {
				hasRequiredPerms = true
				return true
			}
		}
	} else if isStaffCommand {
		for _, role := range i.Member.Roles {
			authorDiscordRole, err := s.State.Role(i.GuildID, role)
			if err != nil {
				log.Println("Error getting role:", err)
				return false
			}
			if utils.StringInSlice(authorDiscordRole.Name, staffRoles) {
				return true
			}
		}
	}

	return hasRequiredPerms

}
