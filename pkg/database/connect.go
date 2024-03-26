package database

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-ini/ini"
	_ "github.com/lib/pq"
)

func ConnectDB() error {
	if DB != nil {
		return errors.New("connection has already been established")
	}

	// username, password, database, host, err := getCredentialsFromConfig("configs/database.conf")
	// port := 4532

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "studenti-db", 5432, "paolomagnani", "p", "italico")
	// connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, database)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("error opening connection to database: %v", err)
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return fmt.Errorf("error connecting to database: %v", err) // change with log.panicf
	}

	fmt.Println("il ping e' andato in buca")

	DB = db
	fmt.Println("Connected to the database!")

	return nil
}

func getCredentialsFromConfig(path string) (string, string, string, string, error) {
	cfg, err := ini.Load(path)
	if err != nil {
		return "", "", "", "", fmt.Errorf("failed to read config file: %v", err)
	}

	postgresSection, err := cfg.GetSection("PostgreSQL")
	if err != nil {
		return "", "", "", "", fmt.Errorf("PostgreSQL section not found in config: %v", err)
	}

	return postgresSection.Key("username").String(),
		postgresSection.Key("password").String(),
		postgresSection.Key("database").String(),
		postgresSection.Key("host").String(),
		nil
}
