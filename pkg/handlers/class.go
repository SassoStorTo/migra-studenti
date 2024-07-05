package handlers

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/SassoStorTo/migra-studenti/pkg/models"
	"github.com/SassoStorTo/migra-studenti/pkg/services/classes"
	impo_service "github.com/SassoStorTo/migra-studenti/pkg/services/import"
	"github.com/SassoStorTo/migra-studenti/pkg/services/majors"
	"github.com/SassoStorTo/migra-studenti/pkg/services/studentclass"
	"github.com/SassoStorTo/migra-studenti/pkg/services/students"
	"github.com/SassoStorTo/migra-studenti/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(400).SendString("File upload error")
	}

	file_path := "/tmp/uploads/" + file.Filename

	// Save the file to the server
	err = c.SaveFile(file, file_path)
	if err != nil {
		return c.Status(500).SendString("Could not save file")
	}

	err = impo_service.ParseFile(file_path, time.Now().Year())
	os.Remove(file_path)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendString("File uploaded successfully")
}

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

	return c.Render("classes/com_info_display", fiber.Map{"Class": class})
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

func GetStudentClassMigration(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	active_students := students.GetAllActiveByClassId(id)
	class := models.GetClassById(id)

	return c.Render("classes/com_table_migration", fiber.Map{"Students": active_students, "Class": class})
}

func GetStudentClassMigrationEdit(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	allowed_id := strings.Split(c.FormValue("allowed-id"), ",")
	allowed_id = allowed_id[:len(allowed_id)-1]
	log.Println(allowed_id)

	class := models.GetClassById(id)
	active_students := students.GetAllActiveByClassId(id)

	return c.Render("classes/com_table_migration_edit", fiber.Map{"Students": active_students, "Class": class, "AllowedId": allowed_id})
}

func GetTablesStudentsOfClass(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	class := models.GetClassById(id)
	active_students := students.GetAllActiveByClassId(id)
	log.Println(active_students)

	return c.Render("classes/com_table_students", fiber.Map{"Students": active_students, "Class": class})
}

func ClassMigrationRefreshPage(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	allowed_id := strings.Split(c.FormValue("allowed-id"), ",")
	allowed_id = allowed_id[:len(allowed_id)-1]
	log.Println(allowed_id)

	class := models.GetClassById(id)
	active_students := students.GetAllActiveByClassId(id)

	if class.Year >= 5 {
		c.Response().Header.Add("HX-Redirect", "/classes/"+strconv.Itoa(id))
		return c.SendStatus(fiber.StatusBadRequest)
	}

	class.Year++
	class.Id = 0
	class.ScholarYearStart++
	class.Save()

	allCurrentClasses := classes.GetAllWithMajors()
	previousClassId := -1

	for _, curr := range *allCurrentClasses {
		if curr.Year == class.Year-1 && curr.Section == class.Section &&
			curr.ScholarYearStart == class.ScholarYearStart && curr.IdMajor == class.IdMajor {
			previousClassId = curr.Id
			break
		}
		if curr.Year == class.Year-1 && curr.Section == class.Section &&
			curr.ScholarYearStart == class.ScholarYearStart+1 && curr.IdMajor == class.IdMajor {
			previousClassId = curr.Id
			break
		}
	}

	if previousClassId == -1 {
		err = classes.Create(class.Year, class.Section, class.ScholarYearStart+1, class.IdMajor)
		if err != nil {
			return err
		}
	}

	for _, stud := range *active_students {
		if !utils.IsItemInList(stud.Id, allowed_id) {
			err = studentclass.Create(stud.Id, previousClassId)
			if err != nil {
				return err
			}
		}
	}

	c.Response().Header.Add("HX-Redirect", "/classes/"+strconv.Itoa(id))
	return c.SendStatus(fiber.StatusOK)
}
