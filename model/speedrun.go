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
	UserId        string
	TimeInSeconds int
	WeaponType    WeaponType `gorm:"type:varchar(255)"`
	ProofLink     string
	Season        int
	IsVerified    bool
}
