package userServices

import (
	"github.com/RazvanBerbece/AzteMarket/src/libs/models/dax"
	"github.com/RazvanBerbece/AzteMarket/src/libs/models/events"
	"github.com/RazvanBerbece/AzteMarket/src/libs/repositories"
	logUtils "github.com/RazvanBerbece/AzteMarket/src/libs/services/logger/utils"
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
