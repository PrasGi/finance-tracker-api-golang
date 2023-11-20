package authrequest

import (
	"errors"
	database "finance-tracker-api/Database"
	models "finance-tracker-api/Models"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type BodySignUpRequest struct {
	Username string `json:"username" validate:"required,min=4,max=20"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (body *BodySignUpRequest) ValidateSignupRequest() error {

	// Validate struct fields
	if err := validate.Struct(body); err != nil {
		return err
	}

	// Check uniqueness of Email in the database
	existingUser := models.User{}
	result := database.DB.Where("email = ?", body.Email).First(&existingUser)
	if result.RowsAffected > 0 {
		return errors.New("email already exists")
	}

	return nil
}
