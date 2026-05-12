package meal

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

var (
	ErrMealNotFound = errors.New("Meal not found")
)

func (r *Repository) FindAll(ctx context.Context) ([]Meal, error) {
	var meals []Meal
	
	err := r.db.WithContext(ctx).Order("created_at DESC").Find(&meals).Error

	if err != nil {
		return nil, err
	}

	return meals, nil
}	

func (r *Repository) FindById(ctx context.Context, id string) (Meal, error) {
	var meal Meal

	err := r.db.WithContext(ctx).Find(&meal, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Meal{}, ErrMealNotFound
		}

		return Meal{}, err
	}

	return meal, nil
}

func (r *Repository) Create(ctx context.Context, meal Meal) (Meal, error) {
	err := r.db.WithContext(ctx).Create(&meal).Error

	if err != nil {
		return Meal{}, nil
	}

	return meal, nil
}

func (r *Repository) Update(ctx context.Context, meal Meal) (Meal, error) {
	err := r.db.WithContext(ctx).Save(&meal).Error

	if err != nil {
		return Meal{}, nil
	}

	return meal, nil
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Delete(&Meal{}, "id = ?", id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrMealNotFound
	}

	return nil
}