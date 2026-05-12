package httpx

import (
	"encoding/json"
	"net/http"
)

type JSONResponse struct {
	Message string `json:"message"`
	Data interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func WriteJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(payload)
}

func WriteSucces(w http.ResponseWriter, status int, message string, data interface{}){
	WriteJSON(w, status, JSONResponse{
		Message: message,
		Data: data,
	})
}

func WriteError(w http.ResponseWriter, status int, message string){
	WriteJSON(w, status, ErrorResponse{
		Message: message,
	})
}