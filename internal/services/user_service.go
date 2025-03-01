package services

import (
	"errors"
	"fmt"
	"learn-golang-mux-api/internal/models"
	"learn-golang-mux-api/internal/repositories"
)

type UserServiceStruct struct {
	Repo *repositories.DatabaseConnection
}

func UserService(repo *repositories.DatabaseConnection) *UserServiceStruct {
	return &UserServiceStruct{Repo: repo}
}

func (service *UserServiceStruct) RegisterUser(name, email, password string) (*models.UserWithPasswordStruct, error) {
	if name == "" || email == "" || password == "" {
		return nil, errors.New("name and email are required")
	}

	user := &models.UserWithPasswordStruct{Name: name, Email: email, Password: password}
	err := service.Repo.CreateUser(user)
	return user, err
}

func (s *UserServiceStruct) GetAllUsers() ([]*models.UserStruct, error) {
	users, err := s.Repo.GetAllUsers()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve users: %w", err)
	}
	return users, nil
}
