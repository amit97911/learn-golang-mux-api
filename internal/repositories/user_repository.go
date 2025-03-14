package repositories

import (
	"database/sql"
	"errors"
	"learn-golang-mux-api/internal/models"
)

func (db *DatabaseConnection) CreateUser(user *models.UserStruct) error {
	query := "INSERT INTO users (name, email,password) VALUES (?, ?, ?)"
	_, err := db.DB.Exec(query, user.Name, user.Email, user.Password)
	return err
}

func (db *DatabaseConnection) GetUser(id uint) (*models.UserDetailsStruct, error) {
	query := "SELECT id, name, email FROM users WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	user := models.UserDetailsStruct{}
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *DatabaseConnection) GetAllUsers() ([]*models.UserDetailsStruct, error) {
	query := "SELECT id, name, email FROM users"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*models.UserDetailsStruct
	for rows.Next() {
		var user = &models.UserDetailsStruct{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (db *DatabaseConnection) AuthenticateUser(email, password string) (*string, error) {
	var hashedPassword string
	// Use a prepared statement to prevent SQL injection
	query := "SELECT password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, email)
	if err := row.Scan(&hashedPassword); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Failed")
		}
		return nil, err
	}
	return &hashedPassword, nil

}
