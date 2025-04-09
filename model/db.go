package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"phoenixbot/bot/env"
)

var Db *gorm.DB

var models []interface{}

func ConnectDB() {

	log.Println("Connecting to database ...")
	dsn := env.DatabaseUser + ":" + env.DatabasePw + "@(" + env.DatabaseIp + ":3306)/phoenix_bot_db?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
		return
	}
	log.Print("DB connected:" + Db.Migrator().CurrentDatabase())
	log.Print("Migrating database ...")
	_ = Db.AutoMigrate(models...)
}
