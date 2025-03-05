package config

import (
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseURL string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	cfg := mysql.Config{
		User:                 GetEnv("DB_USER", "user"),
		Passwd:               GetEnv("DB_PASS", "passs"),
		Net:                  "tcp",
		Addr:                 GetEnv("DB_HOST", "host") + ":" + GetEnv("DB_PORT", "3306"),
		DBName:               GetEnv("DB_NAME", "dbname"),
		AllowNativePasswords: true,
	}
	return &Config{
		Port: GetEnv("PORT", "8080"),
		// DatabaseURL: getEnv("DATABASE_URL", "username:password@tcp(localhost:3306)/database"),
		DatabaseURL: cfg.FormatDSN(),
	}
}

// getEnv retrieves environment variables or sets a default
func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
