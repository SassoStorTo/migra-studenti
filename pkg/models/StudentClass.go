package models

import (
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
		(($1), ($2), ($3));`, s.IdS, s.IdC, FormatTimeForDb(s.CreationDate))

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
