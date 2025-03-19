package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get database credentials from env
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")


	// Construct connection string
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s ",
		dbHost, dbUser, dbPassword, dbName)

	// Connect to the database
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Verify connection
	if err = DB.Ping(); err != nil {
		log.Fatal("Database is not reachable:", err)
	}

	fmt.Println("Database connected!")
}
