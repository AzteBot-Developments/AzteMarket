package sharedConfig

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var _ = godotenv.Load(".env")

var DiscordBotToken = os.Getenv("DISCORD_BOT_TOKEN")
var DiscordBotAppId = os.Getenv("DISCORD_APP_ID")

var DiscordMainGuildId = os.Getenv("DISCORD_MAIN_GUILD_ID")

var DiscordChannelTopicPairs = strings.Split(os.Getenv("DISCORD_CHANNEL_TOPICS"), ",")
