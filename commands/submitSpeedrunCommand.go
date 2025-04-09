package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"phoenixbot/bot/model"
)

func init() {
	registerCommand(submitSpeedRunCommand)
}

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
				Name:        "weapon-type",
				Description: "Weapon Type",
				Required:    true,
				Choices:     weaponTypeChoices,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "minutes",
				Description: "Minutes",
				Required:    true,
				MaxValue:    49,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "seconds",
				Description: "Seconds",
				Required:    true,
				MaxValue:    59,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "proof",
				Description: "Link to a screenshot as proof",
				Required:    true,
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
	if option, ok := optionMap["weapon-type"]; ok {
		response += option.StringValue()
	}
	response += " time taken: "
	if option, ok := optionMap["minutes"]; ok {
		response += fmt.Sprintf(" %02d:", option.IntValue())
	}
	if option, ok := optionMap["seconds"]; ok {
		response += fmt.Sprintf("%02d ", option.IntValue())
	}
	response += " link: "
	if option, ok := optionMap["proof"]; ok {
		response += option.StringValue()
	}
	response += " userID: " + i.Member.User.ID

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

var weaponTypeChoices = []*discordgo.ApplicationCommandOptionChoice{
	{
		Name:  model.WeaponType.String(model.SwordAndShield),
		Value: model.WeaponType.GetWeaponHandle(model.SwordAndShield),
	},
	{
		Name:  model.WeaponType.String(model.DualBlades),
		Value: model.WeaponType.GetWeaponHandle(model.DualBlades),
	},
	{
		Name:  model.WeaponType.String(model.GreatSword),
		Value: model.WeaponType.GetWeaponHandle(model.GreatSword),
	},
	{
		Name:  model.WeaponType.String(model.LongSword),
		Value: model.WeaponType.GetWeaponHandle(model.LongSword),
	},
	{
		Name:  model.WeaponType.String(model.Hammer),
		Value: model.WeaponType.GetWeaponHandle(model.Hammer),
	},
	{
		Name:  model.WeaponType.String(model.HuntingHorn),
		Value: model.WeaponType.GetWeaponHandle(model.HuntingHorn),
	},
	{
		Name:  model.WeaponType.String(model.Lance),
		Value: model.WeaponType.GetWeaponHandle(model.Lance),
	},
	{
		Name:  model.WeaponType.String(model.GunLance),
		Value: model.WeaponType.GetWeaponHandle(model.GunLance),
	},
	{
		Name:  model.WeaponType.String(model.SwitchAxe),
		Value: model.WeaponType.GetWeaponHandle(model.SwitchAxe),
	},
	{
		Name:  model.WeaponType.String(model.ChargeBlade),
		Value: model.WeaponType.GetWeaponHandle(model.ChargeBlade),
	},
	{
		Name:  model.WeaponType.String(model.InsectGlaive),
		Value: model.WeaponType.GetWeaponHandle(model.InsectGlaive),
	},
	{
		Name:  model.WeaponType.String(model.LightBowgun),
		Value: model.WeaponType.GetWeaponHandle(model.LightBowgun),
	},
	{
		Name:  model.WeaponType.String(model.HeavyBowgun),
		Value: model.WeaponType.GetWeaponHandle(model.HeavyBowgun),
	},
	{
		Name:  model.WeaponType.String(model.Bow),
		Value: model.WeaponType.GetWeaponHandle(model.Bow),
	},
}
