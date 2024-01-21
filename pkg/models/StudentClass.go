package models

import (
	"errors"
	"log"
	"time"

	"github.com/SassoStorTo/studenti-italici/api/database"
)

type StudentClass struct {
	IdS          int
	IdC          int
	CreationDate time.Time
}

func NewStudentClass(ids int, idc int, date time.Time) *StudentClass {
	return &StudentClass{ids, idc, date}
}

func (s StudentClass) Save() error {
	res, err := database.DB.Exec(`
		INSERT INTO StudentClass
		(IdS, IdC, CreationDate)
		VALUES
		(($1), ($2), ($3));`, s.IdS, s.IdC, database.FormatTimeForDb(s.CreationDate))

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

func (s StudentClass) Delete() error {
	if s.IdS <= 0 && s.IdC <= 0 {
		return errors.New("[StudentClass] Delete: for deleteings in the db at " +
			"least one Id must be set")
	}

	if s.IdS <= 0 {
		_, err := database.DB.Exec(`
		DELETE FROM studentclass
		WHERE IdC = ($1);`, s.IdC)
		return err
	}
	if s.IdC <= 0 {
		_, err := database.DB.Exec(`
		DELETE FROM studentclass
		WHERE IdS = ($1);`, s.IdS)
		return err
	}

	_, err := database.DB.Exec(`
	DELETE FROM studentclass
	WHERE IdC = ($1) AND IdS = ($2);`, s.IdC, s.IdS)
	return err
}
