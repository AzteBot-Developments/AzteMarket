package repositories

import (
	"fmt"

	"github.com/RazvanBerbece/AzteMarket/pkg/utils"
	"github.com/RazvanBerbece/AzteMarket/src/libs/models/dax"
)

type DbWalletsRepository interface {
	CreateWalletForUser(userId string) (*dax.Wallet, error)
	GetWalletForUser(userId string) (*dax.Wallet, error)
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

func (r WalletsRepository) GetWalletForUser(userId string) (*dax.Wallet, error) {

	query := "SELECT * FROM Wallets WHERE userId = ?"
	row := r.DbContext.SqlDb.QueryRow(query, userId)

	var wallet dax.Wallet
	err := row.Scan(&wallet.UserId,
		&wallet.Id,
		&wallet.Funds,
		&wallet.Inventory)

	if err != nil {
		return nil, fmt.Errorf("an error ocurred while retrieving wallet for user with ID `%s`", userId)
	}

	return &wallet, nil

}
