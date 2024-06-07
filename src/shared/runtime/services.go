package sharedRuntime

import (
	"os"

	"github.com/RazvanBerbece/AzteMarket/src/libs/repositories"
	marketplaceServices "github.com/RazvanBerbece/AzteMarket/src/libs/services/marketplace"
	userServices "github.com/RazvanBerbece/AzteMarket/src/libs/services/user"
	walletServices "github.com/RazvanBerbece/AzteMarket/src/libs/services/wallet"
)

// Connection strings
var MySqlAztebotRootConnectionString = os.Getenv("DB_AZTEBOT_ROOT_CONNSTRING") // in MySQL format (i.e. `root_username:root_password@<unix/tcp>(host:port)/db_name`)
var MySqlAztemarketRootConnectionString = os.Getenv("DB_AZTEMARKET_ROOT_CONNSTRING")

// Create shared services at runtime to use across the app
var UserService = userServices.UserService{
	UsersRepository:   repositories.NewUserRepository(MySqlAztebotRootConnectionString),
	ConsoleLogChannel: LogEventsChannel,
}

var MarketplaceService = marketplaceServices.MarketplaceService{
	StockRepository:   repositories.NewStockRepository(MySqlAztemarketRootConnectionString),
	WalletsRepository: repositories.NewWalletsRepository(MySqlAztemarketRootConnectionString),
	ConsoleLogChannel: LogEventsChannel,
}

var WalletService = walletServices.WalletService{
	WalletsRepository: repositories.NewWalletsRepository(MySqlAztemarketRootConnectionString),
	StockRepository:   repositories.NewStockRepository(MySqlAztemarketRootConnectionString),
	ConsoleLogChannel: LogEventsChannel,
}
