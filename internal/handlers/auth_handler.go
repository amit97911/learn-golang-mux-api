package handlers

import (
	"encoding/json"
	"learn-golang-mux-api/internal/services"
	"net/http"
)

type AuthUserHandlerStruct struct {
	Service *services.AuthServiceStruct
}

func AuthUserHandler(service *services.AuthServiceStruct) *AuthUserHandlerStruct {
	return &AuthUserHandlerStruct{Service: service}
}

func (h *AuthUserHandlerStruct) Login(w http.ResponseWriter, r *http.Request) {
	var userInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	_, err := h.Service.HandleLogin(userInput.Email, userInput.Password)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Set a session or cookie here if needed
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Login successful")
}
