package remoteOnReadyEvent

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Called once the Discord servers confirm a succesful connection.
func DefaultHandler(s *discordgo.Session, event *discordgo.Ready) {
	fmt.Println("YOOOO")
}
