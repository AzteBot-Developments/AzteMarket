package slashCmdUtilHandlers

import "github.com/bwmarrin/discordgo"

func HandleSlashPing(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "AzteMarket Pong!",
		},
	})
}
