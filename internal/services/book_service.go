package services

import (
	"errors"
	"learn-golang-mux-api/internal/models"
	"learn-golang-mux-api/internal/repositories"
	"learn-golang-mux-api/pkg"
	"log"
)

type BookRepositoryStruct struct {
	Repository *repositories.DatabaseConnection
}

func NewBookService(repo *repositories.DatabaseConnection) *BookRepositoryStruct {
	return &BookRepositoryStruct{Repository: repo}
}

func (repo *BookRepositoryStruct) GetAllBooks(pagination *pkg.LimitOffset, query string) ([]*models.BookDetailsStruct, error) {
	books, err := repo.Repository.GetAllBooks(pagination, query)
	if err != nil {
		log.Println("Start: GetAllBooks() service failed")
		log.Println(err)
		log.Println("End: GetAllBooks() service failed")
		return nil, err
	}
	if len(books) == 0 {
		return nil, errors.New("No books found")
	}
	return books, nil
}

func (repo *BookRepositoryStruct) CreateBook(book *models.BookDetailsStruct) (*models.BookDetailsStruct, error) {
	id, err := repo.Repository.CreateBook(book)
	book.Id = id
	if err != nil {
		log.Println("Start: CreateBook() service failed")
		log.Println(err)
		log.Println("End: CreateBook() service failed")
		return nil, err
	}
	return book, err
}

func (repo *BookRepositoryStruct) GetBookById(id string) (*models.BookDetailsStruct, error) {
	book, err := repo.Repository.GetBookById(id)
	if err != nil {
		log.Println("Start: GetBookById() service failed")
		log.Println(err)
		log.Println("End: GetBookById() service failed")
		return nil, err
	}
	return book, nil
}

func (repo *BookRepositoryStruct) UpdateBook(book *models.BookDetailsStruct) (*models.BookDetailsStruct, error) {
	rowsAffected, err := repo.Repository.UpdateBook(book)
	if err != nil {
		log.Println("Start: UpdateBook() service failed")
		log.Println(err)
		log.Println("End: UpdateBook() service failed")
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, errors.New("No book found")
	}

	return book, nil
}

func (repo *BookRepositoryStruct) DeleteBook(id string) error {
	rowsAffected, err := repo.Repository.DeleteBook(id)
	if err != nil {
		log.Println("Start: DeleteBook() service failed")
		log.Println(err)
		log.Println("End: DeleteBook() service failed")
		return err
	}

	if rowsAffected == 0 {
		return errors.New("No book found")
	}

	return nil
}
