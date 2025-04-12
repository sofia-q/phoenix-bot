package components

import (
	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
	"log"
	"phoenixbot/bot/db"
)

func init() {
	registerComponent(verifySpeedrunComponent)
}

var verifySpeedrunComponent = component{
	name: "verify_button_yes",
	handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {

		id := i.Interaction.Message.Embeds[0].Footer.Text

		foundSpeedrun, err := db.FindSpeedrunById(uuid.MustParse(id))
		//TODO: add verify db functionality here.
		//_ = s.ChannelMessageDelete("1358151701420577009", i.Interaction.Message.ID)
		if err != nil {
			respondErr := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "error! DB entry not found",
				},
			})
			if respondErr != nil {
				log.Println(respondErr.Error())
			}
		}
		respondErr := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "speedrun found, ID: " + foundSpeedrun.ID.String(),
			},
		})
		if err != nil {
			log.Println(respondErr.Error())
		}
	},
}
