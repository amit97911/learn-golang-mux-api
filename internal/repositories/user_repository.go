package repositories

import (
	"database/sql"
	"errors"
	"learn-golang-mux-api/internal/models"
)

func (repo *DatabaseConnection) CreateUser(user *models.UserWithPasswordStruct) error {
	query := "INSERT INTO users (name, email,password) VALUES (?, ?, ?)"
	_, err := repo.DB.Exec(query, user.Name, user.Email, user.Password)
	return err
}

func (repo *DatabaseConnection) GetUser(id uint) (*models.UserStruct, error) {
	query := "SELECT id, name, email FROM users WHERE id = ?"
	row := repo.DB.QueryRow(query, id)

	user := models.UserStruct{}
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *DatabaseConnection) GetAllUsers() ([]*models.UserStruct, error) {
	query := "SELECT id, name, email FROM users"
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*models.UserStruct
	for rows.Next() {
		var user = &models.UserStruct{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (repo *DatabaseConnection) AuthenticateUser(email, password string) (*string, error) {
	var hashedPassword string

	// Use a prepared statement to prevent SQL injection
	query := "SELECT password FROM users WHERE email = ?"
	row := repo.DB.QueryRow(query, email)
	if err := row.Scan(&hashedPassword); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Failed")
		}
		return nil, err
	}
	return &hashedPassword, nil

}
