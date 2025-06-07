package components

import (
	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
	"log"
	"phoenixbot/bot/db"
	"phoenixbot/bot/leaderboard"
)

func init() {
	registerComponent(verifySpeedrunComponent)
}

var verifySpeedrunComponent = component{
	name:    "verify_button_yes",
	handler: verifySpeedrun,
}

func verifySpeedrun(s *discordgo.Session, i *discordgo.InteractionCreate) {

	i.Interaction.Message.Content = ""
	i.Interaction.Message.Embeds[0].Title = "Speedrun Verified!"
	var editErr error
	i.Interaction.Message, editErr = s.ChannelMessageEditComplex(
		&discordgo.MessageEdit{
			Content: &i.Interaction.Message.Content,
			Embeds:  &i.Interaction.Message.Embeds,
			AllowedMentions: &discordgo.MessageAllowedMentions{
				Parse: []discordgo.AllowedMentionType{},
				Users: []string{},
			},
			Components: &[]discordgo.MessageComponent{},
			Channel:    i.Interaction.ChannelID,
			ID:         i.Interaction.Message.ID,
		})
	if editErr != nil {
		log.Println(editErr.Error())
	}

	respondErr := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Success! Speedrun verified!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	if respondErr != nil {
		log.Println(respondErr.Error())
	}

	id := i.Interaction.Message.Embeds[0].Footer.Text
	var foundSpeedrun db.Speedrun
	findErr := foundSpeedrun.FindSpeedrunById(uuid.MustParse(id))
	if findErr != nil {
		log.Println(findErr.Error())
		return
	}
	foundSpeedrun.IsVerified = true
	saveErr := foundSpeedrun.Save()

	if saveErr != nil {
		log.Println(saveErr.Error())
		return
	}
	var weaponType db.WeaponType
	foundType, parseErr := weaponType.ParseStringToWeaponType(foundSpeedrun.WeaponType)
	if parseErr != nil {
		log.Println(parseErr.Error())
	}
	leaderboard.UpdateLeaderboard(i.GuildID, s, foundType)
}
