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
		Description: "Gateway to various commands related to the AzteMarket.",
	},
	{
		Name:        "wallet",
		Description: "Displays a member's wallet status (funds, details, etc.).",
	},
}

var RegisteredSlashCommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"ping":   slashCmdUtilHandlers.HandleSlashPing,
	"market": slashCmdMarketHandlers.HandleSlashMarket,
	"wallet": slashCmdWalletHandlers.HandleSlashWallet,
}
