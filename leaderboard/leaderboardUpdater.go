package leaderboard

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"phoenixbot/bot/db"
)

func InitializeLeaderboard(guildId string, s *discordgo.Session) {
	updateOverallLeaderboard(guildId, s)
	for i := 0; i < 14; i++ {
		updateWeaponLeaderboard(db.WeaponType(i), guildId, s)
	}
}

func UpdateLeaderboard(guildId string, s *discordgo.Session, weaponType db.WeaponType) {
	updateOverallLeaderboard(guildId, s)
	updateWeaponLeaderboard(weaponType, guildId, s)
}

func updateOverallLeaderboard(guildId string, s *discordgo.Session) {
	var resultList = new(db.SpeedrunList)
	//TODO: season support here
	err := resultList.FindTop10Overall(guildId, 1)
	if err != nil {
		return
	}

	var leaderboardString = "Top 10 hunters: \n"
	for i, entry := range *resultList {
		userMention := "<@" + entry.UserId + ">"
		leaderboardString += fmt.Sprintf("%d: %s: %02d:%02d \n", i+1, userMention, entry.TimeInSeconds/60, entry.TimeInSeconds%60)
	}
	serverMessageData := &db.LeaderboardMessage{}
	_ = serverMessageData.FindMessageIdForLeaderboardTypeAndGuildId("Leaderboard", guildId)
	if serverMessageData != nil {
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

func updateWeaponLeaderboard(weaponType db.WeaponType, guildId string, s *discordgo.Session) {
	//this is showing errors, help

	var resultList = new(db.SpeedrunList)
	//TODO: season support here
	err := resultList.FindTop5ByWeaponType(weaponType, guildId, 1)
	if err != nil {
		return
	}

	var leaderboardString = "Top 5 " + weaponType.String() + " hunters: \n"
	for i, entry := range *resultList {
		userMention := "<@" + entry.UserId + ">"
		leaderboardString += fmt.Sprintf("%d: %s: %02d:%02d \n", i+1, userMention, entry.TimeInSeconds/60, entry.TimeInSeconds%60)
	}
	serverMessageData := &db.LeaderboardMessage{}
	_ = serverMessageData.FindMessageIdForLeaderboardTypeAndGuildId(weaponType.String(), guildId)
	if serverMessageData != nil {
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
