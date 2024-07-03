package studentclass

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/SassoStorTo/migra-studenti/pkg/database"
	"github.com/SassoStorTo/migra-studenti/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func QueryCreate() string {
	return `
		CREATE TABLE IF NOT EXISTS StudentClass (
			IdS INT,
			IdC INT, 
			CreationDate TIMESTAMP NOT NULL,
			PRIMARY KEY (IdS, IdC, CreationDate),
			FOREIGN KEY (IdS) REFERENCES Students(Id),
			FOREIGN KEY (IdC) REFERENCES Classes(Id)
		);`
}

func Create(idStudent, idClass int) error {
	s := models.NewStudentClass(idStudent, idClass, time.Now())
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
	CreationDate time.Time
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

		fmt.Printf("sandro: %s", result.CreationDate)

		data = append(data, result)
	}

	return &data
}
