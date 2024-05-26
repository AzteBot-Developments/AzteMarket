package globalConfig

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var _ = godotenv.Load(".env")

var DiscordBotToken = os.Getenv("DISCORD_BOT_TOKEN")
var DiscordBotAppId = os.Getenv("DISCORD_APP_ID")

var DiscordMainGuildId = os.Getenv("DISCORD_MAIN_GUILD_ID")

var MySqlAztebotRootConnectionString = os.Getenv("DB_AZTEBOT_ROOT_CONNSTRING") // in MySQL format (i.e. `root_username:root_password@<unix/tcp>(host:port)/db_name`)
var MySqlAztemarketRootConnectionString = os.Getenv("DB_AZTEMARKET_ROOT_CONNSTRING")

var DiscordChannelTopicPairs = strings.Split(os.Getenv("DISCORD_CHANNEL_TOPICS"), ",")
