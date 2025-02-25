package services

import (
	"errors"
	"learn-golang-mux-api/internal/models"
	"learn-golang-mux-api/internal/repositories"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (service *UserService) RegisterUser(name, email string) (*models.User, error) {
	if name == "" || email == "" {
		return nil, errors.New("name and email are required")
	}

	user := &models.User{Name: name, Email: email}
	err := service.Repo.CreateUser(user)
	return user, err
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	users, err := s.Repo.GetAllUsers()
	if err != nil {
		return nil, errors.New("failed to retrieve users")
	}
	return users, nil
}
