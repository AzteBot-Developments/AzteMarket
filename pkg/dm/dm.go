package dm

import (
	"fmt"
	"time"

	"github.com/RazvanBerbece/AzteMarket/pkg/embed"
	"github.com/RazvanBerbece/AzteMarket/src/libs/models/domain"
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

func SendDirectSimpleEmbedToMember(s *discordgo.Session, userId string, title string, text string, color int) error {

	simpleEmbed := embed.SimpleEmbed(title, text, color)

	errDm := DmEmbedUser(s, userId, *simpleEmbed[0])
	if errDm != nil {
		fmt.Printf("Error sending embed DM to member with UID %s: %v\n", userId, errDm)
		return errDm
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

func SendDirectComplexEmbedToMember(s *discordgo.Session, userId string, embed embed.Embed, actionsRow discordgo.ActionsRow, pageSize int, runtimeMap map[string]domain.EmbedData) error {

	originalAllFields := make([]*discordgo.MessageEmbedField, len(embed.Fields))
	copy(originalAllFields, embed.Fields)

	// Only show fields from page 1 in the beginning
	pages := (len(originalAllFields) + pageSize - 1) / pageSize
	embed.Fields = embed.Fields[0:pageSize]
	embed.Footer = &discordgo.MessageEmbedFooter{
		Text: fmt.Sprintf("Page %d / %d", 1, pages),
	}
	msg, err := DmEmbedComplexUser(s, userId, *embed.MessageEmbed, actionsRow)
	if err != nil {
		fmt.Printf("Error sending embed DM to member with UID %s: %v\n", userId, err)
		return err
	}

	// Keep paginated embeds in-memory to enable handling on button presses
	runtimeMap[msg.ID] = domain.EmbedData{
		ChannelId:   msg.ChannelID,
		FieldData:   &originalAllFields, // all fields
		CurrentPage: 1,                  // same for all complex paginated embeds
		Timestamp:   float64(time.Now().Unix()),
	}

	return nil
}

func DmEmbedComplexUser(s *discordgo.Session, userId string, embed discordgo.MessageEmbed, actionsRow discordgo.ActionsRow) (*discordgo.Message, error) {

	channel, err := s.UserChannelCreate(userId)
	if err != nil {
		fmt.Println("error creating embed DM channel: ", err)
		return nil, err
	}

	msg, err := s.ChannelMessageSendComplex(channel.ID, &discordgo.MessageSend{
		Embed:      &embed,
		Components: []discordgo.MessageComponent{actionsRow},
	})
	if err != nil {
		fmt.Printf("Error sending complex DM to user: %v\n", err)
		return nil, err
	}

	return msg, nil

}
