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
		return nil, fmt.Errorf("failed to add given item to the market")
	}

	return itemId, nil
}

func (s MarketplaceService) GetItemFromMarket(itemId string) (*dax.StockItem, error) {

	item, err := s.StockRepository.GetStockItem(itemId)
	if err != nil {
		go logUtils.PublishConsoleLogErrorEvent(s.ConsoleLogChannel, err.Error())
		return nil, fmt.Errorf("failed to retrieve item with ID `%s` from the market", itemId)
	}

	return item, nil
}

func (s MarketplaceService) GetItemFromMarketByName(itemName string) (*dax.StockItem, error) {

	item, err := s.StockRepository.GetStockItemByName(itemName)
	if err != nil {
		go logUtils.PublishConsoleLogErrorEvent(s.ConsoleLogChannel, err.Error())
		return nil, fmt.Errorf("failed to retrieve item with name `%s` from the market", itemName)
	}

	return item, nil
}

func (s MarketplaceService) GetAllItemsOnMarket() ([]dax.StockItem, error) {

	items, err := s.StockRepository.GetAllItems()
	if err != nil {
		go logUtils.PublishConsoleLogErrorEvent(s.ConsoleLogChannel, err.Error())
		return nil, fmt.Errorf("failed to retrieve all items from the market")
	}

	return items, nil
}

func (s MarketplaceService) ClearMarket() (int64, error) {

	deletedCount, err := s.StockRepository.DeleteAllItems()
	if err != nil {
		go logUtils.PublishConsoleLogErrorEvent(s.ConsoleLogChannel, err.Error())
		return -1, fmt.Errorf("failed to clear all items from the market")
	}

	return deletedCount, nil
}

func (s MarketplaceService) RemoveItemFromMarket(itemId string) error {

	err := s.StockRepository.DeleteItem(itemId)
	if err != nil {
		go logUtils.PublishConsoleLogErrorEvent(s.ConsoleLogChannel, err.Error())
		return fmt.Errorf("failed to delete item with ID `%s` from the market", itemId)
	}

	return nil
}

func (s MarketplaceService) BuyItem(buyerUserId string, itemId string) error {

	item, err := s.StockRepository.GetStockItem(itemId)
	if err != nil {
		go logUtils.PublishConsoleLogErrorEvent(s.ConsoleLogChannel, err.Error())
		return fmt.Errorf("failed to retrieve item with ID `%s`", itemId)
	}

	// Ensure that the item can be bought (i.e stock is available)
	if item.NumAvailable <= 0 {
		return fmt.Errorf("the item with ID `%s` is no longer in stock", itemId)
	}

	buyerWallet, err := s.WalletsRepository.GetWalletForUser(buyerUserId)
	if err != nil {
		go logUtils.PublishConsoleLogErrorEvent(s.ConsoleLogChannel, err.Error())
		return fmt.Errorf("failed to retrieve wallet for user `%s`", buyerUserId)
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
		return fmt.Errorf("failed to subtract `%.2f` funds from user with ID `%s`", item.Cost, buyerWallet.UserId)
	}

	// Add item ID to user's inventory
	err = s.WalletsRepository.AddItemToWallet(buyerWallet.Id, item.Id)
	if err != nil {
		go logUtils.PublishConsoleLogErrorEvent(s.ConsoleLogChannel, err.Error())
		return fmt.Errorf("failed to add item `%s` to user with ID `%s`", item.DisplayName, buyerWallet.UserId)
	}

	// Decrement num of available units of this item on the market
	err = s.StockRepository.DecrementAvailableForItem(item.Id, 1)
	if err != nil {
		go logUtils.PublishConsoleLogErrorEvent(s.ConsoleLogChannel, err.Error())
		return fmt.Errorf("failed to remove `%d` units from item `%s`", 1, item.DisplayName)
	}

	return nil

}

func (s MarketplaceService) RemoveStockUnitsForItem(itemId string, multiplier int) error {

	item, err := s.GetItemFromMarket(itemId)
	if err != nil {
		go logUtils.PublishConsoleLogErrorEvent(s.ConsoleLogChannel, err.Error())
		return fmt.Errorf("failed to retrieve item with ID `%s`", itemId)
	}

	// Domain level validation
	if item.NumAvailable < multiplier {
		return fmt.Errorf("cannot remove more units from an item than there are available (`%d` < `%d`)", item.NumAvailable, multiplier)
	}

	err = s.StockRepository.DecrementAvailableForItem(itemId, multiplier)
	if err != nil {
		go logUtils.PublishConsoleLogErrorEvent(s.ConsoleLogChannel, err.Error())
		return fmt.Errorf("failed to remove `%d` units from item `%s`", multiplier, item.DisplayName)
	}

	return nil

}

func (s MarketplaceService) AddStockUnitsForItem(itemId string, multiplier int) error {

	item, err := s.GetItemFromMarket(itemId)
	if err != nil {
		go logUtils.PublishConsoleLogErrorEvent(s.ConsoleLogChannel, err.Error())
		return fmt.Errorf("failed to retrieve item with ID `%s`", itemId)
	}

	err = s.StockRepository.IncrementAvailableForItem(itemId, multiplier)
	if err != nil {
		go logUtils.PublishConsoleLogErrorEvent(s.ConsoleLogChannel, err.Error())
		return fmt.Errorf("failed to add `%d` new units to item `%s`", multiplier, item.DisplayName)
	}

	return nil

}
