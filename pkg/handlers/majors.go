package handlers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/SassoStorTo/studenti-italici/pkg/services/majors"
	"github.com/gofiber/fiber/v2"
)

func GetCreateMajorForm(c *fiber.Ctx) error {
	return c.Render("majors/add_new", fiber.Map{}, "template")
}
func GetTableMajors(c *fiber.Ctx) error {
	majors := *majors.GetAll()
	return c.Render("majors/table_view_majors", fiber.Map{"Majors": majors, "Title": "Articolazioni"}, "template")
}

func AddNewMajor(c *fiber.Ctx) error {
	name := strings.TrimSpace(c.FormValue("name"))
	if name == "" {
		return fmt.Errorf("[Classes] Create: name field empty")
	}
	m, err := majors.Create(name)
	if err != nil {
		return err
	}
	return c.SendString("Articolazione " + m.Name + " creata con successo")
}

func DeleteMajor(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fmt.Errorf("[Classes] Delete: id field incorrect")
	}

	majors.Delete(id)
	return c.Response().CloseBodyStream()
}
