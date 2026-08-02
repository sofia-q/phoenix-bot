package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"net/url"
	"phoenixbot/bot/db"
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
	var newSpeedrun = db.Speedrun{}

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
		if _, parseUrlErr := url.ParseRequestURI(option.StringValue()); parseUrlErr != nil {
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
	newSpeedrun.GuildID = i.GuildID
	saveErr := newSpeedrun.Save()
	if saveErr != nil {
		log.Println(saveErr)
	}
	var runInfo = &discordgo.MessageEmbed{
		Title: "Speedrun submitted",
		Type:  "rich",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "User",
				Value: "<@" + newSpeedrun.UserId + ">",
			},
			{
				Name:  "Weapon Type",
				Value: newSpeedrun.WeaponType,
			},
			{
				Name:  "Time Taken",
				Value: fmt.Sprintf("%02d:%02d", newSpeedrun.TimeInSeconds/60, newSpeedrun.TimeInSeconds%60),
			},
			{
				Name: "Proof:",
			},
		},
		Image: &discordgo.MessageEmbedImage{
			URL: newSpeedrun.ProofLink,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: newSpeedrun.ID.String(),
		},
	}

	_, err := s.ChannelMessageSendComplex("1358151701420577009", &discordgo.MessageSend{
		Content: "note: this is the message in the admin only verify channel",
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
					discordgo.Button{
						// Label is what the user will see on the button.
						Label: "Un-Verify (WIP)",
						// Style provides coloring of the button. There are not so many styles tho.
						Style: discordgo.DangerButton,
						// Disabled allows bot to disable some buttons for users.
						Disabled: true,
						// CustomID is a thing telling Discord which data to send when this button will be pressed.
						CustomID: "fd_no",
					},
				},
			},
		},
		Embeds: []*discordgo.MessageEmbed{runInfo},
	})
	if err != nil {
		log.Println(err.Error())
	}
	_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		// Ignore type for now, they will be discussed in "responses"
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "",
			Flags:   discordgo.MessageFlagsEphemeral,
			Embeds:  []*discordgo.MessageEmbed{runInfo},
			AllowedMentions: &discordgo.MessageAllowedMentions{
				Parse: []discordgo.AllowedMentionType{},
				Users: []string{},
			},
		},
	})
}

var weaponTypeChoices = []*discordgo.ApplicationCommandOptionChoice{
	{
		Name:  db.WeaponType.String(db.SwordAndShield),
		Value: db.WeaponType.String(db.SwordAndShield),
	},
	{
		Name:  db.WeaponType.String(db.DualBlades),
		Value: db.WeaponType.String(db.DualBlades),
	},
	{
		Name:  db.WeaponType.String(db.GreatSword),
		Value: db.WeaponType.String(db.GreatSword),
	},
	{
		Name:  db.WeaponType.String(db.LongSword),
		Value: db.WeaponType.String(db.LongSword),
	},
	{
		Name:  db.WeaponType.String(db.Hammer),
		Value: db.WeaponType.String(db.Hammer),
	},
	{
		Name:  db.WeaponType.String(db.HuntingHorn),
		Value: db.WeaponType.String(db.HuntingHorn),
	},
	{
		Name:  db.WeaponType.String(db.Lance),
		Value: db.WeaponType.String(db.Lance),
	},
	{
		Name:  db.WeaponType.String(db.GunLance),
		Value: db.WeaponType.String(db.GunLance),
	},
	{
		Name:  db.WeaponType.String(db.SwitchAxe),
		Value: db.WeaponType.String(db.SwitchAxe),
	},
	{
		Name:  db.WeaponType.String(db.ChargeBlade),
		Value: db.WeaponType.String(db.ChargeBlade),
	},
	{
		Name:  db.WeaponType.String(db.InsectGlaive),
		Value: db.WeaponType.String(db.InsectGlaive),
	},
	{
		Name:  db.WeaponType.String(db.LightBowgun),
		Value: db.WeaponType.String(db.LightBowgun),
	},
	{
		Name:  db.WeaponType.String(db.HeavyBowgun),
		Value: db.WeaponType.String(db.HeavyBowgun),
	},
	{
		Name:  db.WeaponType.String(db.Bow),
		Value: db.WeaponType.String(db.Bow),
	},
}
