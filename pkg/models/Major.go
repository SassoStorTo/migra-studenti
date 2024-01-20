package models

import (
	"errors"
	"log"

	"github.com/SassoStorTo/studenti-italici/api/database"
)

type Majors struct { //Todo: change name to Major
	Id   int
	Name string
}

func NewMajor(name string) Majors {
	return Majors{Id: -1, Name: name}
}

func (m Majors) Save() error {
	if m.Id != -1 {
		return errors.New("[Major] Save: for saveing in the db the Id must be empty")
	}

	res, err := database.DB.Exec(`
		INSERT INTO majors
		(name)
		VALUES
		($1);`, m.Name)

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
