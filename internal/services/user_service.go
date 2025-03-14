package services

import (
	"fmt"
	"learn-golang-mux-api/internal/models"
	"learn-golang-mux-api/internal/repositories"
)

type UserRepositoryStruct struct {
	Repository *repositories.DatabaseConnection
}

/**************************************************************************************/
func NewUserService(repo *repositories.DatabaseConnection) *UserRepositoryStruct {
	return &UserRepositoryStruct{Repository: repo}
}

/**************************************************************************************/

func (repo *UserRepositoryStruct) RegisterUser(user *models.UserStruct) (*models.UserStruct, error) {
	err := repo.Repository.CreateUser(user)
	return user, err
}

func (repo *UserRepositoryStruct) GetAllUsers() ([]*models.UserDetailsStruct, error) {
	users, err := repo.Repository.GetAllUsers()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve users: %w", err)
	}
	return users, nil
}
