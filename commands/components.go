package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
	"time"
)

var ComponentsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"fd_no": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Huh. I see, maybe some of these resources might help you?",
				Flags:   discordgo.MessageFlagsEphemeral,
				Components: []discordgo.MessageComponent{
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.Button{
								Emoji: &discordgo.ComponentEmoji{
									Name: "📜",
								},
								Label: "Documentation",
								Style: discordgo.LinkButton,
								URL:   "https://discord.com/developers/docs/interactions/message-components#buttons",
							},
							discordgo.Button{
								Emoji: &discordgo.ComponentEmoji{
									Name: "🔧",
								},
								Label: "Discord developers",
								Style: discordgo.LinkButton,
								URL:   "https://discord.gg/discord-developers",
							},
							discordgo.Button{
								Emoji: &discordgo.ComponentEmoji{
									Name: "🦫",
								},
								Label: "Discord Gophers",
								Style: discordgo.LinkButton,
								URL:   "https://discord.gg/7RuRrVHyXF",
							},
						},
					},
				},
			},
		})
		if err != nil {
			panic(err)
		}
	},
	"fd_yes": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Great! If you wanna know more or just have questions, feel free to visit Discord Devs and Discord Gophers server. " +
					"But now, when you know how buttons work, let's move onto select menus (execute `/selects single`)",
				Flags: discordgo.MessageFlagsEphemeral,
				Components: []discordgo.MessageComponent{
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.Button{
								Emoji: &discordgo.ComponentEmoji{
									Name: "🔧",
								},
								Label: "Discord developers",
								Style: discordgo.LinkButton,
								URL:   "https://discord.gg/discord-developers",
							},
							discordgo.Button{
								Emoji: &discordgo.ComponentEmoji{
									Name: "🦫",
								},
								Label: "Discord Gophers",
								Style: discordgo.LinkButton,
								URL:   "https://discord.gg/7RuRrVHyXF",
							},
						},
					},
				},
			},
		})
		if err != nil {
			panic(err)
		}
	},
	"select": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		var response *discordgo.InteractionResponse

		data := i.MessageComponentData()
		switch data.Values[0] {
		case "go":
			response = &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "This is the way.",
					Flags:   discordgo.MessageFlagsEphemeral,
				},
			}
		default:
			response = &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "It is not the way to go.",
					Flags:   discordgo.MessageFlagsEphemeral,
				},
			}
		}
		err := s.InteractionRespond(i.Interaction, response)
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second) // Doing that so user won't see instant response.
		_, err = s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
			Content: "Anyways, now when you know how to use single select menus, let's see how multi select menus work. " +
				"Try calling `/selects multi` command.",
			Flags: discordgo.MessageFlagsEphemeral,
		})
		if err != nil {
			panic(err)
		}
	},
	"stackoverflow_tags": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		data := i.MessageComponentData()

		const stackoverflowFormat = `https://stackoverflow.com/questions/tagged/%s`

		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Here is your stackoverflow URL: " + fmt.Sprintf(stackoverflowFormat, strings.Join(data.Values, "+")),
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second) // Doing that so user won't see instant response.
		_, err = s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
			Content: "But wait, there is more! You can also auto populate the select menu. Try executing `/selects auto-populated`.",
			Flags:   discordgo.MessageFlagsEphemeral,
		})
		if err != nil {
			panic(err)
		}
	},
	"channel_select": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "This is it. You've reached your destination. Your choice was <#" + i.MessageComponentData().Values[0] + ">\n" +
					"If you want to know more, check out the links below",
				Components: []discordgo.MessageComponent{
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.Button{
								Emoji: &discordgo.ComponentEmoji{
									Name: "📜",
								},
								Label: "Documentation",
								Style: discordgo.LinkButton,
								URL:   "https://discord.com/developers/docs/interactions/message-components#select-menus",
							},
							discordgo.Button{
								Emoji: &discordgo.ComponentEmoji{
									Name: "🔧",
								},
								Label: "Discord developers",
								Style: discordgo.LinkButton,
								URL:   "https://discord.gg/discord-developers",
							},
							discordgo.Button{
								Emoji: &discordgo.ComponentEmoji{
									Name: "🦫",
								},
								Label: "Discord Gophers",
								Style: discordgo.LinkButton,
								URL:   "https://discord.gg/7RuRrVHyXF",
							},
						},
					},
				},

				Flags: discordgo.MessageFlagsEphemeral,
			},
		})
		if err != nil {
			panic(err)
		}
	},
}
