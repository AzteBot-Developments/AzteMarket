package repositories

import (
	"fmt"

	"github.com/RazvanBerbece/AzteMarket/src/libs/models/dax"
	"github.com/google/uuid"
)

type DbStockRepository interface {
	AddStockItem(stockItemDisplayName string, stockItemDetails string, cost float64, numAvailable int) (*string, error)
	GetStockItem(stockItemId string) (*dax.StockItem, error)
	GetAllItems() ([]dax.StockItem, error)
	DeleteAllItems() (int64, error)
	DecrementAvailableForItem(stockItemId string) error
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
		&item.Cost,
		&item.NumAvailable)

	if err != nil {
		return nil, fmt.Errorf("an error ocurred while retrieving stock item with ID `%s`", stockItemId)
	}

	return &item, nil

}

func (r StockRepository) AddStockItem(stockItemDisplayName string, stockItemDetails string, cost float64, numAvailable int) (*string, error) {

	stockItem := &dax.StockItem{
		Id:           uuid.New().String(),
		DisplayName:  stockItemDisplayName,
		Details:      stockItemDetails,
		Cost:         cost,
		NumAvailable: numAvailable,
	}

	stmt, err := r.DbContext.SqlDb.Prepare(`
		INSERT INTO 
			Stock(
				id, 
				displayName, 
				details,
				cost,
				numAvailable
			)
		VALUES(?, ?, ?, ?, ?);`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(stockItem.Id, stockItem.DisplayName, stockItem.Details, stockItem.Cost, stockItem.NumAvailable)
	if err != nil {
		return nil, err
	}

	return &stockItem.Id, nil

}

func (r StockRepository) GetAllItems() ([]dax.StockItem, error) {

	var items []dax.StockItem

	rows, err := r.DbContext.SqlDb.Query("SELECT * FROM Stock")
	if err != nil {
		return nil, fmt.Errorf("an error ocurred while retrieving all items: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var item dax.StockItem
		if err := rows.Scan(&item.Id, &item.DisplayName, &item.Details, &item.Cost, &item.NumAvailable); err != nil {
			return nil, fmt.Errorf("error in Stock GetAll: %v", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error in Stock GetAll: %v", err)
	}

	return items, nil
}

func (r StockRepository) DeleteAllItems() (int64, error) {

	query := "DELETE FROM Stock;"

	res, err := r.DbContext.SqlDb.Exec(query)
	if err != nil {
		return -1, fmt.Errorf("error deleting all items from the Stock table: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return -1, fmt.Errorf("error reading rows affected for DeleteAllItems: %w", err)
	}

	return rowsAffected, nil
}

func (r StockRepository) DecrementAvailableForItem(stockItemId string) error {

	stmt, err := r.DbContext.SqlDb.Prepare(`
	UPDATE Stock SET 
		numAvailable = numAvailable - 1
	WHERE id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(stockItemId)
	if err != nil {
		return err
	}

	return nil
}
