package walletServices

import (
	"fmt"
	"strings"

	"github.com/RazvanBerbece/AzteMarket/pkg/dm"
	"github.com/RazvanBerbece/AzteMarket/pkg/embed"
	"github.com/RazvanBerbece/AzteMarket/pkg/utils"
	"github.com/RazvanBerbece/AzteMarket/src/libs/models/dax"
	"github.com/RazvanBerbece/AzteMarket/src/libs/models/events"
	"github.com/RazvanBerbece/AzteMarket/src/libs/repositories"
	logUtils "github.com/RazvanBerbece/AzteMarket/src/libs/services/logger/utils"
	sharedConfig "github.com/RazvanBerbece/AzteMarket/src/shared/config"
	"github.com/bwmarrin/discordgo"
)

type WalletService struct {
	// repos
	WalletsRepository repositories.DbWalletsRepository
	StockRepository   repositories.DbStockRepository
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

func (s WalletService) GetWallet(id string) (*dax.Wallet, error) {

	wallet, err := s.WalletsRepository.GetWallet(id)
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

func (s WalletService) SendFunds(session *discordgo.Session, senderUserId string, receiverWalletId string, funds float64) (float64, error) {

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
	// update receiver wallet with new funds
	receiverWallet, err := s.WalletsRepository.GetWallet(receiverWalletId)
	if err != nil {
		return -1, fmt.Errorf("receiver `%s` does not currently own a wallet for the AzteMarket; please ensure that the receiver owns a wallet to send funds to them", receiverWalletId)
	}

	// Announce target user that they received funds
	announcement := embed.NewEmbed().
		SetAuthor("AzteMarket Wallet Service").
		SetColor(sharedConfig.EmbedColorCode).
		DecorateWithTimestampFooter("Mon, 02 Jan 2006 15:04:05 MST").
		SetDescription(fmt.Sprintf("You have just received `🪙 %.2f` AzteCoins from `%s` !", funds, senderWallet.Id)).
		AddField("🧾 Your updated balance", fmt.Sprintf("`🪙 %.2f` AzteCoins", receiverWallet.Funds), false)

	// Ignore errs
	go dm.DmEmbedUser(session, receiverWallet.UserId, *announcement.MessageEmbed)

	return senderWallet.Funds - funds, nil
}

func (s WalletService) ConsumeItemForUser(userName string, walletId string, itemId string) error {

	wallet, err := s.WalletsRepository.GetWallet(walletId)
	if err != nil {
		return err
	}

	item, err := s.StockRepository.GetStockItem(itemId)
	if err != nil {
		return err
	}

	// Ensure that member actually owns the item
	inventoryString := wallet.Inventory
	ownedItemIds := strings.Split(inventoryString, ",")
	if !utils.StringInSlice(itemId, ownedItemIds) {
		return fmt.Errorf("wallet for user `%s` [`%s`] doesn't own the target item to consume (`%s` [`%s`])", userName, wallet.Id, item.DisplayName, itemId)
	}

	err = s.WalletsRepository.RemoveItemFromWallet(walletId, itemId)
	if err != nil {
		return err
	}

	return nil

}
