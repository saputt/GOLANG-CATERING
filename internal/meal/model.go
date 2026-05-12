package meal

import (
	"catering-api/internal/testimonial"
	"time"
)

type Meal struct {
	Id string `json:"id" gorm:"type:uuid;primaryKey"`	
	Name string `json:"name" gorm:"not null"`
	Price int `json:"price" gorm:"not null"`
	AverageRating float64 `json:"averageRating"`
	Description string `json:"description"`
	ImageUrl string `json:"imageUrl"`
	CreatedAt time.Time `json:"createdAt"`

	Testimonies []testimonial.Testimonial `json:"testimonies,omitempty" gorm:"foreignKey:MealId"`
}

type CreateMealRequest struct {
	Name string `json:"name"`
	Price int `json:"price"`
	AverageRating float64 `json:"averageRating"`
	Description string `json:"description"`
	ImageUrl string `json:"imageUrl"`
}

type UpdateMealRequest struct {
	Name *string `json:"name"`
	Price *int `json:"price"`
	AverageRating *float64 `json:"averageRating"`
	Description *string `json:"description"`
	ImageUrl *string `json:"imageUrl"`
}

