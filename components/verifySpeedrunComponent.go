package components

import (
	"github.com/bwmarrin/discordgo"
)

func init() {
	registerComponent(verifySpeedrunComponent)
}

var verifySpeedrunComponent = component{
	name: "verify_button_yes",
	handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {

		//TODO: add verify db functionality here.

		_ = s.ChannelMessageDelete("1358151701420577009", i.Interaction.Message.ID)

		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "speedrun verified",
			},
		})
		if err != nil {
			panic(err)
		}
	},
}
