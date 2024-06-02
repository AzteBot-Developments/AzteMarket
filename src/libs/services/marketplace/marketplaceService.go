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

func (s MarketplaceService) AddItemForSaleOnMarket(itemName string, itemDetails string, cost float64) error {

	err := s.StockRepository.AddStockItem(itemName, itemDetails, cost)
	if err != nil {
		go logUtils.PublishConsoleLogErrorEvent(s.ConsoleLogChannel, err.Error())
		return err
	}

	return nil

}

func (s MarketplaceService) GetItemFromMarket(itemId string) (*dax.StockItem, error) {

	item, err := s.StockRepository.GetStockItem(itemId)
	if err != nil {
		go logUtils.PublishConsoleLogErrorEvent(s.ConsoleLogChannel, err.Error())
		return nil, err
	}

	return item, err
}
