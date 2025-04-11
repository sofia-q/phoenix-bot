package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

func init() {
	log.Println("registering speedrun table")
	models = append(models, &Speedrun{})
}

type Speedrun struct {
	gorm.Model
	ID            uuid.UUID `gorm:"primary_key"`
	UserId        string    `gorm:"type:varchar(255)"`
	TimeInSeconds int
	WeaponType    string `gorm:"type:varchar(255)"`
	ProofLink     string `gorm:"type:varchar(255)"`
	Season        int
	IsVerified    bool
}

func (speedrun *Speedrun) BeforeCreate(_ *gorm.DB) (err error) {
	speedrun.ID = uuid.New()
	return
}
