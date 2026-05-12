package meal

import (
	"catering-api/internal/auth"
	"catering-api/internal/httpx"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Route("/meals", func(r chi.Router) {
		r.Get("/", h.GetAllMeals)
		r.Get("/{mealId}", h.GetMeal)
		r.With(auth.AdminOnly).Post("/", h.CreateMeal)
		r.With(auth.AdminOnly).Patch("/{mealId}", h.UpdateMeal)
		r.With(auth.AdminOnly).Delete("/{mealId}", h.DeleteMeal)
	})
}

func (h *Handler) GetAllMeals(w http.ResponseWriter, r *http.Request){
	meals, err := h.service.GetAllMeals(r.Context())
	fmt.Println(err)
	if err != nil {
		httpx.WriteError(w, http.StatusInternalServerError, "failed get meals")
		return 
	}

	httpx.WriteSucces(w, http.StatusOK, "Get all meals success", meals)
}

func (h *Handler) GetMeal(w http.ResponseWriter, r *http.Request){
	mealId := chi.URLParam(r, "mealId")
	meal, err := h.service.GetMeal(r.Context(), mealId)
	if err != nil {
		httpx.WriteError(w, http.StatusNotFound, "Meal not found")
		return
	}
	httpx.WriteSucces(w, http.StatusOK, "Get meal success", meal)
}

func (h *Handler) CreateMeal(w http.ResponseWriter, r *http.Request) {
	var req CreateMealRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if(err != nil){
		httpx.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return 
	}

	meal, err := h.service.Create(r.Context(), req)
	if(err != nil){
		httpx.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	httpx.WriteSucces(w, http.StatusOK, "Create meal success", meal)
}

func (h *Handler) UpdateMeal(w http.ResponseWriter, r *http.Request) {
	mealId := chi.URLParam(r, "mealId")
	var req UpdateMealRequest
	
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		httpx.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	meal, err := h.service.Update(r.Context(), mealId, req)
	if err != nil {
		if (err == ErrMealNotFound){
			httpx.WriteError(w, http.StatusNotFound, err.Error())
			return
		}
		httpx.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	httpx.WriteSucces(w, http.StatusOK, "Update meal success", meal)
}

func (h *Handler) DeleteMeal(w http.ResponseWriter, r *http.Request) {
	mealId := chi.URLParam(r, "mealId")
	err := h.service.Delete(r.Context(), mealId)
	if(err != nil) {
		if err == ErrMealNotFound {
			httpx.WriteError(w, http.StatusNotFound, err.Error())
			return
		}
		httpx.WriteError(w, http.StatusInternalServerError, "failed to delete meal")
		return
	}
	httpx.WriteSucces(w, http.StatusOK, "Success delete meal", nil)
}