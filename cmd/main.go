package main

import (
	"learn-golang-mux-api/config"
	"learn-golang-mux-api/internal/handlers"
	"learn-golang-mux-api/internal/repositories"
	"learn-golang-mux-api/internal/services"
	"learn-golang-mux-api/middlewares"
	"learn-golang-mux-api/pkg"
	"log"
	"net/http"

	gorillaHandlers "github.com/gorilla/handlers"
	gorillaMux "github.com/gorilla/mux"
)

func main() {
	cfg := config.LoadConfig()
	repo := repositories.NewUserRepository(cfg.DatabaseURL)
	service := services.NewUserService(repo)
	handler := handlers.NewUserHandler(service)

	router := gorillaMux.NewRouter()

	userRouter := router.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/create", handler.CreateUser).Methods("POST")
	userRouter.HandleFunc("/id/{id}", handler.GetUser).Methods("GET")
	userRouter.HandleFunc("/all", handler.GetAllUsers).Methods("GET")

	// Method Not Allowed Handler
	router.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method Not Allowed"))
	})

	port := cfg.Port
	log.Println("Server started on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, gorillaHandlers.CORS()(router)))
	middlewares.Middleware()
	pkg.Utils()
}
