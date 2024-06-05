package dm

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func DmUser(s *discordgo.Session, userId string, content string) error {

	channel, err := s.UserChannelCreate(userId)
	if err != nil {
		fmt.Println("error creating DM channel: ", err)
		return err
	}

	_, err = s.ChannelMessageSend(channel.ID, content)
	if err != nil {
		fmt.Println("error sending DM message: ", err)
		return err
	}

	return nil

}

func DmEmbedUser(s *discordgo.Session, userId string, embed discordgo.MessageEmbed) error {

	channel, err := s.UserChannelCreate(userId)
	if err != nil {
		fmt.Println("error creating embed DM channel: ", err)
		return err
	}

	_, err = s.ChannelMessageSendEmbed(channel.ID, &embed)
	if err != nil {
		fmt.Println("error sending embed DM message: ", err)
		return err
	}

	return nil

}
