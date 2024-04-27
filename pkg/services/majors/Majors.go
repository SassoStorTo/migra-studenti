package majors

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
		CREATE TABLE IF NOT EXISTS Majors (
			Id SERIAL PRIMARY KEY,
			Name varchar(50) UNIQUE NOT NULL
		);
	`
}

func GetAll() *[]models.Majors {
	rows, err := database.DB.Query(`SELECT Id, Name FROM majors;`)

	if err != nil {
		log.Panic(err.Error())
	}
	defer rows.Close()

	var data []models.Majors
	for rows.Next() {
		var result models.Majors
		err := rows.Scan(&result.Id, &result.Name)
		if err != nil {
			log.Panic(err.Error())
		}
		data = append(data, result)
	}

	return &data
}

func Create(name string) (*models.Majors, error) {
	s := models.NewMajor(name)
	return &s, s.Save()
}

func Delete(id int) error {
	s := &models.Majors{Id: id}
	return s.Delete()
}

func Edit(c *fiber.Ctx) error {
	fmt.Print("Major Edit\n")

	name := strings.TrimSpace(c.FormValue("name"))
	if name == "" {
		return fmt.Errorf("[Classes] Edit: name field empty")
	}

	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		return fmt.Errorf("[Classes] Edit: id field incorrect")
	}

	s := &models.Majors{Id: id, Name: name}
	return s.Update()
}
