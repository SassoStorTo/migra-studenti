package majors

import (
	"log"

	"github.com/SassoStorTo/studenti-italici/api/database"
	"github.com/SassoStorTo/studenti-italici/pkg/models"
)

func QueryCreate() string {
	return `
		CREATE TABLE IF NOT EXISTS Majors (
			Id SERIAL PRIMARY KEY,
			Name varchar(50) UNIQUE NOT NULL
		);
	`
}

func GetAll() *[]models.Majors {
	rows, err := database.DB.Query(`SELECT (Id, Name) FROM majors;`)

	if err != nil {
		log.Panic(err.Error())
	}
	defer rows.Close()

	var data []models.Majors
	for rows.Next() {
		var result models.Majors
		err := rows.Scan(&result.Id, &result.Name)
		if err != nil {
			log.Panic("rotto mentre lettura azzzz")
		}
		data = append(data, result)
	}

	return &data
}
