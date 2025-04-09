package commands

import (
	"github.com/bwmarrin/discordgo"
)

func init() {
	registerCommand(submitSpeedRunCommand)
}

var submitSpeedRunCommand = command{
	name:           submitSpeedrunCommandName,
	discordCommand: submitSpeedrunDiscordCommand,
	commandHandler: handleSubmitSpeedrunCommand,
}

// this is the command name
var submitSpeedrunCommandName = "submit-speedrun"

// this is the metadata for the commands
var submitSpeedrunDiscordCommand = &discordgo.ApplicationCommand{
	Name: submitSpeedrunCommandName,
	// All commands and options must have a description
	// Commands/options without description will fail the registration
	// of the discordCommand.
	Description: "Submit a speedrun for verification",
}

// this is what the command actually does
func handleSubmitSpeedrunCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Hi, this is currently WIP",
		},
	})
}
