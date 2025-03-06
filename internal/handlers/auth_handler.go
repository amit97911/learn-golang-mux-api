package handlers

import (
	"encoding/json"
	"learn-golang-mux-api/internal/services"
	"net/http"
)

type AuthUserServiceStruct struct {
	Service *services.AuthRepositoryStruct
}

/**************************************************************************************/
func NewAuthUserHandler(service *services.AuthRepositoryStruct) *AuthUserServiceStruct {
	return &AuthUserServiceStruct{Service: service}
}

/**************************************************************************************/

func (serv *AuthUserServiceStruct) Login(w http.ResponseWriter, r *http.Request) {
	var (
		userInput struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		token    *string
		err      error
		response map[string]string
	)

	if err = json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		response = map[string]string{
			"message": "Invalid email or password",
		}
		json.NewEncoder(w).Encode(response)
	}

	token, err = serv.Service.HandleLogin(userInput.Email, userInput.Password)
	if err != nil {
		response = map[string]string{
			"message": "Invalid email or password",
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
	}
	response = map[string]string{
		"message": "Login Successful",
		"token":   *token,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
