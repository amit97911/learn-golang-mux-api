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
	userRepo := repositories.NewUserRepository(cfg.DatabaseURL)
	userService := services.UserService(userRepo)
	userHandler := handlers.UserHandler(userService)

	authService := services.AuthUserService(userRepo)
	authHandler := handlers.AuthUserHandler(authService)

	router := gorillaMux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()

	authRouter := apiRouter.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/login", authHandler.Login).Methods("POST")

	userRouter := apiRouter.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/create", userHandler.CreateUser).Methods("POST")
	userRouter.HandleFunc("/id/{id}", userHandler.GetUser).Methods("GET")
	userRouter.HandleFunc("/all", userHandler.GetAllUsers).Methods("GET")

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
