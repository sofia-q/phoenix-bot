package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"phoenixbot/bot/model"
)

func init() {
	registerCommand(submitSpeedRunCommand)
	for i := model.SwordAndShield; i < model.Bow; i++ {
		weaponChoices = append(weaponChoices, &discordgo.ApplicationCommandOptionChoice{
			Name:  model.WeaponType.String(model.WeaponType(i)),
			Value: model.WeaponType.GetWeaponHandle(model.WeaponType(i)),
		})
	}
}

var weaponChoices []*discordgo.ApplicationCommandOptionChoice

var submitSpeedRunCommand = command{
	name: submitSpeedrunCommandName,
	// this is the metadata for the commands
	commandMetadata: &discordgo.ApplicationCommand{
		Name: submitSpeedrunCommandName,
		// All commands and options must have a description
		// Commands/options without description will fail the registration
		// of the command.
		Description: "Submit a speedrun for verification",
		Options: []*discordgo.ApplicationCommandOption{

			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "weapon-type-option",
				Description: "Weapon Type",
				Required:    true,
				Choices:     weaponChoices,
			},
		},
	},
	commandHandler: handleSubmitSpeedrunCommand,
}

// this is the command name
var submitSpeedrunCommandName = "submit-speedrun"

// this is what the command actually does
func handleSubmitSpeedrunCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Access options in the order provided by the user.
	options := i.ApplicationCommandData().Options

	// convert the slice into a map
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	response := "weapon type entered: "

	// Get the value from the option map.
	// When the option exists, ok = true
	if option, ok := optionMap["weapon-type-option"]; ok {
		response += option.StringValue()
	}

	_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		// Ignore type for now, they will be discussed in "responses"
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf(
				response,
			),
		},
	})
}
