package classes

import (
	"fmt"
	"log"
	"strconv"

	"github.com/SassoStorTo/studenti-italici/pkg/database"
	"github.com/SassoStorTo/studenti-italici/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func QueryCreate() string {
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

func GetById(id int) *models.Class { //Todo: change name - usa bene il next
	rows, err := database.DB.Query(`SELECT (Id, Year, Section, ScholarYearStart, IdMajor) 
		FROM majors;`)
	if err != nil {
		log.Panic(err.Error())
	}
	defer rows.Close()

	rows.Next()
	var result models.Class
	err = rows.Scan(&result.Id, &result.Year, &result.Section,
		&result.ScholarYearStart, &result.IdMajor)
	if err != nil {
		log.Panic("rotto mentre lettura azzzz")
	}

	return &result
}

func GetAll() *[]models.Class {
	rows, err := database.DB.Query(`SELECT Id, Year, Section, ScholarYearStart, IdM
									FROM classes;`)
	if err != nil {
		log.Panic(err.Error())
	}
	defer rows.Close()

	var data []models.Class
	for rows.Next() {
		var result models.Class
		err := rows.Scan(&result.Id, &result.Year, &result.Section,
			&result.ScholarYearStart, &result.IdMajor)
		if err != nil {
			log.Panic("rotto mentre lettura azzzz")
		}
		data = append(data, result)
	}

	return &data
}

type ClassWithMajor struct {
	models.Class
	Major string
}

func GetAllWithMajors() *[]ClassWithMajor {
	rows, err := database.DB.Query(`SELECT C.Id, C.Year, C.Section, C.ScholarYearStart, C.IdM, M.Name
									FROM classes AS C INNER JOIN
										 majors AS M ON C.IdM = M.Id;`)
	if err != nil {
		log.Panic(err.Error())
	}
	defer rows.Close()

	var data []ClassWithMajor
	for rows.Next() {
		var result ClassWithMajor
		err := rows.Scan(&result.Id, &result.Year, &result.Section,
			&result.ScholarYearStart, &result.IdMajor, &result.Major)
		if err != nil {
			log.Panic(err.Error())
		}
		data = append(data, result)
	}

	return &data
}

func Create(year int, section string, schoolyear int, idMajor int) error {
	s := models.NewClass(year, section, schoolyear, idMajor)
	return s.Save()
}

func Delete(c *fiber.Ctx) error {
	fmt.Print("Class Delete\n")

	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		return fmt.Errorf("[Classes] Create: id field incorrect")
	}

	s := models.Class{Id: id}
	return s.Delete()
}

func GetByStudentID(studentId int) ClassWithMajor {
	rows, err := database.DB.Query(`SELECT C.Id, C.Year, C.Section, C.ScholarYearStart, C.IdM, M.Name
									FROM classes AS C INNER JOIN
										 majors AS M ON C.IdM = M.Id INNER JOIN
										 studentclass AS SC ON C.Id = SC.IdC
									WHERE SC.IdS = ($1)
									ORDER BY SC.CreationDate DESC
									LIMIT 1;`, studentId)
	if err != nil {
		log.Panic(err.Error())
	}
	defer rows.Close()

	var result ClassWithMajor
	rows.Next()
	err = rows.Scan(&result.Id, &result.Year, &result.Section,
		&result.ScholarYearStart, &result.IdMajor, &result.Major)
	if err != nil {
		log.Panic(err.Error())
	}

	return result
}

func GetLastId() int {
	rows, err := database.DB.Query(`SELECT Id FROM classes ORDER BY Id DESC LIMIT 1`)
	if err != nil {
		log.Panic(err.Error())
		return 0
	}
	defer rows.Close()

	var id int
	if rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			log.Panic(err.Error())
			return -1
		}
	}

	return id
}
