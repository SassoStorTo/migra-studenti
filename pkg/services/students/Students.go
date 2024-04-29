package students

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/SassoStorTo/studenti-italici/pkg/database"
	"github.com/SassoStorTo/studenti-italici/pkg/models"
	"github.com/SassoStorTo/studenti-italici/pkg/utils"
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

func Create(name, lastname string, dateOfBirth time.Time) (*models.Student, error) {
	s := models.NewStuent(name, lastname, dateOfBirth)
	return s, s.Save()
}

func Delete(id int) error {
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

type ClassStudent struct {
	Year     int
	Section  string
	Major    string
	Students *[]models.Student
}

type studentClassRead struct {
	ClassId sql.NullInt64
	Year    sql.NullInt64
	Section sql.NullString
	Major   sql.NullString
	models.Student
}

func GetAll() *[]studentClassRead {
	rows, err := database.DB.Query(`SELECT S.Id, S.Name, S.LastName, S.DateOfBirth
									FROM students AS S;`)
	if err != nil {
		log.Panic(err.Error())
		return nil
	}
	defer rows.Close()

	data := []studentClassRead{}
	for rows.Next() {
		var result studentClassRead
		err := rows.Scan(&result.Id, &result.Name, &result.LastName, &result.DateOfBirth)
		if err != nil {
			log.Panic(err.Error())
			return nil
		}
		data = append(data, result)
	}

	return &data
}

func GetAllWithClass() *[]studentClassRead {
	rows, err := database.DB.Query(`SELECT S.Id, S.Name, S.LastName, S.DateOfBirth, C.Id AS cid, C.Year, C.Section, M.Name
									FROM students AS S LEFT JOIN
										 studentclass AS SC ON S.Id = SC.IdS LEFT JOIN
										 classes AS C ON SC.IdC = C.Id LEFT JOIN
										 majors AS M ON C.IdM = M.Id;`)
	if err != nil {
		log.Panic(err.Error())
		return nil
	}
	defer rows.Close()

	data := []studentClassRead{}
	for rows.Next() {
		var result studentClassRead
		err := rows.Scan(&result.Id, &result.Name, &result.LastName, &result.DateOfBirth,
			&result.ClassId, &result.Year, &result.Section, &result.Major)
		if err != nil {
			log.Panic(err.Error())
			return nil
		}
		data = append(data, result)
	}

	return &data
}

func GetAllAssociatedClass() *map[string]ClassStudent {
	rows, err := database.DB.Query(`SELECT S.Id, S.Name, S.LastName, S.DateOfBirth, C.Id AS Cid, C.Year, C.Section, M.Name
									FROM students AS S INNER JOIN
									     studentclass AS SC ON S.Id = SC.IdS INNER JOIN
										 classes AS C ON SC.IdC = C.Id INNER JOIN
										 majors AS M ON C.IdM = M.Id;`)

	if err != nil {
		log.Panic(err.Error())
		return nil
	}
	defer rows.Close()

	data := make(map[string]ClassStudent)
	for rows.Next() {
		var result studentClassRead
		err := rows.Scan(&result.Id, &result.Name, &result.LastName, &result.DateOfBirth,
			&result.ClassId, &result.Year, &result.Section, &result.Major)
		if err != nil {
			log.Panic(err.Error())
			return nil
		}

		idx := fmt.Sprintf("%d%s%s", result.Year, result.Section, result.Major)
		if _, isKey := data[idx]; !isKey {
			data[idx] = ClassStudent{
				Year: utils.ConvertNullInt64ToInt(result.Year), Section: utils.ConvertNullStringToString(result.Section),
				Major: utils.ConvertNullStringToString(result.Major), Students: &[]models.Student{}}
		}

		*data[idx].Students = append(*data[idx].Students, models.Student{Id: result.Id,
			Name: result.Name})
	}

	return &data
}

func GetAssociatedClass(idClass int) *ClassStudent {
	rows, err := database.DB.Query(`SELECT S.Id, S.Name, S.LastName, S.DateOfBirth, C.Id AS Cid, C.Year, C.Section, M.Name
									FROM students AS S INNER JOIN
									     studentclass AS SC ON S.Id = SC.IdS INNER JOIN
										 classes AS C ON SC.IdC = C.Id INNER JOIN
										 majors AS M ON C.IdM = M.Id
									WHERE C.Id = ($1);`, idClass)

	if err != nil {
		log.Panic(err.Error())
		return nil
	}
	defer rows.Close()

	var data ClassStudent
	if rows.Next() {
		var result studentClassRead
		err := rows.Scan(&result.Id, &result.Name, &result.LastName, &result.DateOfBirth,
			&result.ClassId, &result.Year, &result.Section, &result.Major)
		if err != nil {
			log.Panic(err.Error())
			return nil
		}
	}

	return &data
}

func GetAllByClassId(idClass int) *[]models.Student {
	rows, err := database.DB.Query(`SELECT S.Id, S.Name, S.LastName, S.DateOfBirth
									FROM students AS S INNER JOIN
										 studentclass AS SC ON S.Id = SC.IdS
									WHERE SC.IdC = ($1);`, idClass)
	if err != nil {
		log.Panic(err.Error())
		return nil
	}
	defer rows.Close()

	data := []models.Student{}
	for rows.Next() {
		var result models.Student
		err := rows.Scan(&result.Id, &result.Name, &result.LastName, &result.DateOfBirth)
		if err != nil {
			log.Panic(err.Error())
			return nil
		}
		data = append(data, result)
	}

	return &data
}

func GetLastStudentId() int {
	rows, err := database.DB.Query(`SELECT Id FROM students ORDER BY Id DESC LIMIT 1`)
	if err != nil {
		log.Panic(err.Error())
		return 0
	}
	defer rows.Close()

	var id int
	if rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			log.Panic(err.Error())
			return 0
		}
	}

	return id
}

func GetCurrentMajor(studnetId int) (*models.Majors, error) { // todo: testare questa con piu' di un link a class
	rows, err := database.DB.Query(`SELECT M.Id, M.Name, MAX(SC.CreationDate) AS CreationDate
									FROM students AS S INNER JOIN
										 studentclass AS SC ON S.Id = SC.IdS INNER JOIN
										 classes AS C ON SC.IdC = C.Id INNER JOIN
										 majors AS M ON C.IdM = M.Id
									WHERE S.Id = ($1)
									GROUP BY M.Id, M.Name;`, studnetId)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	var major models.Majors

	if rows.Next() {
		var s time.Time
		err := rows.Scan(&major.Id, &major.Name, &s)
		if err != nil {
			log.Panic(err.Error())
			return nil, err
		}
	}

	return &major, nil
}
