package testimonial

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

var ErrTestimonialNotFound = errors.New("Testimonial not found")

func (r *Repository) FindAllTestimonialByMeal(ctx context.Context, mealId string) ([]Testimonial, error) {
	var testimonials []Testimonial

	err := r.db.WithContext(ctx).Order("created_at DESC").Find(&testimonials, "meal_id = ?", mealId).Error

	if err != nil {
		return []Testimonial{}, err
	}

	return testimonials, nil
}

func (r *Repository) FindTestimonialById(ctx context.Context, testiId string) (Testimonial, error) {
	var testimonial Testimonial
	err := r.db.WithContext(ctx).Find(&testimonial, "id = ?", testiId).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return Testimonial{}, ErrTestimonialNotFound
		}
		return Testimonial{}, err
	}

	return testimonial, nil
}

func (r *Repository) CreateTestimonial(ctx context.Context, testimonial Testimonial) (Testimonial, error) {
	err := r.db.WithContext(ctx).Create(&testimonial).Error

	if err != nil {
		return Testimonial{}, err
	}

	return testimonial, nil
}

func (r *Repository) UpdateTestimonial(ctx context.Context, testimonial Testimonial) (Testimonial, error) {
	err := r.db.WithContext(ctx).Save(&testimonial).Error

	if err != nil {
		return Testimonial{}, err
	}

	return testimonial, nil
}

func (r *Repository) DeleteTestimonial(ctx context.Context, testiId string) error {
	result := r.db.WithContext(ctx).Delete(&Testimonial{}, "id = ?", testiId)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrTestimonialNotFound
	}

	return nil
} 