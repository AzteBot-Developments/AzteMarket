package walletServices

import (
	"fmt"

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

func (s WalletService) DeleteWalletForUser(userId string) error {

	err := s.WalletsRepository.DeleteWalletForUser(userId)
	if err != nil {
		go logUtils.PublishConsoleLogErrorEvent(s.ConsoleLogChannel, err.Error())
		return err
	}

	return nil

}

func (s WalletService) SendFunds(senderUserId string, receiverWalletId string, funds float64) (float64, error) {

	// ensure that sender has a wallet
	senderWallet, err := s.WalletsRepository.GetWalletForUser(senderUserId)
	if err != nil {
		return -1, fmt.Errorf("sender `%s` does not currently own a wallet for the AzteMarket; please create one in order to send funds", senderUserId)
	}

	// Validation
	if funds > senderWallet.Funds {
		return -1, fmt.Errorf("sender `%s` blocked from transfering more AzteCoins than they own (available: `🪙 %.2f`)", senderUserId, senderWallet.Funds)
	}
	if senderWallet.Id == receiverWalletId {
		return -1, fmt.Errorf("sender `%s` blocked from sending funds to own wallet", senderUserId)
	}

	// ensure that receiver has a wallet
	_, err = s.WalletsRepository.GetWallet(receiverWalletId)
	if err != nil {
		return -1, fmt.Errorf("receiver `%s` does not currently own a wallet for the AzteMarket; please ensure that the receiver owns a wallet to send funds to them", receiverWalletId)
	}

	// Remove funds from sender and update in DB
	err = s.WalletsRepository.SubtractFundsFromWallet(senderWallet.Id, funds)
	if err != nil {
		return -1, err
	}

	// Add funds to receiver and update in DB
	err = s.WalletsRepository.AddFundsToWallet(receiverWalletId, funds)
	if err != nil {
		return -1, err
	}

	return senderWallet.Funds - funds, nil
}
