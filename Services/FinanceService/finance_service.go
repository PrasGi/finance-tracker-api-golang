package financeservice

import (
	"errors"
	database "finance-tracker-api/Database"
	models "finance-tracker-api/Models"
)

func Index(finance *[]models.Finance) error {

	err := database.DB.Preload("User").Preload("Category").Find(&finance).Error

	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func Store(finance *models.Finance) error {

	result := database.DB.Create(&finance)

	if result.Error != nil {
		return errors.New(result.Error.Error())
	}

	return nil
}

func Show(finance *models.Finance, id int64) error {

	result := database.DB.Preload("User").Preload("Category").First(&finance, "id = ?", id)

	if result.Error != nil {
		return errors.New(result.Error.Error())
	}

	return nil
}

func Update(finance *models.Finance, id int64) error {

	result := database.DB.Model(models.Finance{}).Where("id = ?", id).Updates(&finance)

	if result.Error != nil {
		return errors.New(result.Error.Error())
	}

	return nil
}

func Destroy(finance *models.Finance, id int64) error {

	result := database.DB.Delete(&finance, "id = ?", id)

	if result.Error != nil {
		return errors.New(result.Error.Error())
	}

	return nil
}
