package marketplaceServices

import (
	"github.com/RazvanBerbece/AzteMarket/src/libs/models/dax"
	"github.com/RazvanBerbece/AzteMarket/src/libs/models/events"
	"github.com/RazvanBerbece/AzteMarket/src/libs/repositories"
	logUtils "github.com/RazvanBerbece/AzteMarket/src/libs/services/logger/utils"
)

type MarketplaceService struct {
	// repos
	StockRepository repositories.DbStockRepository
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
