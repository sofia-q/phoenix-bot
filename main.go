package main

import (
	"log"
	"os"
	"os/signal"
	"phoenixbot/bot/commands"
	"phoenixbot/bot/env"
	"phoenixbot/bot/model"

	"github.com/bwmarrin/discordgo"
)

// Bot parameters

var s *discordgo.Session

func init() {

	var err error
	s, err = discordgo.New("Bot " + *env.BotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commands.CommandHandlers[i.ApplicationCommandData().Name]; ok {
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

	registeredCommands := commands.AddCommands(s)

	model.ConnectDB()

	defer func() {
		_ = s.Close()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	if *env.RemoveCommands {
		commands.RemoveCommands(s, registeredCommands)
	}

	log.Println("Gracefully shutting down.")
}
