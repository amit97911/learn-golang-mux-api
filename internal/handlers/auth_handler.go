package handlers

import (
	"encoding/json"
	"learn-golang-mux-api/internal/services"
	"net/http"
)

type AuthUserHandlerStruct struct {
	Service *services.UserServiceStruct
}

func AuthUserHandler(service *services.UserServiceStruct) *AuthUserHandlerStruct {
	return &AuthUserHandlerStruct{Service: service}
}

func (h *AuthUserHandlerStruct) Login(w http.ResponseWriter, r *http.Request) {
	response := map[string]any{
		"message": "Flow Pending",
	}
	json.NewEncoder(w).Encode(response)
}
