package env

import (
	"flag"
	"fmt"
	"os"
	"phoenixbot/bot/db"

	"github.com/joho/godotenv"
)

var (
	GuildID        = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
	BotToken       = flag.String("token", "", "Bot access token")
	RemoveCommands = flag.Bool("rmcmd", false, "Remove all commands after shutdowning or not")
)

func init() {

	token := LoadVar("BOT_TOKEN")
	if token == "" {
		fmt.Println("No BOT_TOKEN environment variable found")
		return
	}
	_ = flag.Set("token", token)

	guild := LoadVar("GUILD_ID")
	if guild == "" {
		fmt.Println("No guild ID environment variable found")

	} else {
		_ = flag.Set("guild", guild)
	}
	flag.Parse()

	databaseIp := LoadVar("DATABASE_IP")
	if databaseIp == "" {
		fmt.Println("DATABASE_IP environment variable not found")
		return
	}
	databaseUser := LoadVar("DATABASE_USER")
	if databaseUser == "" {
		fmt.Println("DATABASE_IP environment variable not found")
		return
	}
	databasePw := LoadVar("DATABASE_PW")
	if databasePw == "" {
		fmt.Println("DATABASE_IP environment variable not found")
		return
	}

	db.ConnectDB(databaseUser, databasePw, databaseIp)
}

func LoadVar(key string) string {
	// Attempt to load .env file
	_ = godotenv.Load()

	// Return the environment variable
	return os.Getenv(key)
}
