package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"phoenixbot/bot/env"
)

var Db *gorm.DB

func ConnectDB() {

	databaseIp := env.LoadVar("DATABASE_IP")
	if databaseIp == "" {
		fmt.Println("DATABASE_IP environment variable not found")
		return
	}
	databaseUser := env.LoadVar("DATABASE_USER")
	if databaseUser == "" {
		fmt.Println("DATABASE_IP environment variable not found")
		return
	}
	databasePw := env.LoadVar("DATABASE_PW")
	if databasePw == "" {
		fmt.Println("DATABASE_IP environment variable not found")
		return
	}
	log.Println("Connecting to database ...")
	dsn := databaseUser + ":" + databasePw + "@(" + databaseIp + ":3306)/?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
		return
	}
	_ = Db.Exec("CREATE DATABASE bot_database;")
	log.Print("DB connected:" + Db.Migrator().CurrentDatabase())
}
