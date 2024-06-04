package slashCmdMarketHandlers

import (
	"fmt"

	"github.com/RazvanBerbece/AzteMarket/pkg/interaction"
	sharedRuntime "github.com/RazvanBerbece/AzteMarket/src/shared/runtime"
	"github.com/bwmarrin/discordgo"
)

func HandleSlashClearMarket(s *discordgo.Session, i *discordgo.InteractionCreate) {

	deletedCount, err := sharedRuntime.MarketplaceService.ClearMarket()
	if err != nil {
		interaction.ErrorEmbedResponseEdit(s, i.Interaction, err.Error())
		return
	}

	interaction.SendSimpleEmbedSlashResponse(s, i.Interaction, fmt.Sprintf("Successfully cleared `%d` items off the AzteMarket.", deletedCount))

}
