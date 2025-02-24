package main

import (
	"fmt"
	"learn-golang-mux-api/config"
	"learn-golang-mux-api/internal/handlers"
	"learn-golang-mux-api/internal/models"
	"learn-golang-mux-api/internal/repositories"
	"learn-golang-mux-api/internal/services"
	"learn-golang-mux-api/middlewares"
	"learn-golang-mux-api/pkg"
)

func main() {
	fmt.Println("main.go")
	config.Config()
	models.User()
	repositories.UserRepository()
	services.UserService()
	handlers.UserHandler()
	middlewares.Middleware()
	pkg.Utils()
}
