package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

var models []interface{}

// Models returns every entity registered for schema management. Atlas reads
// this when generating migrations, so registering a model in its own init is
// all that is needed to bring it into the schema.
func Models() []interface{} {
	return models
}

func ConnectDB(user, pw, ip string) {

	log.Println("Connecting to database ...")
	dsn := user + ":" + pw + "@(" + ip + ":3306)/phoenix_bot_db?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	log.Print("DB connected:" + db.Migrator().CurrentDatabase())
}
