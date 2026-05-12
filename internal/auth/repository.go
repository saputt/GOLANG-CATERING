package auth

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

var ErrUserNotFound = errors.New("User not found")

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) FindUserByEmail(ctx context.Context, email string) (User, error) {
	var user User

	err := r.db.WithContext(ctx).
		First(&user, "email = ?", email).Error
	
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return User{}, ErrUserNotFound
		}
		return User{}, err
	}

	return user, nil
}

func (r *Repository) CreateUser(ctx context.Context, user User) (User, error) {
	err := r.db.WithContext(ctx).Create(&user).Error

	if err != nil {	
		return User{}, err
	}

	return user, nil
}