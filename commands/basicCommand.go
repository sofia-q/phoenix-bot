package commands

import (
	"github.com/bwmarrin/discordgo"
)

func init() {
	registerCommand(basicCommand)
}

var basicCommand = command{
	name: basicCommandName,
	//this is the metadata for the commands
	discordCommand: &discordgo.ApplicationCommand{
		Name: basicCommandName,
		// All commands and options must have a description
		// Commands/options without description will fail the registration
		// of the command.
		Description: "Basic command",
	},
	commandHandler: handleBasicCommand,
}

// this is the command name
var basicCommandName = "basic-command"

// this is what the command actually does
func handleBasicCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Hey there! Congratulations, you just executed your first slash command",
		},
	})
}
