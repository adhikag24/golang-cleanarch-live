package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func CreateConnection() *sql.DB {

	// Open the connection

	dbConn, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// check the connection
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")

	return dbConn
}
