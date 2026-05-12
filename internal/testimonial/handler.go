package testimonial

import (
	"catering-api/internal/httpx"
	"encoding/json"
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
	r.Route("/testimonials", func(r chi.Router) {
		r.Get("/{mealId}", h.GetAllTestimonialByMeal)
		r.Post("/", h.CreateTestimonial)
		r.Patch("/{testiId}", h.UpdateTestimonial)
		r.Delete("/{testiId}", h.DeleteTestimonial)
	})
}

func (h *Handler) GetAllTestimonialByMeal(w http.ResponseWriter, r *http.Request) {
	mealId := chi.URLParam(r, "mealId")
	testimonials, err := h.service.FindAllTestimonialByMeal(r.Context(), mealId)
	if err != nil {
		httpx.WriteError(w, http.StatusInternalServerError, "failed get testimonials")
		return
	}
	httpx.WriteSucces(w, http.StatusOK, "Get all testimonials success", testimonials)
}

func (h *Handler) CreateTestimonial(w http.ResponseWriter, r *http.Request) {
	var req CreateTestimonialRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		httpx.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.service.CreateTestimonial(r.Context(), req)

	httpx.WriteSucces(w, http.StatusOK, "Create testimonial success", result)
}

func (h *Handler) UpdateTestimonial(w http.ResponseWriter, r *http.Request) {
	var req UpdateTestimonialRequest
	
	testimonialId := chi.URLParam(r, "testiId")

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		httpx.WriteError(w, http.StatusBadRequest, "Invalid input")
		return
	}

	testimonial, err := h.service.UpdateTestimonial(r.Context(), testimonialId, req)

	if err != nil {
		if err == ErrTestimonialNotFound {
			httpx.WriteError(w, http.StatusNotFound, err.Error())
			return
		}
		httpx.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	httpx.WriteSucces(w, http.StatusOK, "Update testimonial success", testimonial)
}

func (h *Handler) DeleteTestimonial(w http.ResponseWriter, r *http.Request) {
	testimonialId := chi.URLParam(r, "testiId")

	err := h.service.DeleteTestimonial(r.Context(), testimonialId)

	if err != nil {
		if err == ErrTestimonialNotFound {
			httpx.WriteError(w, http.StatusNotFound, err.Error())
			return
		}
		httpx.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
}