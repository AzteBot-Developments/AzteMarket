package globalConfig

import (
	"os"

	"github.com/joho/godotenv"
)

var _ = godotenv.Load(".env")

var DiscordAztebotToken = os.Getenv("DISCORD_BOT_TOKEN")
var DiscordAztebotAppId = os.Getenv("DISCORD_APP_ID")

var DiscordMainGuildId = os.Getenv("DISCORD_MAIN_GUILD_ID")

var MySqlRootConnectionString = os.Getenv("DB_ROOT_CONNSTRING") // in MySQL format (i.e. `root_username:root_password@<unix/tcp>(host:port)/db_name`)
