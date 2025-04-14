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
		foundSpeedrun.IsVerified = true
		err = foundSpeedrun.Save()
		if err != nil {
			log.Println(err.Error())
		}
		i.Interaction.Message.Content = ""
		i.Interaction.Message.Embeds[0].Title = "Speedrun Verified!"
		i.Interaction.Message, err = s.ChannelMessageEditComplex(
			&discordgo.MessageEdit{
				Content: &i.Interaction.Message.Content,
				Embeds:  &i.Interaction.Message.Embeds,
				AllowedMentions: &discordgo.MessageAllowedMentions{
					Parse: []discordgo.AllowedMentionType{},
					Users: []string{},
				},
				Components: &[]discordgo.MessageComponent{},
				Channel:    "1358151701420577009",
				ID:         i.Interaction.Message.ID,
			})
		if err != nil {
			log.Println(err.Error())
		}
		respondErr := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Success! Speedrun verified!",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		if err != nil {
			log.Println(respondErr.Error())
		}
	},
}
