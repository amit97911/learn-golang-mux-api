package repositories

import (
	"database/sql"
	"learn-golang-mux-api/internal/models"
	"learn-golang-mux-api/pkg"
	"log"
)

func (db *DatabaseConnection) GetAllBooks(pagination *pkg.LimitOffset, search string) ([]*models.BookDetailsStruct, error) {
	var (
		rows *sql.Rows
		err  error
	)
	if search == "" {
		query := "SELECT id, title, author, year, description FROM books limit ? offset ?"
		rows, err = db.DB.Query(query, pagination.Limit, pagination.Offset)
	} else {
		query := "SELECT id, title, author, year, description FROM books WHERE title LIKE ? OR author LIKE ? OR description LIKE ? limit ? offset ?"
		likePattern := "%" + search + "%"
		rows, err = db.DB.Query(query, likePattern, likePattern, likePattern, pagination.Limit, pagination.Offset)
	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var books []*models.BookDetailsStruct
	for rows.Next() {
		var book = &models.BookDetailsStruct{}
		if err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.Year, &book.Description); err != nil {
			log.Println(err)
			return books, err
		}
		books = append(books, book)
	}
	return books, nil

}

func (db *DatabaseConnection) CreateBook(book *models.BookDetailsStruct) (int64, error) {
	query := "INSERT INTO books (title, author, year, description) VALUES (?, ?, ?, ?)"
	result, err := db.DB.Exec(query, book.Title, book.Author, book.Year, book.Description)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (db *DatabaseConnection) GetBookById(id string) (*models.BookDetailsStruct, error) {
	query := "SELECT id, title, author, year, description FROM books WHERE id = ?"
	row := db.DB.QueryRow(query, id)
	var book = &models.BookDetailsStruct{}
	if err := row.Scan(&book.Id, &book.Title, &book.Author, &book.Year, &book.Description); err != nil {
		return nil, err
	}
	return book, nil
}

func (db *DatabaseConnection) UpdateBook(book *models.BookDetailsStruct) (int64, error) {
	query := "UPDATE books SET title = ?, author = ?, year = ?, description = ? WHERE id = ?"
	res, err := db.DB.Exec(query, book.Title, book.Author, book.Year, book.Description, book.Id)
	if err != nil {
		return 0, err
	}
	updated, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return updated, nil
}

func (db *DatabaseConnection) DeleteBook(id string) (int64, error) {
	query := "DELETE FROM books WHERE id = ?"
	res, err := db.DB.Exec(query, id)
	if err != nil {
		return 0, err
	}
	affectedRows, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return affectedRows, nil
}
