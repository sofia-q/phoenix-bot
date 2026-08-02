package leaderboard

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
	"log"
	"phoenixbot/bot/db"
)

func InitializeLeaderboard(guildId string, s *discordgo.Session) {
	seasonId, ok := currentSeasonId(guildId)
	if !ok {
		return
	}

	updateOverallLeaderboard(guildId, seasonId, s)
	for i := 0; i < 14; i++ {
		updateWeaponLeaderboard(db.WeaponType(i), guildId, seasonId, s)
	}
}

func UpdateLeaderboard(guildId string, s *discordgo.Session, weaponType db.WeaponType) {
	seasonId, ok := currentSeasonId(guildId)
	if !ok {
		return
	}

	updateOverallLeaderboard(guildId, seasonId, s)
	updateWeaponLeaderboard(weaponType, guildId, seasonId, s)
}

// currentSeasonId reports which season the leaderboards should be showing.
// Between seasons there is nothing to render, so callers stop rather than
// overwrite the standings the last season ended on.
func currentSeasonId(guildId string) (seasonId uuid.UUID, ok bool) {
	season := db.Season{}
	found, err := season.FindCurrentForGuild(guildId)
	if err != nil {
		log.Println("Something went wrong loading the current season: " + err.Error())
		return uuid.UUID{}, false
	}
	return season.ID, found
}

func updateOverallLeaderboard(guildId string, seasonId uuid.UUID, s *discordgo.Session) {
	var resultList = new(db.SpeedrunList)
	err := resultList.FindTop10Overall(seasonId)
	if err != nil {
		log.Println("Something went wrong loading the overall leaderboard: " + err.Error())
		return
	}

	var leaderboardString = "Top 10 hunters: \n"
	for i, entry := range *resultList {
		userMention := "<@" + entry.UserId + ">"
		leaderboardString += fmt.Sprintf("%d: %s: %02d:%02d \n", i+1, userMention, entry.TimeInSeconds/60, entry.TimeInSeconds%60)
	}
	serverMessageData := &db.LeaderboardMessage{}
	_ = serverMessageData.FindMessageIdForLeaderboardTypeAndGuildId("Leaderboard", guildId)
	// A guild that never ran the leaderboard setup has no message to edit.
	if serverMessageData.MessageID != "" {
		_, editErr := s.ChannelMessageEditComplex(&discordgo.MessageEdit{
			ID:      serverMessageData.MessageID,
			Channel: serverMessageData.ChannelID,
			Content: &leaderboardString,
			AllowedMentions: &discordgo.MessageAllowedMentions{
				Parse: []discordgo.AllowedMentionType{},
				Users: []string{},
			},
		})
		if editErr != nil {
			log.Println(editErr.Error())
		}
	}

}

func updateWeaponLeaderboard(weaponType db.WeaponType, guildId string, seasonId uuid.UUID, s *discordgo.Session) {
	//this is showing errors, help

	var resultList = new(db.SpeedrunList)
	err := resultList.FindTop5ByWeaponType(weaponType, seasonId)
	if err != nil {
		log.Println("Something went wrong loading the " + weaponType.String() + " leaderboard: " + err.Error())
		return
	}

	var leaderboardString = "Top 5 " + weaponType.String() + " hunters: \n"
	for i, entry := range *resultList {
		userMention := "<@" + entry.UserId + ">"
		leaderboardString += fmt.Sprintf("%d: %s: %02d:%02d \n", i+1, userMention, entry.TimeInSeconds/60, entry.TimeInSeconds%60)
	}
	serverMessageData := &db.LeaderboardMessage{}
	_ = serverMessageData.FindMessageIdForLeaderboardTypeAndGuildId(weaponType.String(), guildId)
	// A guild that never ran the leaderboard setup has no message to edit.
	if serverMessageData.MessageID != "" {
		_, editErr := s.ChannelMessageEditComplex(&discordgo.MessageEdit{
			ID:      serverMessageData.MessageID,
			Channel: serverMessageData.ChannelID,
			Content: &leaderboardString,
			AllowedMentions: &discordgo.MessageAllowedMentions{
				Parse: []discordgo.AllowedMentionType{},
				Users: []string{},
			},
		})
		if editErr != nil {
			log.Println(editErr.Error())
		}
	}
}
