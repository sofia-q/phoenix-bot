package db

import (
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	log.Println("registering speedrun table")
	models = append(models, &Speedrun{})
}

type SpeedrunList []Speedrun

type Speedrun struct {
	gorm.Model
	ID            uuid.UUID `gorm:"type:char(36);primary_key"`
	UserId        string    `gorm:"type:varchar(20);not null"`
	TimeInSeconds int       `gorm:"not null"`
	WeaponType    string    `gorm:"type:varchar(255);not null"`
	ProofLink     string    `gorm:"type:varchar(255);not null"`
	IsVerified    bool      `gorm:"not null;default:false"`
	// Every run belongs to a season, and the season is what says which guild it
	// was run in - there is deliberately no GuildID here to disagree with it.
	SeasonID uuid.UUID `gorm:"type:char(36);not null;index"`
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

func (speedrunList *SpeedrunList) FindTop10Overall(seasonId uuid.UUID) (err error) {
	// The season identifies the guild, so there is nothing else to scope by.
	// deleted_at has to be filtered by hand: gorm only excludes soft deleted
	// rows from queries it builds itself, not from raw SQL.
	query := `
	SELECT *
	FROM (
		SELECT *,
			   ROW_NUMBER() OVER (
				   PARTITION BY user_id
				   ORDER BY time_in_seconds, created_at
			   ) AS user_rank
		FROM speedruns
		WHERE season_id = ? AND is_verified = true AND deleted_at IS NULL
	) AS best_user_runs
	WHERE user_rank = 1
	ORDER BY time_in_seconds, created_at
	LIMIT 10;
	`
	return db.Raw(query, seasonId.String()).Scan(speedrunList).Error
}

func (speedrunList *SpeedrunList) FindTop5ByWeaponType(weaponType WeaponType, seasonId uuid.UUID) (err error) {
	// See FindTop10Overall on the season_id and deleted_at conditions.
	query := `
	SELECT *
	FROM (
		SELECT *,
			ROW_NUMBER() OVER (
				PARTITION BY weapon_type
				ORDER BY time_in_seconds, created_at
			) AS weapon_rank
		FROM (
			SELECT *,
				   ROW_NUMBER() OVER (
					   PARTITION BY weapon_type, user_id
					   ORDER BY time_in_seconds, created_at
				   ) AS user_weapon_rank
			FROM speedruns
			WHERE weapon_type = ? AND season_id = ? AND is_verified = true AND deleted_at IS NULL
		) AS unique_user_runs
		WHERE user_weapon_rank = 1
	) AS ranked_runs
	WHERE weapon_rank <= 5
	ORDER BY weapon_rank;
	`
	return db.Raw(query, weaponType.String(), seasonId.String()).Scan(speedrunList).Error
}
