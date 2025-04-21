package commands

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"phoenixbot/bot/db"
)

func init() {
	registerCommand(setupLeaderboardCommand)
}

var setupLeaderboardCommand = command{
	name: setupLeaderboardCommandName,
	//this is the metadata for the commands
	commandMetadata: &discordgo.ApplicationCommand{
		Name: setupLeaderboardCommandName,
		// All commands and options must have a description
		// Commands/options without description will fail the registration
		// of the command.
		Description: "sets up the leaderboard to print in this channel",
	},
	commandHandler: setUpLeaderboard,
}

// this is the command name
var setupLeaderboardCommandName = "setup-leaderboard"

// this is what the command actually does
func setUpLeaderboard(s *discordgo.Session, i *discordgo.InteractionCreate) {

	_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   discordgo.MessageFlagsEphemeral,
			Content: "Leaderboard setup complete!",
		},
	})

	//future todo: delete old entries if existing

	msg, err := s.ChannelMessageSendComplex(i.ChannelID, &discordgo.MessageSend{
		Content: "placeholder for overall leaderboard",
	})
	if err != nil {
		log.Println("Something went wrong sending placeholder message for leaderboard")
		return
	}

	newMessage := db.LeaderboardMessage{
		GuildID:         i.GuildID,
		LeaderboardType: "Leaderboard",
		MessageID:       msg.ID,
		ChannelID:       i.ChannelID,
	}
	_ = newMessage.Save()

	for j := 0; j < 14; j++ {

		msg, err := s.ChannelMessageSendComplex(i.ChannelID, &discordgo.MessageSend{
			Content: "placeholder for weapon type: " + db.WeaponType(j).String(),
		})
		if err != nil {
			log.Println("Something went wrong sending placeholder message for leaderboard")
			return
		}

		newMessage := db.LeaderboardMessage{
			GuildID:         i.GuildID,
			LeaderboardType: db.WeaponType(j).String(),
			MessageID:       msg.ID,
			ChannelID:       i.ChannelID,
		}
		_ = newMessage.Save()
	}

	//todo: update leaderboard
}
