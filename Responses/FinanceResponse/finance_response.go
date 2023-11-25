package financeresponse

import (
	models "finance-tracker-api/Models"
	"time"
)

type FinanceResponse struct {
	ID           uint      `json:"id"`
	Title        string    `json:"title"`
	Amount       float64   `json:"amount"`
	UserId       uint      `json:"user_id"`
	UserName     string    `json:"user_name"`
	CategoryId   uint      `json:"category_id"`
	CategoryName string    `json:"category_name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func FinanceResponseOne(finance models.Finance) FinanceResponse {
	financeResponse := FinanceResponse{
		ID:           finance.ID,
		Title:        finance.Title,
		Amount:       finance.Amount,
		UserId:       finance.UserId,
		UserName:     finance.User.Username,
		CategoryId:   finance.CategoryId,
		CategoryName: finance.Category.Name,
		CreatedAt:    finance.CreatedAt,
		UpdatedAt:    finance.UpdatedAt,
	}

	return financeResponse
}

func FinanceResponseMany(finance []models.Finance) []FinanceResponse {

	var financesResponse []FinanceResponse

	for _, finance := range finance {

		financeResponse := FinanceResponse{
			ID:           finance.ID,
			Title:        finance.Title,
			Amount:       finance.Amount,
			UserId:       finance.UserId,
			UserName:     finance.User.Username,
			CategoryId:   finance.CategoryId,
			CategoryName: finance.Category.Name,
			CreatedAt:    finance.CreatedAt,
			UpdatedAt:    finance.UpdatedAt,
		}

		financesResponse = append(financesResponse, financeResponse)
	}

	return financesResponse
}
