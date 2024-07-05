package users

import (
	"log"

	"github.com/SassoStorTo/migra-studenti/pkg/database"
	"github.com/SassoStorTo/migra-studenti/pkg/models"
)

func QueryCreate() string {
	return `
		CREATE TABLE IF NOT EXISTS Users (
			Id SERIAL PRIMARY KEY,
			Name VARCHAR(50),
			Email VARCHAR(150) NOT NULL UNIQUE,
			Hd VARCHAR(100) NOT NULL,
			VerifiedEmail BOOL NOT NULL,
			Picture VARCHAR(500),
			IsAdmin BOOL DEFAULT FALSE,
			IsEditor BOOL DEFAULT FALSE
		)`
}

func GetAll() *[]models.User {
	rows, err := database.DB.Query(`SELECT Id, Name, Email, Hd, VerifiedEmail, IsAdmin, IsEditor FROM Users;`)
	if err != nil {
		log.Panic(err.Error())
		return nil
	}
	defer rows.Close()

	data := []models.User{}
	for rows.Next() {
		var result models.User
		err := rows.Scan(&result.Id, &result.Name, &result.Email, &result.Hd,
			&result.VerifiedEmail, &result.IsAdmin, &result.IsEditor)
		if err != nil {
			log.Panic(err.Error())
			return nil
		}
		data = append(data, result)
	}

	return &data
}
