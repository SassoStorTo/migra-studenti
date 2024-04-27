package handlers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/SassoStorTo/studenti-italici/pkg/models"
	"github.com/SassoStorTo/studenti-italici/pkg/services/classes"
	"github.com/SassoStorTo/studenti-italici/pkg/services/majors"
	"github.com/gofiber/fiber/v2"
)

func GetAllClasses(c *fiber.Ctx) error {
	classes := models.GetAllClasses()
	return c.Render("classes/table_view_classes", fiber.Map{"Classes": classes, "TableTitle": "Tutte le classi"}, "template")
}

func GetCreateClassForm(c *fiber.Ctx) error {
	majors := majors.GetAll()
	currentYear := time.Now().Year()
	return c.Render("classes/add_new", fiber.Map{"Majors": majors, "DefaultYear": currentYear}, "template")
}

func AddNewClass(c *fiber.Ctx) error {
	class := new(models.Class)
	if err := c.BodyParser(class); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	year, err := strconv.Atoi(c.FormValue("year"))
	if err != nil {
		return fmt.Errorf("[Classes] Create: year incorrect")
	}
	section := strings.TrimSpace(c.FormValue("section"))
	if section == "" {
		return fmt.Errorf("[Classes] Create: section empty")
	}
	schoolyear, err := strconv.Atoi(c.FormValue("schoolyearstart"))
	if err != nil {
		return fmt.Errorf("[Classes] Create: schoolyear incorrect")
	}
	idMajor, err := strconv.Atoi(c.FormValue("idmajor"))
	if err != nil {
		return fmt.Errorf("[Classes] Create: major id incorrect")
	}

	err = classes.Create(year, section, schoolyear, idMajor)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	c.Response().Header.Add("HX-Redirect", "/classes")
	return c.SendStatus(fiber.StatusOK)
}
