package main

import (
	database "finance-tracker-api/Database"
	helpers "finance-tracker-api/Helpers"
	models "finance-tracker-api/Models"
)

func init() {
	helpers.LoadEnv()
	database.InitDB()
}

func main() {
	database.DB.Migrator().DropTable(&models.User{})
	database.DB.AutoMigrate(&models.User{})

	database.DB.AutoMigrate(&models.PersonalAccessToken{})
	database.DB.AutoMigrate(&models.Category{})
}
