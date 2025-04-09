package env

import (
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	GuildID        = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
	BotToken       = flag.String("token", "", "Bot access token")
	RemoveCommands = flag.Bool("rmcmd", false, "Remove all commands after shutdowning or not")
	DatabaseIp     = ""
	DatabaseUser   = ""
	DatabasePw     = ""
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

	DatabaseIp = LoadVar("DATABASE_IP")
	if DatabaseIp == "" {
		fmt.Println("DATABASE_IP environment variable not found")
		return
	}
	DatabaseUser = LoadVar("DATABASE_USER")
	if DatabaseUser == "" {
		fmt.Println("DATABASE_IP environment variable not found")
		return
	}
	DatabasePw = LoadVar("DATABASE_PW")
	if DatabasePw == "" {
		fmt.Println("DATABASE_IP environment variable not found")
		return
	}

}

func LoadVar(key string) string {
	// Attempt to load .env file
	_ = godotenv.Load()

	// Return the environment variable
	return os.Getenv(key)
}
