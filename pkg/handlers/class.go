package handlers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/SassoStorTo/studenti-italici/pkg/models"
	"github.com/SassoStorTo/studenti-italici/pkg/services/classes"
	"github.com/SassoStorTo/studenti-italici/pkg/services/majors"
	"github.com/SassoStorTo/studenti-italici/pkg/services/students"
	"github.com/gofiber/fiber/v2"
)

func GetAllClasses(c *fiber.Ctx) error {
	classes := models.GetAllClasses()
	return c.Render("classes/table_view_classes", fiber.Map{"Classes": classes, "TableTitle": "Tutte le classi"}, "template")
}

func GetClassInfo(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fmt.Errorf("[Classes] GetClassInfo: id incorrect")
	}

	class := models.GetClassById(id)
	active_students := students.GetAllActiveByClassId(id)
	old_students := students.GetAllOldByClassId(id)

	return c.Render("classes/info", fiber.Map{"Class": class, "Students": active_students, "Old": old_students}, "template")
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

func SaveEditClass(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fmt.Errorf("[Classes] Edit: id field incorrect")
	}

	year, err := strconv.Atoi(c.FormValue("year"))
	if err != nil {
		return fmt.Errorf("[Classes] Edit: year incorrect")
	}
	section := strings.TrimSpace(c.FormValue("section"))
	if section == "" {
		return fmt.Errorf("[Classes] Edit: section empty")
	}
	schoolyear, err := strconv.Atoi(c.FormValue("scholaryearstart"))
	if err != nil {
		return fmt.Errorf("[Classes] Edit: schoolyear incorrect")
	}
	// idMajor, err := strconv.Atoi(c.FormValue("idmajor"))
	// if err != nil {
	// 	return fmt.Errorf("[Classes] Edit: major id incorrect")
	// }

	class := models.GetClassById(id)
	if class == nil {
		return fmt.Errorf("class not found")
	}
	class.Year = year
	class.Section = section
	class.ScholarYearStart = schoolyear
	// class.IdMajor = idMajor

	err = class.Update()
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	return c.Render("classes/com_info_form", fiber.Map{"Class": class})
}

func GetFomrComponentEditClass(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	class := models.GetClassById(id)
	if class == nil {
		return fmt.Errorf("student not found")
	}

	return c.Render("classes/com_info_form", fiber.Map{"Class": class})
}

func GetFomrComponentDisplayClass(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	class := models.GetClassById(id)
	if class == nil {
		return fmt.Errorf("student not found")
	}

	return c.Render("classes/com_info_display", fiber.Map{"Class": class})
}
