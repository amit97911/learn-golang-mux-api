package services

import (
	"errors"
	"fmt"
	"learn-golang-mux-api/internal/repositories"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceStruct struct {
	Repo *repositories.DatabaseConnection
}

func AuthUserService(repo *repositories.DatabaseConnection) *AuthServiceStruct {
	return &AuthServiceStruct{Repo: repo}
}

func (s *AuthServiceStruct) HandleLogin(email, password string) (*bool, error) {
	var userExists bool
	hashedPassword, err := s.Repo.AuthenticateUser(email, password)
	if err != nil {
		return &userExists, fmt.Errorf("failed to retrieve users: %w", err)
	}
	// Compare the provided password with the hashed password
	if err = bcrypt.CompareHashAndPassword([]byte(*hashedPassword), []byte(password)); err != nil {
		log.Println("Comparison failed")
		return &userExists, errors.New("authentication failed")
	}

	return &userExists, nil
}
