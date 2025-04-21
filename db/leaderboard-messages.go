package db

import "log"

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
	return Db.Save(&leaderboardMessage).Error
}
