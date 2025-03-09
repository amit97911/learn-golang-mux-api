package handlers

import (
	"encoding/json"
	"fmt"
	"learn-golang-mux-api/internal/models"
	"learn-golang-mux-api/internal/services"
	"learn-golang-mux-api/pkg"
	"net/http"

	"github.com/gorilla/mux"
)

type BookServiceStruct struct {
	Service *services.BookRepositoryStruct
}

func NewBookHandler(serv *services.BookRepositoryStruct) *BookServiceStruct {
	return &BookServiceStruct{Service: serv}
}

func (serv *BookServiceStruct) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	var (
		pagination *pkg.LimitOffset
		err        error
		query      string
		books      []*models.BookDetailsStruct
	)
	pagination, err = pkg.Paginate(r.URL.Query().Get("limit"), r.URL.Query().Get("offset"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	query = r.URL.Query().Get("query")

	books, err = serv.Service.GetAllBooks(pagination, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response := map[string]any{
		"message":    "Books fetched",
		"data":       books,
		"pagination": pagination,
	}
	json.NewEncoder(w).Encode(response)
}

func (serv *BookServiceStruct) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book *models.BookDetailsStruct

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if book.Title == "" || book.Author == "" || book.Year == 0 || book.Description == "" {
		response := map[string]string{
			"message": "title, author, year and description are required",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	createdBook, err := serv.Service.CreateBook(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]any{
		"message": "Book created",
		"data":    createdBook,
	}

	json.NewEncoder(w).Encode(response)

}

func (serv *BookServiceStruct) GetBook(w http.ResponseWriter, r *http.Request) {
	var (
		book *models.BookDetailsStruct
		err  error
	)
	vars := mux.Vars(r)
	id := vars["id"]

	book, err = serv.Service.GetBookById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response := map[string]any{
		"message": "Book fetched",
		"data":    book,
	}
	json.NewEncoder(w).Encode(response)
}

func (serv *BookServiceStruct) UpdateBook(w http.ResponseWriter, r *http.Request) {
	var (
		book *models.BookDetailsStruct
		err  error
	)
	book, err = serv.Service.GetBookById(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if book.Title == "" || book.Author == "" || book.Year == 0 || book.Description == "" {
		response := map[string]string{
			"message": "title, author, year and description are required",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	updatedBook, err := serv.Service.UpdateBook(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]any{
		"message": "Book updated",
		"data":    updatedBook,
	}

	json.NewEncoder(w).Encode(response)
}

func (serv *BookServiceStruct) DeleteBook(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	id := vars["id"]

	err = serv.Service.DeleteBook(id)
	// err = serv.Service.DeleteBook(r.URL.Query().Get("id"))
	fmt.Println(r.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"message": "Book deleted",
	}

	json.NewEncoder(w).Encode(response)
}
