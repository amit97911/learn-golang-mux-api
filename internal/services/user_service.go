package services

import (
	"errors"
	"fmt"
	"learn-golang-mux-api/internal/models"
	"learn-golang-mux-api/internal/repositories"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (service *UserService) RegisterUser(name, email, password string) (*models.User, error) {
	if name == "" || email == "" || password == "" {
		return nil, errors.New("name and email are required")
	}

	user := &models.User{Name: name, Email: email, Password: password}
	err := service.Repo.CreateUser(user)
	return user, err
}

func (s *UserService) GetAllUsers() ([]*models.User, error) {
	users, err := s.Repo.GetAllUsers()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve users: %w", err)
	}
	return users, nil
}
