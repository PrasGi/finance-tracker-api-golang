package financerequest

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type BodyFinanceRequest struct {
	Title      string  `json:"title" validate:"required"`
	Amount     float64 `json:"amount" validate:"required"`
	CategoryId uint    `json:"category_id" validate:"required"`
	UserId     uint    `json:"user_id" validate:"required"`
}

func (body *BodyFinanceRequest) ValidateCreateRequest() error {

	// Validate struct fields
	if err := validate.Struct(body); err != nil {
		return err
	}

	return nil
}
