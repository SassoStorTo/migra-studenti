package studentclass

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/SassoStorTo/studenti-italici/pkg/database"
	"github.com/SassoStorTo/studenti-italici/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func QueryCreate() string {
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

func CreateTableStudentClass() { //Todo: chek if it's really used
	_, err := database.DB.Exec(QueryCreate())

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

func Create(c *fiber.Ctx) error {
	fmt.Print("student-class Create\n")

	ids, err := strconv.Atoi(c.FormValue("ids"))
	if err != nil {
		return fmt.Errorf("[Classes] Create: ids field incorrect")
	}
	idc, err := strconv.Atoi(c.FormValue("idc"))
	if err != nil {
		return fmt.Errorf("[Classes] Create: idc field incorrect")
	}

	s := models.NewStudentClass(ids, idc, time.Now())
	return s.Save()
}
func Delete(c *fiber.Ctx) error {
	fmt.Print("student-class Delete\n")

	ids, err := strconv.Atoi(c.FormValue("ids"))
	if err != nil {
		return fmt.Errorf("[Classes] Delete: ids field incorrect")
	}
	idc, err := strconv.Atoi(c.FormValue("idc"))
	if err != nil {
		return fmt.Errorf("[Classes] Delete: idc field incorrect")
	}

	s := models.StudentClass{IdS: ids, IdC: idc}
	return s.Delete()
}

type ClassWithMajor struct {
	models.Class
	Major        string
	CreationDate string
}

func GetStudentHistory(id int) *[]ClassWithMajor {
	rows, err := database.DB.Query(`SELECT C.Id, C.Year, C.Section, C.ScholarYearStart, M.Name, SC.CreationDate
									FROM studentclass AS SC INNER JOIN
										 classes AS C ON SC.IdC = C.Id INNER JOIN
										 majors AS M ON C.IdM = M.Id
									WHERE IdS = $1
									ORDER BY SC.CreationDate DESC;`, id)

	if err != nil {
		log.Panic(err.Error())
	}
	defer rows.Close()

	var data []ClassWithMajor
	for rows.Next() {
		var result ClassWithMajor
		err := rows.Scan(&result.Id, &result.Year, &result.Section,
			&result.ScholarYearStart, &result.Major, &result.CreationDate)
		if err != nil {
			log.Panic("rotto mentre lettura azzzz")
		}
		data = append(data, result)
	}

	return &data
}
