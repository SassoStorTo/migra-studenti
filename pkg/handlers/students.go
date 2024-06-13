package handlers

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/SassoStorTo/studenti-italici/pkg/models"
	"github.com/SassoStorTo/studenti-italici/pkg/services/classes"
	"github.com/SassoStorTo/studenti-italici/pkg/services/studentclass"
	"github.com/SassoStorTo/studenti-italici/pkg/services/students"
	"github.com/gofiber/fiber/v2"
)

// func GetTablesStudents(c *fiber.Ctx) error {
// 	classesWithStudents := students.GetAllAssociatedClass()

// 	if len(*classesWithStudents) == 0 {
// 		log.Print("No students")
// 		return c.Render("render_helper", fiber.Map{"content": "Non ci sono studenti"}, "template")
// 	}

//		return c.Render("students/table_view_students", fiber.Map{"Classess": classesWithStudents, "TableTitle": "Tutti gli studenti"}, "template")
//	}

func GetTablesStudents(c *fiber.Ctx) error {
	allStudent := students.GetAllWithClass()
	// allStudent := students.GetAll()

	if len(*allStudent) == 0 {
		log.Print("No students")
		return c.Render("render_helper", fiber.Map{"content": "Non ci sono studenti"}, "template")
	}

	return c.Render("students/table_view", fiber.Map{"Students": allStudent}, "template")
}

func AddNewStudent(c *fiber.Ctx) error {
	name := strings.TrimSpace(c.FormValue("name"))
	if name == "" {
		return fmt.Errorf("[Students] Create: name empty")
	}
	lastname := strings.TrimSpace(c.FormValue("lastname"))
	if lastname == "" {
		return fmt.Errorf("[Students] Create: lastname empty")
	}

	dateStr := c.FormValue("date")
	// dateOfBirth, err := time.Parse("aaaa-mm-dd", dateStr)
	dateOfBirth, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return fmt.Errorf("[Students] Create: date incorrect")
	}

	classId, err := strconv.Atoi(c.FormValue("idclass"))
	if err != nil {
		return fmt.Errorf("[Students] Create: class id incorrect")
	}

	class := models.GetClassById(classId)
	if class == nil {
		return fmt.Errorf("[Students] Create: class not found")
	}

	stud := models.NewStuent(name, lastname, dateOfBirth)
	err = stud.Save()
	if err != nil {
		return err
	}
	stud.Id = students.GetLastId()

	err = studentclass.Create(stud.Id, classId)
	if err != nil {
		return err
	}

	c.Response().Header.Add("HX-Redirect", "/students")
	return c.SendStatus(fiber.StatusOK)
}

func GetCreateStuduentForm(c *fiber.Ctx) error {
	class := models.GetAllClasses()
	return c.Render("students/form_create", fiber.Map{"Classes": class}, "template")
}

func GetStudentInfo(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	stud := models.GetStudentById(id)
	if stud == nil {
		return fmt.Errorf("student not found")
	}

	class := classes.GetByStudentID(stud.Id)
	history := studentclass.GetStudentHistory(id)

	return c.Render("students/info", fiber.Map{"Student": stud, "History": history,
		"Class": class}, "template")
}

func GetFomrComponentEditStudent(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	stud := models.GetStudentById(id)
	if stud == nil {
		return fmt.Errorf("student not found")
	}
	currClass := classes.GetByStudentID(stud.Id)

	allClasses := classes.GetAllWithMajors()

	return c.Render("students/com_info_form", fiber.Map{"Student": stud,
		"Classes": allClasses, "CurrentClass": currClass})
}

func GetFomrComponentDisplayStudent(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	stud := models.GetStudentById(id)
	if stud == nil {
		return fmt.Errorf("student not found")
	}

	class := classes.GetByStudentID(stud.Id)

	return c.Render("students/com_info_display", fiber.Map{"Student": stud, "Class": class})
}

func SaveEditStudent(c *fiber.Ctx) error {
	idStr := c.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("[Students] update: id incorrect")
	}

	name := strings.TrimSpace(c.FormValue("name"))
	if name == "" {
		return fmt.Errorf("[Students] update: name empty")
	}
	lastname := strings.TrimSpace(c.FormValue("lastname"))
	if lastname == "" {
		return fmt.Errorf("[Students] update: lastname empty")
	}

	dateStr := c.FormValue("date")
	dateOfBirth, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return fmt.Errorf("[Students] update: date incorrect")
	}

	stud := models.GetStudentById(id)
	if stud == nil {
		return fmt.Errorf("[Students] update: student not found")
	}

	stud.Name = name
	stud.LastName = lastname
	stud.DateOfBirth = dateOfBirth
	stud.Update()

	// class := classes.GetByStudentID(id)

	idClassStr := c.FormValue("idclass")
	newClassId, err := strconv.Atoi(idClassStr)
	if err != nil {
		return fmt.Errorf("[Students] update: id incorrect")
	}

	class := models.GetClassById(newClassId)
	if class == nil {
		return fmt.Errorf("[Students] update: class not found")
	}

	err = studentclass.Create(stud.Id, class.Id)
	if err != nil {
		return fmt.Errorf("[Students] update: " + err.Error())
	}

	classWithMajor := classes.GetByStudentID(stud.Id)

	return c.Render("students/com_info_display", fiber.Map{"Student": stud, "Class": classWithMajor})
}

func DeleteStudent(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	students.Delete(id)

	c.Response().Header.Add("HX-Redirect", "/students")
	return c.SendStatus(fiber.StatusOK)
}
