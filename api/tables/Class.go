package tables

import (
	"log"

	"github.com/SassoStorTo/studenti-italici/api/database"
	"github.com/SassoStorTo/studenti-italici/pkg/models"
)

func QueryCreateClasses() string {
	return `
		CREATE TABLE IF NOT EXISTS Classes (
			Id SERIAL PRIMARY KEY,
			Year INT NOT NULL CHECK (Year BETWEEN 1 AND 5), 
			Section VARCHAR(5) NOT NULL,
			ScholarYearStart INT NOT NULL CHECK (ScholarYearStart > 0),
			IdM INT,
			FOREIGN KEY (IdM) REFERENCES majors(Id)
		);`
}

func GetAllClasses() *[]models.Class {
	rows, err := database.DB.Query(`SELECT (Id, Year, Section, ScholarYearStart) FROM majors;`)

	if err != nil {
		log.Panic(err.Error())
	}
	defer rows.Close()

	var data []models.Class
	for rows.Next() {
		var result models.Class
		err := rows.Scan(&result.Id, &result.Year, &result.Section)
		if err != nil {
			log.Panic("rotto mentre lettura azzzz")
		}
		data = append(data, result)
	}

	return &data
}
