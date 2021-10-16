package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func CreateConnection() (*sql.DB, error) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	// Open the connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		panic(err)
	}

	// check the connection
	start := time.Now()

	for db.Ping() != nil {
		if start.After(start.Add(10 * time.Second)) {
			log.Fatalf("Failed to connect to database: %v", err)
			break
		}
	}

	fmt.Println("Successfully connected!")

	return db, err
}
