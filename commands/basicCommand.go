package commands

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func init() {
	log.Printf("Registering command: " + basicCommandName)
	Commands = append(Commands, basicCommand.discordCommand)
	CommandHandlers[basicCommandName] = basicCommand.commandHandler
}

var basicCommand = command{
	name:           basicCommandName,
	discordCommand: basicDiscordCommand,
	commandHandler: handleBasicCommand,
}

// this is the command name
var basicCommandName = "basic-command"

// this is the metadata for the commands
var basicDiscordCommand = &discordgo.ApplicationCommand{
	Name: basicCommandName,
	// All commands and options must have a description
	// Commands/options without description will fail the registration
	// of the discordCommand.
	Description: "Basic discordCommand",
}

// this is what the command actually does
func handleBasicCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
_:
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Hey there! Congratulations, you just executed your first slash discordCommand",
		},
	})
}
