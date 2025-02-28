package services

import (
	"learn-golang-mux-api/internal/repositories"
)

func AuthUserService(repo *repositories.UserRepository) *UserServiceStruct {
	return &UserServiceStruct{Repo: repo}
}
