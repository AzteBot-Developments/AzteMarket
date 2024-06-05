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
		Name:        "market-buy-item",
		Description: "Buys an item from the AzteMarket using the funds available in the user's wallet.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "id",
				Description: "The ID of the item to buy from the market.",
				Required:    true,
			},
		},
	},
	{
		Name:        "wallet",
		Description: "Displays the command's author wallet status (funds, details, etc.).",
	},
	{
		Name:        "wallet-create",
		Description: "Creates an AzteMarket wallet for the author of the command !",
	},
	{
		Name:        "wallet-delete",
		Description: "Deletes the command author's AzteMarket wallet.",
	},
	{
		Name:        "wallet-send-funds",
		Description: "Sends AzteCoins to another member's wallet.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "id",
				Description: "The wallet ID to send the funds to.",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "funds",
				Description: "The amount of funds to send. (max. 50000 per transaction)",
				Required:    true,
			},
		},
	},
}

var RegisteredSlashCommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"ping":              slashCmdUtilHandlers.HandleSlashPing,
	"market":            slashCmdMarketHandlers.HandleSlashViewMarket,
	"market-clear":      slashCmdMarketHandlers.HandleSlashClearMarket,
	"market-see-item":   slashCmdMarketHandlers.HandleSlashViewItemOnMarket,
	"market-add-stock":  slashCmdMarketHandlers.HandleSlashAddStock,
	"market-buy-item":   slashCmdMarketHandlers.HandleSlashBuyItem,
	"wallet":            slashCmdWalletHandlers.HandleSlashWallet,
	"wallet-create":     slashCmdWalletHandlers.HandleSlashCreateWallet,
	"wallet-delete":     slashCmdWalletHandlers.HandleSlashDeleteWallet,
	"wallet-send-funds": slashCmdWalletHandlers.HandleSlashSendFundsFromWallet,
}
