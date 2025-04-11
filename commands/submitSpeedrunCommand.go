package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"net/url"
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
	var newSpeedrun = model.Speedrun{}

	// Get the value from the option map.
	// When the option exists, ok = true
	if option, ok := optionMap["weapon-type"]; ok {
		newSpeedrun.WeaponType = option.StringValue()
	}
	if option, ok := optionMap["minutes"]; ok {
		newSpeedrun.TimeInSeconds = int(option.IntValue() * 60)
	}
	if option, ok := optionMap["seconds"]; ok {
		newSpeedrun.TimeInSeconds += int(option.IntValue())
	}
	if option, ok := optionMap["proof"]; ok {
		newSpeedrun.ProofLink = option.StringValue()
		if _, err := url.ParseRequestURI(option.StringValue()); err != nil {
			_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				// Ignore type for now, they will be discussed in "responses"
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Error! Invalid image URL!",
					Flags:   discordgo.MessageFlagsEphemeral,
				},
			})
			return
		}
	}
	newSpeedrun.UserId = i.Member.User.ID
	newSpeedrun.IsVerified = false
	newSpeedrun.Season = 1
	_ = model.Db.Create(&newSpeedrun)
	response := "weapon type entered: " + newSpeedrun.WeaponType
	response += " time taken: "
	response += fmt.Sprintf(" %02d:", newSpeedrun.TimeInSeconds/60)
	response += fmt.Sprintf("%02d ", newSpeedrun.TimeInSeconds%60)
	response += " link: "
	response += newSpeedrun.ProofLink
	response += " user: " + "<@" + newSpeedrun.UserId + ">"
	response += " new run ID: " + newSpeedrun.ID.String()

	_, err := s.ChannelMessageSendComplex("1358151701420577009", &discordgo.MessageSend{
		Content: response + " note: this is the message in the admin only verify channel",
		Components: []discordgo.MessageComponent{
			// ActionRow is a container of all buttons within the same row.
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						// Label is what the user will see on the button.
						Label: "Verify",
						// Style provides coloring of the button. There are not so many styles tho.
						Style: discordgo.SuccessButton,
						// Disabled allows bot to disable some buttons for users.
						Disabled: false,
						// CustomID is a thing telling Discord which data to send when this button will be pressed.
						CustomID: "verify_button_yes",
					},
				},
			},
		},
		AllowedMentions: &discordgo.MessageAllowedMentions{
			Parse: []discordgo.AllowedMentionType{},
			Users: []string{},
		},
	})
	if err != nil {
		log.Printf(err.Error())
	}
	_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		// Ignore type for now, they will be discussed in "responses"
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf(
				response,
			),
			AllowedMentions: &discordgo.MessageAllowedMentions{
				Parse: []discordgo.AllowedMentionType{},
				Users: []string{},
			},
		},
	})
}

var weaponTypeChoices = []*discordgo.ApplicationCommandOptionChoice{
	{
		Name:  model.WeaponType.String(model.SwordAndShield),
		Value: model.WeaponType.String(model.SwordAndShield),
	},
	{
		Name:  model.WeaponType.String(model.DualBlades),
		Value: model.WeaponType.String(model.DualBlades),
	},
	{
		Name:  model.WeaponType.String(model.GreatSword),
		Value: model.WeaponType.String(model.GreatSword),
	},
	{
		Name:  model.WeaponType.String(model.LongSword),
		Value: model.WeaponType.String(model.LongSword),
	},
	{
		Name:  model.WeaponType.String(model.Hammer),
		Value: model.WeaponType.String(model.Hammer),
	},
	{
		Name:  model.WeaponType.String(model.HuntingHorn),
		Value: model.WeaponType.String(model.HuntingHorn),
	},
	{
		Name:  model.WeaponType.String(model.Lance),
		Value: model.WeaponType.String(model.Lance),
	},
	{
		Name:  model.WeaponType.String(model.GunLance),
		Value: model.WeaponType.String(model.GunLance),
	},
	{
		Name:  model.WeaponType.String(model.SwitchAxe),
		Value: model.WeaponType.String(model.SwitchAxe),
	},
	{
		Name:  model.WeaponType.String(model.ChargeBlade),
		Value: model.WeaponType.String(model.ChargeBlade),
	},
	{
		Name:  model.WeaponType.String(model.InsectGlaive),
		Value: model.WeaponType.String(model.InsectGlaive),
	},
	{
		Name:  model.WeaponType.String(model.LightBowgun),
		Value: model.WeaponType.String(model.LightBowgun),
	},
	{
		Name:  model.WeaponType.String(model.HeavyBowgun),
		Value: model.WeaponType.String(model.HeavyBowgun),
	},
	{
		Name:  model.WeaponType.String(model.Bow),
		Value: model.WeaponType.String(model.Bow),
	},
}
