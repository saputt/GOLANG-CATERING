package meal

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return  &Service{
		repo: repo,
	}
}

func (s *Service) GetAllMeals(ctx context.Context) ([]Meal, error){
	return s.repo.FindAll(ctx)
}

func (s *Service) GetMeal(ctx context.Context, id string) (Meal, error) {
	return s.repo.FindById(ctx, id)
}

func (s *Service) Create(ctx context.Context, req CreateMealRequest) (Meal, error) {
	if(req.Name == ""){
		return Meal{}, errors.New("Name is required")
	}
	if(req.Price < 0){
		return Meal{}, errors.New("Price must be large than zero")
	}

	meal := Meal{
		Id: uuid.NewString(),
		Name: req.Name,
		Price: req.Price,
		AverageRating: req.AverageRating,
		Description: req.Description,
		ImageUrl: req.ImageUrl,
	}

	return s.repo.Create(ctx, meal)
}

func (s *Service) Update(ctx context.Context, id string, req UpdateMealRequest) (Meal, error) {
	meal, err := s.repo.FindById(ctx, id)
	if err != nil {
		return Meal{}, err
	}
	if req.Name != nil {
		meal.Name = *req.Name
	}
	if req.Price != nil {
		meal.Price = *req.Price
	}
	if req.AverageRating != nil {
		meal.AverageRating = *req.AverageRating
	}
	if req.Description != nil {
		meal.Description = *req.Description
	}
	if req.ImageUrl != nil {
		meal.ImageUrl = *req.ImageUrl
	}

	return s.repo.Update(ctx, meal)
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}