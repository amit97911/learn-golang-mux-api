package repositories

import (
	"learn-golang-mux-api/internal/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(databaseUrl string) *UserRepository {
	db, err := gorm.Open(mysql.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database!")
	}
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	db.AutoMigrate(&models.User{})
	return &UserRepository{DB: db}
}

func (repo *UserRepository) CreateUser(user *models.User) error {
	return repo.DB.Create(user).Error
}

func (repo *UserRepository) GetUser(id uint) (*models.User, error) {
	var user models.User
	err := repo.DB.First(&user, id).Error
	return &user, err
}

func (repo *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := repo.DB.Find(&users).Error
	return users, err
}
