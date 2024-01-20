package models

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/SassoStorTo/studenti-italici/api/database"
)

type Student struct {
	Id          int
	Name        string
	LastName    string
	DateOfBirth time.Time
}

func NewStuent(name string, lastname string, dateofbirth time.Time) *Student {
	return &Student{
		Id:          -1,
		Name:        name,
		LastName:    lastname,
		DateOfBirth: dateofbirth}
}

func FormatTimeForDb(date time.Time) string {
	return fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d",
		date.Year(), date.Month(), date.Day(),
		date.Hour(), date.Minute(), date.Second())

}

func (s Student) Save() error {
	if s.Id != -1 {
		return errors.New("[Student] Save: for saveing in the db the Id must be empty")
	}

	res, err := database.DB.Exec(`
		INSERT INTO Students
		(Name, LastName, DateOfBirth)
		VALUES
		(($1), ($2), ($3));`, s.Name, s.LastName, FormatTimeForDb(s.DateOfBirth))

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

func (s Student) Update() error {
	if s.Id == -1 {
		return errors.New("[Student] Save: for updateing in the db the Id must be set")
	}

	res, err := database.DB.Exec(`
		UPDATE Students
		SET Name = ($2),
			LastName = ($3),
			DateOfBirth = ($4)
		WHERE Id=($1);`, s.Id, s.Name, s.LastName, FormatTimeForDb(s.DateOfBirth))

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
