package database

import (
	helpers "finance-tracker-api/Helpers"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	destination := os.Getenv("DB_URL")

	DB, err = gorm.Open(mysql.Open(destination), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	helpers.PanicIfErrSystem(err)
}
