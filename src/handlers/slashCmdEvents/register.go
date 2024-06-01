package slashCmdEvents

import (
	"fmt"

	logUtils "github.com/RazvanBerbece/AzteMarket/src/libs/services/logger/utils"
	sharedConfig "github.com/RazvanBerbece/AzteMarket/src/shared/config"
	sharedRuntime "github.com/RazvanBerbece/AzteMarket/src/shared/runtime"
	"github.com/bwmarrin/discordgo"
)

func CreateAndRegisterSlashEventHandlers(s *discordgo.Session, mainGuildOnly bool, commands []*discordgo.ApplicationCommand) error {

	log := fmt.Sprintf("[STARTUP] Overwriting %d slash commands...", len(commands))
	go logUtils.PublishConsoleLogInfoEvent(sharedRuntime.LogEventsChannel, log)

	// Create the code-defined commands in the target guild(s)
	err := CreateSlashCommandsOnRemoteGuilds(s, sharedConfig.DiscordBotAppId, mainGuildOnly, &sharedConfig.DiscordMainGuildId, commands)
	if err != nil {
		go logUtils.PublishConsoleLogErrorEvent(sharedRuntime.LogEventsChannel, err.Error())
		return err
	}

	// Run actual slash command handlers
	go HandleSlashEvents(s)

	return nil
}

func CreateSlashCommandsOnRemoteGuilds(s *discordgo.Session, appId string, mainGuildOnly bool, mainGuildId *string, commands []*discordgo.ApplicationCommand) error {

	if mainGuildOnly {
		// Register commands only for the main guild
		// This is more performant when the bot is not supposed to be in more guilds
		_, err := s.ApplicationCommandBulkOverwrite(appId, *mainGuildId, commands)
		if err != nil {
			return fmt.Errorf("an error ocurred while bulk overwriting slash commands in main guild: %v", err)
		}
	} else {
		// For each guild where the bot exists in, register the available commands
		for _, guild := range s.State.Guilds {
			_, err := s.ApplicationCommandBulkOverwrite(appId, guild.ID, commands)
			if err != nil {
				return fmt.Errorf("an error ocurred while bulk overwriting slash commands in guild %s: %v", guild.Name, err)
			}
		}
	}

	return nil
}
