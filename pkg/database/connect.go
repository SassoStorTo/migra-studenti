package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() error {
	if DB != nil {
		return errors.New("connection has already been established")
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Panicf("error opening connection to database: %v\n", err)
	}

	if err = db.Ping(); err != nil {
		db.Close()
		log.Panicf("error connecting to database: %v\n", err)
	}

	DB = db
	fmt.Println("Connected to the database!")

	return nil
}
