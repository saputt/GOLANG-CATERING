package auth

import (
	"catering-api/internal/testimonial"
	"time"
)

type UserRole string

const (
	RoleAdmin UserRole = "admin"
	RoleUser UserRole = "user"
)	

type User struct {
	Id string `json:"id" gorm:"type:uuid;primaryKey"`
	Username string `json:"username" gorm:"not null"`
	Email string `json:"email" gorm:"uniqueIndex;not null"`
	Password string `json:"password" gorm:"not null"`
	Role UserRole `json:"role" gorm:"type:varchar(20);default:user;not null"`
	CreatedAt time.Time `json:"createdAt"`

	Reviews []testimonial.Testimonial `json:"reviews,omitempty" gorm:"foreignKey:MealId"`
}

func (User) TableName() string {
	return "users"
}

type UserResponse struct {
	Id string `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Role UserRole `json:"role"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User UserResponse `json:"user"`
}

type RegisterResponse struct {
	User UserResponse `json:"user"`
}