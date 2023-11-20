package authrequest

import "github.com/go-playground/validator/v10"

func init() {
	validate = validator.New()
}

type BodySignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (body *BodySignInRequest) ValidateSignInRequest() error {

	// Validate struct fields
	if err := validate.Struct(body); err != nil {
		return err
	}

	return nil
}
