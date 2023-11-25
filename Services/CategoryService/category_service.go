package categoryservice

import (
	"errors"
	database "finance-tracker-api/Database"
	models "finance-tracker-api/Models"
)

func Index(categories *[]models.Category) error {

	err := database.DB.Find(&categories).Error

	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func Store(category *models.Category) error {

	result := database.DB.Create(&category)

	if result.Error != nil {
		return errors.New(result.Error.Error())
	}

	return nil
}

func Show(category *models.Category, id int64) error {

	result := database.DB.First(&category, "id = ?", id)

	if result.Error != nil {
		return errors.New(result.Error.Error())
	}

	return nil
}

func Update(category *models.Category, id int64) error {

	result := database.DB.Model(models.Category{}).Where("id = ?", id).Updates(&category)

	if result.Error != nil {
		return errors.New(result.Error.Error())
	}

	return nil
}

func Destroy(category *models.Category, id int64) error {

	result := database.DB.Delete(&category, "id = ?", id)

	if result.Error != nil {
		return errors.New(result.Error.Error())
	}

	return nil
}
