package handlers

import (
	"encoding/json"
	"learn-golang-mux-api/internal/models"
	"learn-golang-mux-api/internal/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type UserHandlerStruct struct {
	Service *services.UserServiceStruct
}

func UserHandler(service *services.UserServiceStruct) *UserHandlerStruct {
	return &UserHandlerStruct{Service: service}
}

// CreateUser handles user creation
func (h *UserHandlerStruct) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.UserWithPasswordStruct

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hash)

	createdUser, err := h.Service.RegisterUser(user.Name, user.Email, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]any{
		"message": "User created",
		"data": map[string]any{
			"name":  createdUser.Name,
			"email": createdUser.Email,
		},
	}

	json.NewEncoder(w).Encode(response)

}

// GetUser retrieves a user by ID
func (h *UserHandlerStruct) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.Service.Repo.GetUser(uint(id))
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *UserHandlerStruct) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.Service.GetAllUsers()
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
