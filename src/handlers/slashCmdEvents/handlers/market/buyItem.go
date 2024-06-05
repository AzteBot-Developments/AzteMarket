package slashCmdMarketHandlers

import (
	"fmt"

	"github.com/RazvanBerbece/AzteMarket/pkg/interaction"
	"github.com/bwmarrin/discordgo"
)

func HandleSlashBuyItem(s *discordgo.Session, i *discordgo.InteractionCreate) {

	// Retrieve input args from command
	itemId := i.ApplicationCommandData().Options[0].StringValue()

	// Input validation
	if len(itemId) <= 0 || len(itemId) > 40 {
		interaction.SendErrorEmbedResponse(s, i.Interaction, fmt.Sprintf("Invalid input argument (term: `%s`)", i.ApplicationCommandData().Options[0].Name))
		return
	}

	// Final response to interaction
	interaction.SendSimpleEmbedSlashResponse(s, i.Interaction, fmt.Sprintf("Bought item with ID `%s`", itemId))

}
