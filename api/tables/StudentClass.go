package tables

import (
	"log"

	"github.com/SassoStorTo/studenti-italici/api/database"
	"github.com/SassoStorTo/studenti-italici/pkg/models"
)

func QueryCreateStudentClass() string {
	return `
		CREATE TABLE IF NOT EXISTS StudentClass (
			IdS INT,
			IdC INT, 
			CreationDate TIMESTAMP NOT NULL,
			PRIMARY KEY (IdS, IdC),
			FOREIGN KEY (IdS) REFERENCES Students(Id),
			FOREIGN KEY (IdC) REFERENCES Classes(Id)
		);`
}

func CreateTableStudentClass() {
	_, err := database.DB.Exec(QueryCreateClasses())

	if err != nil {
		log.Panic(err)
	}
}

func GetAllStudentClass() *[]models.Class {
	rows, err := database.DB.Query(`SELECT (IdS, IdC, DateCreation) FROM StudentClass;`)

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
