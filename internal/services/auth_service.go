package services

import (
	"errors"
	"fmt"
	"learn-golang-mux-api/internal/repositories"
	"learn-golang-mux-api/middlewares"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type AuthRepositoryStruct struct {
	Repository *repositories.DatabaseConnection
}

/**************************************************************************************/
func NewAuthUserService(repo *repositories.DatabaseConnection) *AuthRepositoryStruct {
	return &AuthRepositoryStruct{Repository: repo}
}

/**************************************************************************************/

func (repo *AuthRepositoryStruct) HandleLogin(email, password string) (*string, error) {
	var token string
	hashedPassword, err := repo.Repository.AuthenticateUser(email, password)
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
