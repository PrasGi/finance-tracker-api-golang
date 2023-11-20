package userservice

import (
	"errors"
	database "finance-tracker-api/Database"
	models "finance-tracker-api/Models"
)

func Store(user *models.User) error {

	result := database.DB.Create(&user)

	if result.Error != nil {
		return errors.New(result.Error.Error())
	}

	return nil
}
