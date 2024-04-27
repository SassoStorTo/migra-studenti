package classes

import (
	"fmt"
	"log"
	"strconv"
	"strings"

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
	rows, err := database.DB.Query(`SELECT (Id, Year, Section, ScholarYearStart, IdMajor) 
		FROM majors;`)
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

func Edit(c *fiber.Ctx) error {
	fmt.Print("Class Edit\n")

	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		return fmt.Errorf("[Classes] Create: id field incorrect")
	}

	year, err := strconv.Atoi(c.FormValue("year"))
	if err != nil {
		return fmt.Errorf("[Classes] Create: year incorrect")
	}
	section := strings.TrimSpace(c.FormValue("section"))
	if section == "" {
		return fmt.Errorf("[Classes] Create: section empty")
	}
	schoolyear, err := strconv.Atoi(c.FormValue("schoolyear"))
	if err != nil {
		return fmt.Errorf("[Classes] Create: schoolyear incorrect")
	}
	idMajor, err := strconv.Atoi(c.FormValue("idmajor"))
	if err != nil {
		return fmt.Errorf("[Classes] Create: major id incorrect")
	}

	s := &models.Class{Id: id, Year: year, Section: section,
		ScholarYearStart: schoolyear, IdMajor: idMajor}
	return s.Update()
}
