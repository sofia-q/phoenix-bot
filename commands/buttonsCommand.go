package commands

import (
	"github.com/bwmarrin/discordgo"
)

func init() {
	registerCommand(buttonCommand)
}

var buttonCommand = command{
	name: buttonCommandName,
	//this is the metadata for the commands
	commandMetadata: &discordgo.ApplicationCommand{
		Name: buttonCommandName,
		// All commands and options must have a description
		// Commands/options without description will fail the registration
		// of the command.
		Description: "Test the buttons if you got courage",
	},
	commandHandler: handleButtonCommand,
}

// this is the command name
var buttonCommandName = "buttons"

// this is what the command actually does
func handleButtonCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Are you comfortable with buttons and other message components?",
			Flags:   discordgo.MessageFlagsEphemeral,
			// Buttons and other components are specified in Components field.
			Components: []discordgo.MessageComponent{
				// ActionRow is a container of all buttons within the same row.
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							// Label is what the user will see on the button.
							Label: "Yes",
							// Style provides coloring of the button. There are not so many styles tho.
							Style: discordgo.SuccessButton,
							// Disabled allows bot to disable some buttons for users.
							Disabled: false,
							// CustomID is a thing telling Discord which data to send when this button will be pressed.
							CustomID: "fd_yes",
						},
						discordgo.Button{
							Label:    "No",
							Style:    discordgo.DangerButton,
							Disabled: false,
							CustomID: "fd_no",
						},
						discordgo.Button{
							Label:    "I don't know",
							Style:    discordgo.LinkButton,
							Disabled: false,
							// Link buttons don't require CustomID and do not trigger the gateway/HTTP event
							URL: "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
							Emoji: &discordgo.ComponentEmoji{
								Name: "🤷",
							},
						},
					},
				},
				// The message may have multiple actions rows.
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "Discord Developers server",
							Style:    discordgo.LinkButton,
							Disabled: false,
							URL:      "https://discord.gg/discord-developers",
						},
					},
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}
}
