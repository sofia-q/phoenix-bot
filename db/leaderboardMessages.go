package db

import (
	"log"
)

func init() {
	log.Println("registering leaderboard message table")
	models = append(models, &LeaderboardMessage{})
}

type LeaderboardMessage struct {
	GuildID         string `gorm:"primary_key"`
	LeaderboardType string `gorm:"primary_key"`
	ChannelID       string
	MessageID       string
}

func (leaderboardMessage *LeaderboardMessage) Save() (err error) {
	return db.Save(&leaderboardMessage).Error
}

func (leaderboardMessage *LeaderboardMessage) FindMessageIdForLeaderboardTypeAndGuildId(leaderboardType string, guildId string) error {
	return db.Where(&LeaderboardMessage{LeaderboardType: leaderboardType, GuildID: guildId}).Find(leaderboardMessage).Error
}
