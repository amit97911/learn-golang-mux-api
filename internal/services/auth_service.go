package services

import (
	"errors"
	"fmt"
	"learn-golang-mux-api/internal/repositories"
	"learn-golang-mux-api/middlewares"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceStruct struct {
	Repo *repositories.DatabaseConnection
}

func AuthUserService(repo *repositories.DatabaseConnection) *AuthServiceStruct {
	return &AuthServiceStruct{Repo: repo}
}

func (s *AuthServiceStruct) HandleLogin(email, password string) (*string, error) {
	var token string
	hashedPassword, err := s.Repo.AuthenticateUser(email, password)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve users: %w", err)
	}

	// Compare the provided password with the hashed password
	if err = bcrypt.CompareHashAndPassword([]byte(*hashedPassword), []byte(password)); err != nil {
		log.Println("Comparison failed")
		return nil, errors.New("authentication failed")
	}
	token, err = middlewares.GenerateToken(email)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}
	return &token, nil
}
