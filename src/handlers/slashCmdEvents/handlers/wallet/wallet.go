package slashCmdWalletHandlers

import (
	"github.com/bwmarrin/discordgo"
)

func HandleSlashWallet(s *discordgo.Session, i *discordgo.InteractionCreate) {

	// authorUserId := i.Member.User.ID
	// wallet, err := sharedRuntime.WalletService.GetWalletForUser(authorUserId)
	// if err != nil {
	// 	interaction.SendErrorEmbedResponse(s, i.Interaction, err.Error())
	// 	go logUtils.PublishDiscordLogErrorEvent(sharedRuntime.LogEventsChannel, s, "Debug", sharedConfig.DiscordChannelTopicPairs, err.Error())
	// 	return
	// }

	// // Final response to interaction
	// interaction.SendSimpleEmbedSlashResponse(s, i.Interaction, wallet.Funds)
}
