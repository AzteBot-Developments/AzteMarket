package remoteOnMessageEvent

import "github.com/bwmarrin/discordgo"

func Db(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Content == "db" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

}
