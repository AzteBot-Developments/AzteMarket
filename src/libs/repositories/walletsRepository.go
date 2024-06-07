package repositories

import (
	"fmt"

	"github.com/RazvanBerbece/AzteMarket/pkg/utils"
	"github.com/RazvanBerbece/AzteMarket/src/libs/models/dax"
)

type DbWalletsRepository interface {
	GetWallet(id string) (*dax.Wallet, error)
	CreateWalletForUser(userId string) (*dax.Wallet, error)
	GetWalletForUser(userId string) (*dax.Wallet, error)
	DeleteWalletForUser(userId string) error
	GetWalletIdForUser(userId string) (*string, error)
	AddFundsToWallet(id string, funds float64) error
	SubtractFundsFromWallet(id string, funds float64) error
	AddItemToWallet(id string, itemId string) error
	RemoveItemFromWallet(id string, itemId string) error
	// GetWalletById(id string)
}

type WalletsRepository struct {
	DbContext AztemarketDbContext
}

func NewWalletsRepository(connString string) WalletsRepository {
	repo := WalletsRepository{AztemarketDbContext{
		ConnectionString: connString,
	}}
	repo.DbContext.Connect()
	return repo
}

func (r WalletsRepository) CreateWalletForUser(userId string) (*dax.Wallet, error) {

	wallet := &dax.Wallet{
		Id:        utils.NewSuffixedGuid("@OTA"),
		UserId:    userId,
		Funds:     0,
		Inventory: "",
	}

	stmt, err := r.DbContext.SqlDb.Prepare(`
		INSERT INTO 
			Wallets(
				id, 
				userId, 
				funds,
				inventory
			)
		VALUES(?, ?, ?, ?);`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(wallet.Id, wallet.UserId, wallet.Funds, wallet.Inventory)
	if err != nil {
		return nil, err
	}

	return wallet, nil

}

func (r WalletsRepository) GetWalletIdForUser(userId string) (*string, error) {

	query := "SELECT id FROM Wallets WHERE userId = ?"
	row := r.DbContext.SqlDb.QueryRow(query, userId)

	var id string
	err := row.Scan(id)

	if err != nil {
		return nil, err
	}

	return &id, nil

}

func (r WalletsRepository) GetWallet(id string) (*dax.Wallet, error) {

	query := "SELECT * FROM Wallets WHERE id = ?"
	row := r.DbContext.SqlDb.QueryRow(query, id)

	var wallet dax.Wallet
	err := row.Scan(&wallet.UserId,
		&wallet.Id,
		&wallet.Funds,
		&wallet.Inventory)

	if err != nil {
		return nil, err
	}

	return &wallet, nil

}

func (r WalletsRepository) GetWalletForUser(userId string) (*dax.Wallet, error) {

	query := "SELECT * FROM Wallets WHERE userId = ?"
	row := r.DbContext.SqlDb.QueryRow(query, userId)

	var wallet dax.Wallet
	err := row.Scan(&wallet.UserId,
		&wallet.Id,
		&wallet.Funds,
		&wallet.Inventory)

	if err != nil {
		return nil, err
	}

	return &wallet, nil

}

func (r WalletsRepository) DeleteWalletForUser(userId string) error {

	query := "DELETE FROM Wallets WHERE userId = ?"

	_, err := r.DbContext.SqlDb.Exec(query, userId)
	if err != nil {
		return fmt.Errorf("error deleting wallet entry for user: %w", err)
	}

	return nil
}

func (r WalletsRepository) AddFundsToWallet(id string, funds float64) error {

	stmt, err := r.DbContext.SqlDb.Prepare(`
	UPDATE Wallets SET 
		funds = funds + ?
	WHERE id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(funds, id)
	if err != nil {
		return err
	}

	return nil
}

func (r WalletsRepository) SubtractFundsFromWallet(id string, funds float64) error {

	stmt, err := r.DbContext.SqlDb.Prepare(`
	UPDATE Wallets SET 
		funds = funds - ?
	WHERE id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(funds, id)
	if err != nil {
		return err
	}

	return nil
}

func (r WalletsRepository) GetWalletInventory(id string) (string, error) {

	query := "SELECT inventory FROM Wallets WHERE id = ?"
	row := r.DbContext.SqlDb.QueryRow(query, id)

	var inventoryString string
	err := row.Scan(&inventoryString)

	if err != nil {
		return "", err
	}

	return inventoryString, nil

}

func (r WalletsRepository) AddItemToWallet(id string, itemId string) error {

	// Get current inventory state
	inventory, err := r.GetWalletInventory(id)
	if err != nil {
		return err
	}

	// Append new item ID in-memory
	inventory += fmt.Sprintf("%s,", itemId)

	// Set upstream
	stmt, err := r.DbContext.SqlDb.Prepare(`
			UPDATE Wallets SET 
				inventory = ?
			WHERE id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(inventory, id)
	if err != nil {
		return err
	}

	return nil
}

func (r WalletsRepository) RemoveItemFromWallet(id string, itemId string) error {

	return fmt.Errorf("not supported yet")

}
