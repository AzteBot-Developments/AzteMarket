package slashCmdEvents

import (
	slashCmdMarketHandlers "github.com/RazvanBerbece/AzteMarket/src/handlers/slashCmdEvents/handlers/market"
	slashCmdUtilHandlers "github.com/RazvanBerbece/AzteMarket/src/handlers/slashCmdEvents/handlers/utils"
	slashCmdWalletHandlers "github.com/RazvanBerbece/AzteMarket/src/handlers/slashCmdEvents/handlers/wallet"
	"github.com/bwmarrin/discordgo"
)

// Command tuning and validation
var TransferReferenceMinLength = 5
var TransferReferenceMaxLength = 256

var ItemNameMinLength = 5
var ItemNameMaxLength = 128

var MultiplierMinValue = 1.0

var DefinedSlashCommands = []*discordgo.ApplicationCommand{
	{
		Name:        "ping",
		Description: "Basic ping slash interaction for the AzteMarket.",
	},
	{
		Name:        "help",
		Description: "Summary of the AzteMarket slash commands.",
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
		Name:        "market-add-units",
		Description: "Adds units of a certain item to be available for sale on the market.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "id",
				Description: "The ID of the item to supply with extra units.",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "multiplier",
				Description: "How many units to supply.",
				Required:    true,
				MinValue:    &MultiplierMinValue,
				MaxValue:    50000,
			},
		},
	},
	{
		Name:        "market-remove-units",
		Description: "Removes units of a certain item so they're not available for sale on the market anymore.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "id",
				Description: "The ID of the item to remove extra units from.",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "multiplier",
				Description: "How many units to remove.",
				Required:    true,
				MinValue:    &MultiplierMinValue,
				MaxValue:    50000,
			},
		},
	},
	{
		Name:        "market-remove-item",
		Description: "Removes an item and all its vailable units from the AzteMarket.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "id",
				Description: "The ID of the item to remove from the market.",
				Required:    true,
			},
		},
	},
	{
		Name:        "wallet",
		Description: "Displays the command's author wallet status (funds, details, etc.).",
	},
	{
		Name:        "wallet-view",
		Description: "Privileged command to see a member's wallet status (funds, details, etc.).",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "id",
				Description: "The wallet / user ID to see the wallet status for.",
				Required:    true,
			},
		},
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
				Description: "The wallet / user ID to send the funds to.",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "funds",
				Description: "The amount of funds to send. (max. 50000 per transaction)",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "reference",
				Description: "A message to attach to the transfer",
				Required:    false,
				MinLength:   &TransferReferenceMinLength,
				MaxLength:   TransferReferenceMaxLength,
			},
		},
	},
	{
		Name:        "wallet-use-item",
		Description: "Uses an item from the owner's wallet and removes it from their inventory.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "id",
				Description: "The wallet ID from which to use the item.",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "item-name",
				Description: "The name of the item to use from the target wallet",
				Required:    true,
				MinLength:   &ItemNameMinLength,
				MaxLength:   ItemNameMaxLength,
			},
		},
	},
}

var RegisteredSlashCommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"ping":                slashCmdUtilHandlers.HandleSlashPing,
	"help":                HandleSlashAzteMarketHelp,
	"market":              slashCmdMarketHandlers.HandleSlashViewMarket,
	"market-clear":        slashCmdMarketHandlers.HandleSlashClearMarket,
	"market-see-item":     slashCmdMarketHandlers.HandleSlashViewItemOnMarket,
	"market-add-stock":    slashCmdMarketHandlers.HandleSlashAddStock,
	"market-buy-item":     slashCmdMarketHandlers.HandleSlashBuyItem,
	"market-remove-item":  slashCmdMarketHandlers.HandleSlashRemoveItemFromMarket,
	"market-add-units":    slashCmdMarketHandlers.HandleSlashAddAvailableUnitsForItem,
	"market-remove-units": slashCmdMarketHandlers.HandleSlashRemoveAvailableUnitsForItem,
	"wallet":              slashCmdWalletHandlers.HandleSlashWallet,
	"wallet-view":         slashCmdWalletHandlers.HandleSlashViewWallet,
	"wallet-create":       slashCmdWalletHandlers.HandleSlashCreateWallet,
	"wallet-delete":       slashCmdWalletHandlers.HandleSlashDeleteWallet,
	"wallet-send-funds":   slashCmdWalletHandlers.HandleSlashSendFundsFromWallet,
	"wallet-use-item":     slashCmdWalletHandlers.HandleSlashUseItemFromWallet,
}
