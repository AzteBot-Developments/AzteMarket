package interaction

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func GetCommandId(s *discordgo.Session, appId string, guildId string, cmdName string) (string, error) {

	commands, err := s.ApplicationCommands(appId, guildId)
	if err != nil {
		return "", fmt.Errorf("error fetching commands: %v", err)
	}

	for _, cmd := range commands {
		if cmd.Name == cmdName {
			return cmd.ID, nil
		}
	}

	return "", fmt.Errorf("command '%s' not found", cmdName)
}
