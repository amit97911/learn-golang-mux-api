package repositories

import (
	"database/sql"
	"log"
)

type DatabaseConnection struct {
	DB *sql.DB
}

func DBConnect(databaseUrl string) *DatabaseConnection {
	db, err := sql.Open("mysql", databaseUrl)
	if err != nil {
		log.Fatal("Failed to connect database!", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	return &DatabaseConnection{DB: db}
}
