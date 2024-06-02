package interaction

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func GetDiscordRole(s *discordgo.Session, guildId string, roleId string) (*discordgo.Role, error) {

	roles, err := s.GuildRoles(guildId)
	if err != nil {
		fmt.Printf("Error retrieving roles for guild with ID %s: %v\n", guildId, err)
		return nil, err
	}

	for _, role := range roles {
		if role.ID == roleId {
			return role, nil
		}
	}

	return nil, fmt.Errorf("a role with ID %s hasn't been found", roleId)

}
