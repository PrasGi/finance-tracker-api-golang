package categoryrequest

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type BodyCategoryRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func (body *BodyCategoryRequest) ValidateCreateRequest() error {

	// Validate struct fields
	if err := validate.Struct(body); err != nil {
		return err
	}

	return nil
}
