package db

import (
	"log"
)

func init() {
	log.Println("registering leaderboard message table")
	models = append(models, &LeaderboardMessage{})
}

type LeaderboardMessage struct {
	GuildID         string `gorm:"type:varchar(20);primary_key"`
	LeaderboardType string `gorm:"type:varchar(64);primary_key"`
	ChannelID       string `gorm:"type:varchar(20);not null"`
	MessageID       string `gorm:"type:varchar(20);not null"`
}

func (leaderboardMessage *LeaderboardMessage) Save() (err error) {
	return db.Save(&leaderboardMessage).Error
}

func (leaderboardMessage *LeaderboardMessage) FindMessageIdForLeaderboardTypeAndGuildId(leaderboardType string, guildId string) error {
	return db.Where(&LeaderboardMessage{LeaderboardType: leaderboardType, GuildID: guildId}).Find(leaderboardMessage).Error
}
