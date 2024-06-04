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
		Name:        "market",
		Description: "Open the OTA market to see what benefits are available for buying.",
	},
	{
		Name:        "market-clear",
		Description: "Restricted command. Resets the OTA market bringing it to its initial state.",
	},
	{
		Name:        "market-see-item",
		Description: "Sees details about a certain item currently on sale on the OTA market.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "id",
				Description: "The id of the stock item to see from the market.",
				Required:    true,
			},
		},
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
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "max-items",
				Description: "The total amount of items of this kind that can be purchased by the community.",
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
	"market":           slashCmdMarketHandlers.HandleSlashViewMarket,
	"market-clear":     slashCmdMarketHandlers.HandleSlashClearMarket,
	"market-see-item":  slashCmdMarketHandlers.HandleSlashViewItemOnMarket,
	"market-add-stock": slashCmdMarketHandlers.HandleSlashAddStock,
	"wallet":           slashCmdWalletHandlers.HandleSlashWallet,
}
