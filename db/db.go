package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

var models []interface{}

func ConnectDB(user, pw, ip string) {

	log.Println("Connecting to database ...")
	dsn := user + ":" + pw + "@(" + ip + ":3306)/phoenix_bot_db?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
		return
	}
	log.Print("DB connected:" + db.Migrator().CurrentDatabase())
	log.Print("Migrating database ...")
	if err := db.AutoMigrate(models...); err != nil {
		panic("failed to migrate database: " + err.Error())
	}
}
