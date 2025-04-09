package model

import (
	"gorm.io/gorm"
	"log"
)

func init() {
	log.Println("registering speedrun table")
	models = append(models, &Speedrun{})
}

type Speedrun struct {
	gorm.Model
	UserId     string
	timeTaken  string
	weaponType WeaponType `gorm:"type:varchar(255)"`
	proofLink  string
	season     int
}
