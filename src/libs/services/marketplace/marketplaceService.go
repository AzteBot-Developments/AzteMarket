package marketplaceServices

import (
	"fmt"
	"strings"

	"github.com/RazvanBerbece/AzteMarket/src/libs/models/dax"
	"github.com/RazvanBerbece/AzteMarket/src/libs/models/events"
	"github.com/RazvanBerbece/AzteMarket/src/libs/repositories"
	logUtils "github.com/RazvanBerbece/AzteMarket/src/libs/services/logger/utils"
)

type MarketplaceService struct {
	// repos
	StockRepository   repositories.DbStockRepository
	WalletsRepository repositories.DbWalletsRepository
	// log channels
	ConsoleLogChannel chan events.LogEvent
}

func (s MarketplaceService) AddItemForSaleOnMarket(itemName string, itemDetails string, cost float64, numAvailable int) (*string, error) {

	itemId, err := s.StockRepository.AddStockItem(itemName, itemDetails, cost, numAvailable)
	if err != nil {
		go logUtils.PublishConsoleLogErrorEvent(s.ConsoleLogChannel, err.Error())
		return nil, err
	}

	return itemId, nil

}

func (s MarketplaceService) GetItemFromMarket(itemId string) (*dax.StockItem, error) {

	item, err := s.StockRepository.GetStockItem(itemId)
	if err != nil {
		go logUtils.PublishConsoleLogErrorEvent(s.ConsoleLogChannel, err.Error())
		return nil, err
	}

	return item, nil
}

func (s MarketplaceService) GetItemFromMarketByName(itemName string) (*dax.StockItem, error) {

	item, err := s.StockRepository.GetStockItemByName(itemName)
	if err != nil {
		go logUtils.PublishConsoleLogErrorEvent(s.ConsoleLogChannel, err.Error())
		return nil, err
	}

	return item, nil
}

func (s MarketplaceService) GetAllItemsOnMarket() ([]dax.StockItem, error) {

	items, err := s.StockRepository.GetAllItems()
	if err != nil {
		go logUtils.PublishConsoleLogErrorEvent(s.ConsoleLogChannel, err.Error())
		return nil, err
	}

	return items, nil
}

func (s MarketplaceService) ClearMarket() (int64, error) {

	deletedCount, err := s.StockRepository.DeleteAllItems()
	if err != nil {
		go logUtils.PublishConsoleLogErrorEvent(s.ConsoleLogChannel, err.Error())
		return -1, err
	}

	return deletedCount, nil
}

func (s MarketplaceService) BuyItem(buyerUserId string, itemId string) error {

	item, err := s.StockRepository.GetStockItem(itemId)
	if err != nil {
		go logUtils.PublishConsoleLogErrorEvent(s.ConsoleLogChannel, err.Error())
		return err
	}

	// Ensure that the item can be bought (i.e stock is available)
	if item.NumAvailable <= 0 {
		return fmt.Errorf("the item with ID `%s` is no longer in stock", itemId)
	}

	buyerWallet, err := s.WalletsRepository.GetWalletForUser(buyerUserId)
	if err != nil {
		go logUtils.PublishConsoleLogErrorEvent(s.ConsoleLogChannel, err.Error())
		return err
	}

	// Ensure that user has enough funds to buy the item
	tax := 0.005
	if buyerWallet.Funds < item.Cost+tax {
		return fmt.Errorf("cannot buy item `%s` because the buyer's wallet doesn't have enough available funds (available: `%.2f`)", itemId, buyerWallet.Funds)
	}

	// Only allow a maximum of items of the same ID / name in one's wallet at all times
	threshold := 2
	inventoryString := buyerWallet.Inventory
	itemIds := strings.Split(inventoryString, ",")
	itemCount := 0
	for _, id := range itemIds {
		if id == itemId {
			itemCount += 1
		}
	}
	if itemCount > threshold {
		return fmt.Errorf("cannot buy item `%s` because the buyer's wallet has reached the limit of items of this type", itemId)
	}

	// Subtract funds
	// TODO: Could send to server wallet instead ? If I ever get around to doing a server wallet and ICOs
	err = s.WalletsRepository.SubtractFundsFromWallet(buyerWallet.Id, item.Cost)
	if err != nil {
		go logUtils.PublishConsoleLogErrorEvent(s.ConsoleLogChannel, err.Error())
		return err
	}

	// Add item ID to user's inventory
	err = s.WalletsRepository.AddItemToWallet(buyerWallet.Id, item.Id)
	if err != nil {
		go logUtils.PublishConsoleLogErrorEvent(s.ConsoleLogChannel, err.Error())
		return err
	}

	// Decrement num of available units of this item on the market
	err = s.StockRepository.DecrementAvailableForItem(item.Id)
	if err != nil {
		go logUtils.PublishConsoleLogErrorEvent(s.ConsoleLogChannel, err.Error())
		return err
	}

	return nil

}
