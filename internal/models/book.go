package models

type BookDetailsStruct struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Year        int    `json:"year"`
	Description string `json:"description"`
}
