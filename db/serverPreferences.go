package db

import (
	"errors"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func init() {
	log.Println("registering server preferences table")
	models = append(models, &ServerPreferences{})
}

type ServerPreferences struct {
	GuildID              string `gorm:"primary_key"`
	Season               int
	IsSeasonActive       bool
	LeaderboardChannelID string
}

func (serverPreferences *ServerPreferences) Save() (err error) {
	return db.Save(serverPreferences).Error
}

// FindOrCreateForGuild loads a guild's preferences, creating the row the first
// time it is asked for. There is therefore no missing case for callers to
// handle, and no need to remember to set GuildID before saving.
//
// Note that this writes: a guild is given a row as soon as anything asks about
// it, not when it first configures something.
func (serverPreferences *ServerPreferences) FindOrCreateForGuild(guildId string) error {
	err := db.First(serverPreferences, "guild_id = ?", guildId).Error
	if err == nil {
		return nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// No row yet. Another handler may be creating one for this guild at the
	// same moment, so let a conflicting insert pass and read back whichever
	// row won rather than failing on the duplicate key.
	*serverPreferences = ServerPreferences{GuildID: guildId}
	result := db.Clauses(clause.OnConflict{DoNothing: true}).Create(serverPreferences)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 1 {
		return nil
	}

	return db.First(serverPreferences, "guild_id = ?", guildId).Error
}
