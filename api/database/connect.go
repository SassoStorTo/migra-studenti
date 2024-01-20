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
	// username, password, dbname, host, err := getCredentialsFromConfig("configs/database.conf")
	/////////////////////////////////

	// username, password, _, host, err := getCredentialsFromConfig("configs/database.conf")

	// if err != nil {
	// 	return fmt.Errorf("failed to read config file: %v", err)
	// }

	// connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", username, password, host, database)
	// connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", "paolomagnani", "p", "localhost", "")

	// connStr := fmt.Sprintf("postgres://%s:%s@%s?sslmode=disable", username, password, host)
	// alternative
	// connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, 5432, username, password, dbname)
	// connStr := fmt.Sprintf("port=%d user=%s password=%s dbname=%s sslmode=disable", 5432, username, password, dbname)
	// connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, 5432, username, password)

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "paolomagnani", "p", "italico")

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
