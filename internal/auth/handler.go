package auth

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
	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", h.Login)
		r.Post("/register", h.Register)
	})
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		httpx.WriteError(w, http.StatusBadRequest, "Invalid body")
		return
	}

	result, err := h.service.Register(r.Context(), req)
	if err != nil {
		switch err {
		case ErrEmailAlreadyExist:
			httpx.WriteError(w, http.StatusConflict, err.Error())
		case ErrInvalidInput:
			httpx.WriteError(w, http.StatusBadRequest, err.Error())
		default:
			httpx.WriteError(w, http.StatusBadRequest, err.Error())
		}
		return
	}

	httpx.WriteSucces(w, http.StatusOK, "Register success", result)
}
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		httpx.WriteError(w, http.StatusBadRequest, "Invalid body")
		return
	}

	result, err := h.service.Login(r.Context(), req)
	if err != nil {
		switch err {
		case ErrInvalidCredential:
			httpx.WriteError(w, http.StatusForbidden, err.Error())
		case ErrInvalidInput:
			httpx.WriteError(w, http.StatusBadRequest, err.Error())
		default:
			httpx.WriteError(w, http.StatusBadRequest, err.Error())
		}
		return
	}

	httpx.WriteSucces(w, http.StatusOK, "Login success", result)
}