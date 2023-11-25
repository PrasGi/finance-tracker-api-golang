package categoryresponse

import (
	models "finance-tracker-api/Models"
	"time"
)

type CategoryResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func CategoryResponseOne(category models.Category) CategoryResponse {
	categoryResponse := CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}

	return categoryResponse
}

func CategoryResponseMany(category []models.Category) []CategoryResponse {

	var categoriesResponse []CategoryResponse

	for _, category := range category {

		categoryResponse := CategoryResponse{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
			CreatedAt:   category.CreatedAt,
			UpdatedAt:   category.UpdatedAt,
		}

		categoriesResponse = append(categoriesResponse, categoryResponse)
	}

	return categoriesResponse
}
