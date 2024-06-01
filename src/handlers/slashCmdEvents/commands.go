package slashCmdEvents

import (
	slashCmdUtilHandlers "github.com/RazvanBerbece/AzteMarket/src/handlers/slashCmdEvents/handlers/utils"
	"github.com/bwmarrin/discordgo"
)

var DefinedSlashCommands = []*discordgo.ApplicationCommand{
	{
		Name:        "ping",
		Description: "Basic ping slash interaction for the AzteBot.",
	},
}

var RegisteredSlashCommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"ping": slashCmdUtilHandlers.HandleSlashPing,
}
