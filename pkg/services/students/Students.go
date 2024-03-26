package students

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/SassoStorTo/studenti-italici/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func QueryCreate() string {
	return `
		CREATE TABLE IF NOT EXISTS Students (
			Id SERIAL PRIMARY KEY,
			Name varchar(50) NOT NULL,
			LastName varchar(50) NOT NULL,
			DateOfBirth TIMESTAMP NOT NULL
		);`
}

func Create(c *fiber.Ctx) error {
	fmt.Print("Student Create\n")

	name := strings.TrimSpace(c.FormValue("name"))
	if name == "" {
		return fmt.Errorf("[Students] Create: name empty")
	}
	lastname := strings.TrimSpace(c.FormValue("lastname"))
	if lastname == "" {
		return fmt.Errorf("[Students] Create: lastname empty")
	}

	year, err := strconv.Atoi(c.FormValue("year"))
	if err != nil {
		return fmt.Errorf("[Students] Create: year incorrect")
	}
	month, err := strconv.Atoi(c.FormValue("month"))
	if err != nil {
		return fmt.Errorf("[Students] Create: month incorrect")
	}
	day, err := strconv.Atoi(c.FormValue("day"))
	if err != nil {
		return fmt.Errorf("[Students] Create: day incorrect")
	}
	dateOfBirth := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)

	s := models.NewStuent(name, lastname, dateOfBirth)
	return s.Save()
}

func Delete(c *fiber.Ctx) error {
	fmt.Print("Student Delete\n")

	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		return err
	}

	s := models.Student{Id: id}
	return s.Delete()
}

func Edit(c *fiber.Ctx) error {
	fmt.Print("Student Edit\n")

	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		return err
	}

	name := strings.TrimSpace(c.FormValue("name"))
	if name == "" {
		return fmt.Errorf("[Students] Create: name empty")
	}
	lastname := strings.TrimSpace(c.FormValue("lastname"))
	if lastname == "" {
		return fmt.Errorf("[Students] Create: lastname empty")
	}

	year, err := strconv.Atoi(c.FormValue("year"))
	if err != nil {
		return fmt.Errorf("[Students] Create: year incorrect")
	}
	month, err := strconv.Atoi(c.FormValue("month"))
	if err != nil {
		return fmt.Errorf("[Students] Create: month incorrect")
	}
	day, err := strconv.Atoi(c.FormValue("day"))
	if err != nil {
		return fmt.Errorf("[Students] Create: day incorrect")
	}
	dateOfBirth := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)

	s := models.Student{Id: id, Name: name, LastName: lastname,
		DateOfBirth: dateOfBirth}
	return s.Update()
}
