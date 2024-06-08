package userServices

import (
	"log"

	"github.com/RazvanBerbece/AzteMarket/pkg/utils"
	"github.com/RazvanBerbece/AzteMarket/src/libs/models/dax"
	"github.com/RazvanBerbece/AzteMarket/src/libs/models/events"
	"github.com/RazvanBerbece/AzteMarket/src/libs/repositories"
	logUtils "github.com/RazvanBerbece/AzteMarket/src/libs/services/logger/utils"
	"github.com/bwmarrin/discordgo"
)

type UserService struct {
	// repos
	UsersRepository repositories.DbUserRepository
	// log channels
	ConsoleLogChannel chan events.LogEvent
}

func (s UserService) GetUser(userId string) (*dax.User, error) {

	user, err := s.UsersRepository.GetUser(userId)
	if err != nil {
		go logUtils.PublishConsoleLogErrorEvent(s.ConsoleLogChannel, err.Error())
		return nil, err
	}

	return user, nil

}

func (s UserService) UserIsOfStaffType(session *discordgo.Session, guildId string, userId string, staffRoles []string) bool {

	user, err := session.GuildMember(guildId, userId)
	if err != nil {
		log.Printf("Cannot retrieve Discord user with id %s: %v", userId, err)
		return false
	}

	roles, err := session.GuildRoles(guildId)
	if err != nil {
		log.Printf("Cannot retrieve Discord roles guild with id %s: %v", guildId, err)
		return false
	}
	for _, discordRole := range roles {
		if utils.StringInSlice(discordRole.ID, user.Roles) && utils.StringInSlice(discordRole.Name, staffRoles) {
			return true
		}
	}

	return false
}
