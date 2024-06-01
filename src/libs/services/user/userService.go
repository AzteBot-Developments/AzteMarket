package userServices

import (
	"fmt"

	"github.com/RazvanBerbece/AzteMarket/src/libs/models/dax"
	"github.com/RazvanBerbece/AzteMarket/src/libs/models/events"
	"github.com/RazvanBerbece/AzteMarket/src/libs/repositories"
	loggerService "github.com/RazvanBerbece/AzteMarket/src/libs/services/logger/strategies"
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
		s.ConsoleLogChannel <- events.LogEvent{
			Logger: loggerService.NewConsoleLogger(),
			Msg:    fmt.Sprintf("Could not read user %s from DB: %v", userId, err),
			Type:   "ERROR",
		}
		return nil, err
	}

	return user, nil

}
