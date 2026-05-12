package testimonial

import "time"

type Testimonial struct {
	Id string `json:"id" gorm:"type:uuid;primaryKey"`
	Review string `json:"review" gorm:"no null"`
	Rating int `json:"rating" gorm:"not null"`
	UserId string `json:"userId" gorm:"not null"`
	MealId string `json:"mealId" gorm:"not null"`

	CreatedAt time.Time `json:"createdAt"`
}

type CreateTestimonialRequest struct {
	Review string `json:"review"`
	Rating int `json:"rating"`
	UserId string `json:"userId"`
	MealId string `json:"mealId"`
}

type UpdateTestimonialRequest struct {
	Review *string `json:"review"`
	Rating *int `json:"rating"`
}