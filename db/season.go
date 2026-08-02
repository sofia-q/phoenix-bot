package db

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	log.Println("registering season table")
	models = append(models, &Season{})
}

type Season struct {
	ID      uuid.UUID `gorm:"type:char(36);primary_key"`
	GuildID string    `gorm:"type:varchar(20);not null;uniqueIndex:idx_seasons_guild_number"`
	Number  int       `gorm:"not null;uniqueIndex:idx_seasons_guild_number"`
	// Optional: a season started without one falls back to its number. See
	// DisplayName.
	Name      string    `gorm:"type:varchar(255)"`
	StartedAt time.Time `gorm:"not null"`
	// Nil while the season is running, which is what makes a season current.
	EndedAt *time.Time

	// Declares the foreign key on speedruns. Runs hold a season and a season is
	// never deleted out from under them, so removing one with runs attached
	// fails rather than orphaning or destroying them.
	Speedruns []Speedrun `gorm:"foreignKey:SeasonID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}

func (season *Season) BeforeCreate(_ *gorm.DB) (err error) {
	season.ID = uuid.New()
	return
}

// DisplayName is what a season is called in messages: the name it was given, or
// one built from its number when it was started without a name.
func (season Season) DisplayName() string {
	if season.Name != "" {
		return season.Name
	}
	return fmt.Sprintf("Season %d", season.Number)
}

func (season *Season) Save() (err error) {
	return db.Save(season).Error
}

// FindCurrentForGuild loads the season a guild is currently running. A guild
// between seasons, or one that has never started a season, has none: found
// reports which, rather than that being an error.
func (season *Season) FindCurrentForGuild(guildId string) (found bool, err error) {
	err = db.First(season, "guild_id = ? AND ended_at IS NULL", guildId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
