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

type UserServiceStruct struct {
	Service *services.UserRepositoryStruct
}

/**************************************************************************************/
func NewUserHandler(serv *services.UserRepositoryStruct) *UserServiceStruct {
	return &UserServiceStruct{Service: serv}
}

/**************************************************************************************/

func (serv *UserServiceStruct) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user *models.UserStruct
	var userDet *models.UserDetailsStruct

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if user.Name == "" || user.Email == "" || user.Password == "" {
		response := map[string]any{
			"message": "name, email and password are required",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hash)

	createdUser, err := serv.Service.RegisterUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userDet = &models.UserDetailsStruct{
		ID:    createdUser.ID,
		Name:  createdUser.Name,
		Email: createdUser.Email,
	}

	response := map[string]any{
		"message": "User created",
		"data":    *userDet,
	}

	json.NewEncoder(w).Encode(response)

}

// GetUser retrieves a user by ID
func (serv *UserServiceStruct) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := serv.Service.Repository.GetUser(uint(id))
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (serv *UserServiceStruct) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := serv.Service.GetAllUsers()
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
