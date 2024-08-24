package economy

import (
	"fmt"

	"github.com/RazvanBerbece/AzteMarket/src/libs/models/dax"
	"github.com/RazvanBerbece/AzteMarket/src/libs/models/events"
	"github.com/RazvanBerbece/AzteMarket/src/libs/repositories"
	logUtils "github.com/RazvanBerbece/AzteMarket/src/libs/services/logger/utils"
)

type EconomyService struct {
	// repos
	CurrencySystemStateRepositoryRepository repositories.DbCurrencySystemStateRepositoryRepository
	// log channels
	ConsoleLogChannel chan events.LogEvent
}

func (s EconomyService) CreateCurrencySystem(guildId string, currencyName string, totalCurrencyAvailable float64, totalCurrencyInFlow float64, dateOfLastReplenish int64) (*dax.CurrencySystemState, error) {

	currencySystem, err := s.CurrencySystemStateRepositoryRepository.CreateCurrencySystem(guildId, currencyName, totalCurrencyAvailable, totalCurrencyInFlow, dateOfLastReplenish)
	if err != nil {
		go logUtils.PublishConsoleLogErrorEvent(s.ConsoleLogChannel, err.Error())
		return nil, fmt.Errorf("failed to create currency system for guild `%s`", guildId)
	}

	return currencySystem, nil
}

func (s EconomyService) GetCurrencyStateForGuild(guildId string) (*dax.CurrencySystemState, error) {

	currencySystem, err := s.CurrencySystemStateRepositoryRepository.GetCurrencyStateForGuild(guildId)
	if err != nil {
		go logUtils.PublishConsoleLogErrorEvent(s.ConsoleLogChannel, err.Error())
		return nil, fmt.Errorf("failed to retrieve currency system state for guild `%s`", guildId)
	}

	return currencySystem, nil
}

func (s EconomyService) ReplenishCurrencyForGuild(guildId string, currencyAmount float64) error {

	err := s.CurrencySystemStateRepositoryRepository.ReplenishCurrencyForGuild(guildId, currencyAmount)
	if err != nil {
		go logUtils.PublishConsoleLogErrorEvent(s.ConsoleLogChannel, err.Error())
		return fmt.Errorf("failed to replenish currency for guild `%s`", guildId)
	}

	return nil
}

func (s EconomyService) AllocateFlowingCurrencyForGuild(guildId string, currencyAmount float64) error {

	err := s.CurrencySystemStateRepositoryRepository.AllocateFlowingCurrencyForGuild(guildId, currencyAmount)
	if err != nil {
		go logUtils.PublishConsoleLogErrorEvent(s.ConsoleLogChannel, err.Error())
		return fmt.Errorf("failed to allocate flowing currency for guild `%s`", guildId)
	}

	return nil
}

func (s EconomyService) DeallocateFlowingCurrencyForGuild(guildId string, currencyAmount float64) error {

	err := s.CurrencySystemStateRepositoryRepository.DeallocateFlowingCurrencyForGuild(guildId, currencyAmount)
	if err != nil {
		go logUtils.PublishConsoleLogErrorEvent(s.ConsoleLogChannel, err.Error())
		return fmt.Errorf("failed to deallocate flowing currency for guild `%s`", guildId)
	}

	return nil
}

func (s EconomyService) DeleteCurrencySystem(guildId string) error {

	err := s.CurrencySystemStateRepositoryRepository.DeleteCurrencySystem(guildId)
	if err != nil {
		go logUtils.PublishConsoleLogErrorEvent(s.ConsoleLogChannel, err.Error())
		return fmt.Errorf("failed to delete currency system for guild `%s`", guildId)
	}

	return nil
}
