package repositories

import (
	"fmt"

	"github.com/RazvanBerbece/AzteMarket/src/libs/models/dax"
	"github.com/google/uuid"
)

type DbStockRepository interface {
	AddStockItem(stockItemDisplayName string, stockItemDetails string, cost float64) error
	GetStockItem(stockItemId string) (*dax.StockItem, error)
}

type StockRepository struct {
	DbContext AztemarketDbContext
}

func NewStockRepository(connString string) StockRepository {
	repo := StockRepository{AztemarketDbContext{
		ConnectionString: connString,
	}}
	repo.DbContext.Connect()
	return repo
}

func (r StockRepository) GetStockItem(stockItemId string) (*dax.StockItem, error) {

	query := "SELECT * FROM Stock WHERE id = ?"
	row := r.DbContext.SqlDb.QueryRow(query, stockItemId)

	var item dax.StockItem
	err := row.Scan(&item.Id,
		&item.DisplayName,
		&item.Details,
		&item.Cost)

	if err != nil {
		return nil, fmt.Errorf("an error ocurred while retrieving stock item with ID `%s`", stockItemId)
	}

	return &item, nil

}

func (r StockRepository) AddStockItem(stockItemDisplayName string, stockItemDetails string, cost float64) error {

	stockItem := &dax.StockItem{
		Id:          uuid.New().String(),
		DisplayName: stockItemDisplayName,
		Details:     stockItemDetails,
		Cost:        cost,
	}

	stmt, err := r.DbContext.SqlDb.Prepare(`
		INSERT INTO 
			Stock(
				id, 
				displayName, 
				details,
				cost
			)
		VALUES(?, ?, ?, ?);`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(stockItem.Id, stockItem.DisplayName, stockItem.Details, stockItem.Cost)
	if err != nil {
		return err
	}

	return nil

}
