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
