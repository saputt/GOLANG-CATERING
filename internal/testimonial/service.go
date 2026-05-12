package testimonial

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) FindAllTestimonialByMeal(ctx context.Context, mealId string) ([]Testimonial, error) {
	return s.repo.FindAllTestimonialByMeal(ctx, mealId)
}

func (s *Service) FindTestimonialById(ctx context.Context, testiId string) (Testimonial, error) {
	return s.repo.FindTestimonialById(ctx, testiId)
}

func (s *Service) CreateTestimonial(ctx context.Context, req CreateTestimonialRequest) (Testimonial, error) {
	if req.MealId == "" || req.Review == "" || req.UserId == "" {
		return Testimonial{}, errors.New("Invalid input")
	}

	if req.Rating < 0 {
		return Testimonial{}, errors.New("Rating must be large than zero")
	}

	testimonial := Testimonial {
		Id: uuid.NewString(),
		Review: req.Review,
		Rating: req.Rating,
		UserId: req.UserId,
		MealId: req.MealId,
	}	

	return s.repo.CreateTestimonial(ctx, testimonial)
}

func (s *Service) UpdateTestimonial(ctx context.Context, testiId string, req UpdateTestimonialRequest) (Testimonial, error) {
	testimonial, err := s.repo.FindTestimonialById(ctx, testiId)

	if err != nil {
		return Testimonial{}, err
	}

	if req.Rating != nil {
		testimonial.Rating = *req.Rating
	}
	if req.Review != nil {
		testimonial.Review = *req.Review
	}
	
	return s.repo.UpdateTestimonial(ctx, testimonial)
}

func (s *Service) DeleteTestimonial(ctx context.Context, testiId string) error {
	return s.repo.DeleteTestimonial(ctx, testiId)
}