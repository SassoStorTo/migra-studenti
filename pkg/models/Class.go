package models

import (
	"errors"
	"log"

	"github.com/SassoStorTo/studenti-italici/api/database"
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
