package config

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var DB *sql.DB

func connectDB() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	connStr := fmt.Sprintf("postgress://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser, dbPassword, dbHost, dbPort, dbName, dbSSLMode)

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Error pinging database")
	}
	log.Println("Connected to database successfully")
}
