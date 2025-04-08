package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"phoenixbot/bot/commandhandlers"
	"phoenixbot/bot/commands"
	"phoenixbot/bot/db"
	"phoenixbot/bot/env"

	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"
)

// Bot parameters
var (
	GuildID        = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
	BotToken       = flag.String("token", "", "Bot access token")
	RemoveCommands = flag.Bool("rmcmd", true, "Remove all commands after shutdowning or not")
)

var Db gorm.DB

var s *discordgo.Session

func init() {

	token := env.LoadVar("BOT_TOKEN")
	if token == "" {
		fmt.Println("No BOT_TOKEN environment variable found")
		return
	}
	flag.Set("token", token)

	guild := env.LoadVar("GUILD_ID")
	if guild == "" {
		fmt.Println("No guild ID environment variable found")

	} else {
		flag.Set("guild", guild)
	}
	flag.Parse()

	var err error
	s, err = discordgo.New("Bot " + *BotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandhandlers.CommandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

func main() {

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	registeredCommands := addCommands()

	db.ConnectDB()

	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	if *RemoveCommands {
		removeCommands(registeredCommands)
	}

	log.Println("Gracefully shutting down.")
}

func removeCommands(registeredCommands []*discordgo.ApplicationCommand) {
	log.Println("Removing commands...")
	// // We need to fetch the commands, since deleting requires the command ID.
	// // We are doing this from the returned commands on line 375, because using
	// // this will delete all the commands, which might not be desirable, so we
	// // are deleting only the commands that we added.
	// registeredCommands, err := s.ApplicationCommands(s.State.User.ID, *GuildID)
	// if err != nil {
	// 	log.Fatalf("Could not fetch registered commands: %v", err)
	// }

	for _, v := range registeredCommands {
		err := s.ApplicationCommandDelete(s.State.User.ID, *GuildID, v.ID)
		if err != nil {
			log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
		}
	}
}

func addCommands() []*discordgo.ApplicationCommand {
	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands.Commands))
	for i, v := range commands.Commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, *GuildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}
	return registeredCommands
}
