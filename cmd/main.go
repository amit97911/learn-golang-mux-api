package main

import (
	"learn-golang-mux-api/config"
	"learn-golang-mux-api/internal/handlers"
	"learn-golang-mux-api/internal/repositories"
	"learn-golang-mux-api/internal/services"
	"learn-golang-mux-api/middlewares"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	gorillaHandlers "github.com/gorilla/handlers"
	gorillaMux "github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var (
	environment string
	logDir      string = "storage/logs/development"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	environment = config.GetEnv("ENVIRONMENT", "development")
	if environment == "production" {
		logDir = "storage/logs/production"
	} else {
		logDir = "storage/logs/development"
	}

	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		log.Fatal("Failed to create log directory:", err)
	}

	currentDate := time.Now().Format("2006-01-02")
	logFilename := filepath.Join(logDir, currentDate+".log")

	file, err := os.OpenFile(logFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}

	log.SetOutput(file)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
}

func main() {
	cfg := config.LoadConfig()
	db := repositories.DBConnect(cfg.DatabaseURL)

	userService := services.NewUserService(db)
	authService := services.NewAuthUserService(db)
	bookService := services.NewBookService(db)

	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthUserHandler(authService)
	bookHandler := handlers.NewBookHandler(bookService)

	router := gorillaMux.NewRouter()

	apiRouter := router.PathPrefix("/api").Subrouter()

	authRouter := apiRouter.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/login", authHandler.Login).Methods("POST")

	userRouter := apiRouter.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/create", userHandler.CreateUser).Methods("POST")

	protectedUserRouter := apiRouter.PathPrefix("/user").Subrouter()
	protectedUserRouter.Use(middlewares.AuthMiddleware)
	protectedUserRouter.HandleFunc("/id/{id}", userHandler.GetUser).Methods("GET")
	protectedUserRouter.HandleFunc("/all", userHandler.GetAllUsers).Methods("GET")

	protectedBookRouter := apiRouter.PathPrefix("/book").Subrouter()
	protectedBookRouter.Use(middlewares.AuthMiddleware)
	protectedBookRouter.HandleFunc("/all", bookHandler.GetAllBooks).Methods("GET")
	protectedBookRouter.HandleFunc("/id/{id}", bookHandler.GetBook).Methods("GET")
	protectedBookRouter.HandleFunc("/create", bookHandler.CreateBook).Methods("POST")
	protectedBookRouter.HandleFunc("/update", bookHandler.UpdateBook).Methods("PUT")
	protectedBookRouter.HandleFunc("/delete/{id}", bookHandler.DeleteBook).Methods("DELETE")

	port := cfg.Port
	log.Println("Server started on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, gorillaHandlers.CORS()(router)))
	middlewares.Middleware()
}
