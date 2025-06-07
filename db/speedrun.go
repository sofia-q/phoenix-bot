package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

func init() {
	log.Println("registering speedrun table")
	models = append(models, &Speedrun{})
}

type SpeedrunList []Speedrun

type Speedrun struct {
	gorm.Model
	ID            uuid.UUID `gorm:"primary_key"`
	UserId        string    `gorm:"type:varchar(255)"`
	TimeInSeconds int
	WeaponType    string `gorm:"type:varchar(255)"`
	ProofLink     string `gorm:"type:varchar(255)"`
	Season        int
	IsVerified    bool
	GuildID       string
}

func (speedrun *Speedrun) BeforeCreate(_ *gorm.DB) (err error) {
	speedrun.ID = uuid.New()
	return
}

func (speedrun *Speedrun) FindSpeedrunById(uuid uuid.UUID) (err error) {
	return db.First(&speedrun, "id = ?", uuid.String()).Error
}

func (speedrun *Speedrun) Save() (err error) {
	return db.Save(&speedrun).Error
}

func (speedrunList *SpeedrunList) FindTop10Overall(guildId string, season int) (err error) {
	query := `
	SELECT *
	FROM (
		SELECT *,
			   ROW_NUMBER() OVER (
				   PARTITION BY user_id
				   ORDER BY time_in_seconds
			   ) AS user_rank
		FROM speedruns
		WHERE guild_id = ? AND season = ? AND is_verified = true
	) AS best_user_runs
	WHERE user_rank = 1
	ORDER BY time_in_seconds
	LIMIT 10;
	`
	return db.Raw(query, guildId, season).Scan(speedrunList).Error
}

func (speedrunList *SpeedrunList) FindTop5ByWeaponType(weaponType WeaponType, guildId string, season int) (err error) {

	query := `
	SELECT *
	FROM (
		SELECT *,
			ROW_NUMBER() OVER (
				PARTITION BY weapon_type
				ORDER BY time_in_seconds
			) AS weapon_rank
		FROM (
			SELECT *,
				   ROW_NUMBER() OVER (
					   PARTITION BY weapon_type, user_id
					   ORDER BY time_in_seconds
				   ) AS user_weapon_rank
			FROM speedruns
			WHERE weapon_type = ? AND guild_id = ? AND season = ? AND is_verified = true
		) AS unique_user_runs
		WHERE user_weapon_rank = 1
	) AS ranked_runs
	WHERE weapon_rank <= 5
	ORDER BY weapon_rank;
	`
	return db.Raw(query, weaponType.String(), guildId, season).Scan(speedrunList).Error
}
