package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/SassoStorTo/migra-studenti/pkg/database"
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
	if student := GetStudentById(s.IdS); student == nil {
		return fmt.Errorf("the student does not exist")
	}
	if class := GetClassById(s.IdC); class == nil {
		return fmt.Errorf("the class does not exist")
	}

	// rows, err := database.DB.Query("SELECT * FROM studentclass WHERE IdS = $1 AND IdC = $2;", s.IdS, s.IdC)
	// if err != nil {
	// 	return err
	// }
	// defer rows.Close()
	// if rows.Next() {
	// 	return fmt.Errorf("the link already exist")
	// }

	res, err := database.DB.Exec(`
		INSERT INTO StudentClass
		(IdS, IdC, CreationDate)
		VALUES
		(($1), ($2), ($3));`, s.IdS, s.IdC, database.FormatTimeForDb(s.CreationDate))

	if err != nil {
		return err
	}

	num, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if num != 1 {
		return fmt.Errorf("wrong number of affected rows [%d]", num)
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

func GeltLastId() int {
	var id int
	row := database.DB.QueryRow("SELECT MAX(Id) FROM studentclass;")
	row.Scan(&id)
	return id
}
