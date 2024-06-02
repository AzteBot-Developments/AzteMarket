package slashCmdMarketHandlers

import (
	"github.com/RazvanBerbece/AzteMarket/pkg/interaction"
	"github.com/bwmarrin/discordgo"
)

func HandleSlashViewMarket(s *discordgo.Session, i *discordgo.InteractionCreate) {

	// TODO

	// Final response to interaction
	interaction.SendSimpleEmbedSlashResponse(s, i.Interaction, "Not available yet.")

}
