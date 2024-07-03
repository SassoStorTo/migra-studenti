package models

import (
	"errors"
	"fmt"
	"log"

	"github.com/SassoStorTo/migra-studenti/pkg/database"
)

type Class struct {
	Id               int
	Year             int
	Section          string
	ScholarYearStart int
	IdMajor          int
}

func NewClass(year int, section string, scholarYearStart int, idm int) *Class {
	return &Class{
		Id:               -1,
		Year:             year,
		Section:          section,
		ScholarYearStart: scholarYearStart,
		IdMajor:          idm}
}

func (c Class) Save() error {
	if c.Id != -1 {
		return errors.New("[Class] Save: for saveing in the db the Id must be empty")
	}

	res, err := database.DB.Exec(`
		INSERT INTO classes
		(Year, Section, ScholarYearStart, IdM)
		VALUES
		(($1), ($2), ($3), ($4));`, c.Year, c.Section, c.ScholarYearStart, c.IdMajor)

	if err != nil {
		log.Panic(err.Error())
	}

	num, err := res.RowsAffected()
	if err != nil {
		log.Panic(err)
	}

	if num != 1 {
		log.Panicf("Wrong number of affected rows [%d]",
			num)
	}

	return nil
}

func (c Class) Update() error {
	if c.Id == -1 {
		return errors.New("[Student] Save: " +
			"for updateing in the db the Id must be set")
	}

	res, err := database.DB.Exec(`
		UPDATE Classes
		SET Year = ($2),
			Section = ($3),
			ScholarYearStart = ($4),
			IdM = ($5)
		WHERE Id = ($1);`, c.Id, c.Year, c.Section,
		c.ScholarYearStart, c.IdMajor)

	if err != nil {
		log.Panic(err.Error())
	}

	num, err := res.RowsAffected()
	if err != nil {
		log.Panic(err)
	}

	if num != 1 {
		log.Panicf("Wrong number of affected rows [%d]", num)
	}

	return nil
}

func (c Class) Delete() error {
	if c.Id == -1 {
		return errors.New("[Class] Delete: for deleteings in the db the Id must be set")
	}

	err := StudentClass{IdC: c.Id}.Delete()
	if err != nil {
		return err
	}

	res, err := database.DB.Exec(`
		DELETE FROM classes
		WHERE id = ($1);`, c.Id)
	if err != nil {
		return err
	}

	num, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if num != 1 {
		return fmt.Errorf("[Class] Save: wrong number of affected rows [%d]", num)
	}

	return nil
}

func GetClassById(id int) *Class {
	rows, err := database.DB.Query(`SELECT Id, Year, Section, ScholarYearStart, IdM
									FROM classes WHERE Id = $1;`, id)

	if err != nil {
		log.Panic(err.Error())
		return nil
	}
	defer rows.Close()
	if !rows.Next() {
		return nil
	}

	var result Class
	err = rows.Scan(&result.Id, &result.Year, &result.Section,
		&result.ScholarYearStart, &result.IdMajor)
	if err != nil {
		log.Panic(err.Error())
		return nil
	}

	return &result
}

type ClassView struct {
	Id               int
	Year             int
	Section          string
	ScholarYearStart int
	Major            string
	NumberStudents   int
}

func GetAllClasses() []*ClassView {
	rows, err := database.DB.Query(`SELECT C.Id, C.Year, C.Section, C.ScholarYearStart, M.name, COUNT(SC.Name) AS NumberStudents
									FROM classes AS C INNER JOIN
										 majors AS M ON C.IdM = M.Id LEFT JOIN
										 allactivestudentsclass AS SC ON C.Id = SC.IdC 
									GROUP BY C.ScholarYearStart, C.Year, C.Section, C.Id, M.name
									ORDER BY C.ScholarYearStart, C.Year, C.Section;`)

	if err != nil {
		log.Panic(err.Error())
		return nil
	}
	defer rows.Close()

	fmt.Println("Rows closed")

	var result []*ClassView
	for rows.Next() {
		var c ClassView
		err = rows.Scan(&c.Id, &c.Year, &c.Section,
			&c.ScholarYearStart, &c.Major, &c.NumberStudents)
		if err != nil {
			log.Panic(err.Error())
			return nil
		}

		result = append(result, &c)
	}

	return result
}
