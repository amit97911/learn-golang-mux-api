package repositories

import (
	"database/sql"
	"learn-golang-mux-api/internal/models"

	"log"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(databaseUrl string) *UserRepository {
	db, err := sql.Open("mysql", databaseUrl)
	if err != nil {
		log.Fatal("Failed to connect database!", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	return &UserRepository{DB: db}
}

func (repo *UserRepository) CreateUser(user *models.User) error {
	query := "INSERT INTO users (name, email) VALUES (?, ?)"
	_, err := repo.DB.Exec(query, user.Name, user.Email)
	return err
}

func (repo *UserRepository) GetUser(id uint) (*models.User, error) {
	query := "SELECT id, name, email FROM users WHERE id = ?"
	row := repo.DB.QueryRow(query, id)

	user := models.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) GetAllUsers() ([]*models.User, error) {
	query := "SELECT id, name, email FROM users"
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*models.User
	for rows.Next() {
		var user = &models.User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}
