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

func (s Student) Save() error {
	if s.Id != -1 {
		return errors.New("[Student] Save: for saveing in the db the Id must be empty")
	}

	res, err := database.DB.Exec(`
		INSERT INTO Students
		(Name, LastName, DateOfBirth)
		VALUES
		(($1), ($2), ($3));`, s.Name, s.LastName, database.FormatTimeForDb(s.DateOfBirth))

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
		WHERE Id=($1);`, s.Id, s.Name, s.LastName, database.FormatTimeForDb(s.DateOfBirth))

	if err != nil {
		return err
	}

	num, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if num != 1 {
		return fmt.Errorf("[Student] Save: wrong number of affected rows [%d]", num)
	}

	return nil
}

func (s Student) Delete() error {
	if s.Id == -1 {
		return errors.New("[Student] Save: for deleteings in the db the Id must be set")
	}

	err := StudentClass{IdS: s.Id}.Delete()
	if err != nil {
		return err
	}

	res, err := database.DB.Exec(`
		DELETE FROM students
		WHERE Id = ($1);`, s.Id)
	if err != nil {
		return err
	}

	num, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if num != 1 {
		return fmt.Errorf("[Student] Save: wrong number of affected rows [%d]", num)
	}

	return nil
}
