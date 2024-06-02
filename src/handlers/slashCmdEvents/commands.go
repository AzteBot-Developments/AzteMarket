package slashCmdEvents

import (
	slashCmdMarketHandlers "github.com/RazvanBerbece/AzteMarket/src/handlers/slashCmdEvents/handlers/market"
	slashCmdUtilHandlers "github.com/RazvanBerbece/AzteMarket/src/handlers/slashCmdEvents/handlers/utils"
	slashCmdWalletHandlers "github.com/RazvanBerbece/AzteMarket/src/handlers/slashCmdEvents/handlers/wallet"
	"github.com/bwmarrin/discordgo"
)

var DefinedSlashCommands = []*discordgo.ApplicationCommand{
	{
		Name:        "ping",
		Description: "Basic ping slash interaction for the AzteMarket.",
	},
	{
		Name:        "market-add-stock",
		Description: "Adds a new stock item to sell on the OTA marketplace.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "name",
				Description: "The name of the new stock item to put on the market for sale.",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "details",
				Description: "More details about the new stock item.",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "cost",
				Description: "The cost in AzteCoins for the new stock item.",
				Required:    true,
			},
		},
	},
	{
		Name:        "wallet",
		Description: "Displays the command's author wallet status (funds, details, etc.).",
	},
}

var RegisteredSlashCommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"ping":             slashCmdUtilHandlers.HandleSlashPing,
	"market-add-stock": slashCmdMarketHandlers.HandleSlashAddStock,
	"wallet":           slashCmdWalletHandlers.HandleSlashWallet,
}
