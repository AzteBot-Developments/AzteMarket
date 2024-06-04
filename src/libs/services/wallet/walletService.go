package walletServices

import (
	"github.com/RazvanBerbece/AzteMarket/src/libs/models/dax"
	"github.com/RazvanBerbece/AzteMarket/src/libs/models/events"
	"github.com/RazvanBerbece/AzteMarket/src/libs/repositories"
	logUtils "github.com/RazvanBerbece/AzteMarket/src/libs/services/logger/utils"
)

type WalletService struct {
	// repos
	WalletsRepository repositories.DbWalletsRepository
	// log channels
	ConsoleLogChannel chan events.LogEvent
}

func (s WalletService) CreateWalletForUser(userId string) (*dax.Wallet, error) {

	wallet, err := s.WalletsRepository.CreateWalletForUser(userId)
	if err != nil {
		go logUtils.PublishConsoleLogErrorEvent(s.ConsoleLogChannel, err.Error())
		return nil, err
	}

	return wallet, nil

}

func (s WalletService) GetWalletForUser(userId string) (*dax.Wallet, error) {

	wallet, err := s.WalletsRepository.GetWalletForUser(userId)
	if err != nil {
		go logUtils.PublishConsoleLogErrorEvent(s.ConsoleLogChannel, err.Error())
		return nil, err
	}

	return wallet, nil

}
