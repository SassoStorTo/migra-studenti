package models

import (
	"errors"
	"fmt"

	"github.com/SassoStorTo/studenti-italici/pkg/database"
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

func (m Majors) Delete() error {
	if m.Id == -1 {
		return errors.New("[Major] Delete: for deleteings in the db the Id must be set")
	}

	// deletes all the associated classes
	sos, err := m.getAssociatedClasses()
	if err != nil {
		return nil
	}
	for _, t := range *sos {
		t.Delete()
	}

	// delete major
	res, err := database.DB.Exec(`
		DELETE FROM majors
		WHERE id = ($1);`, m.Id)
	if err != nil {
		return err
	}

	// safety check
	num, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if num != 1 {
		return fmt.Errorf("[Class] Save: wrong number of affected rows [%d]", num)
	}
	return nil
}

func (m Majors) Update() error {
	if m.Id == -1 {
		return errors.New("[Major] Save: for updateing in the db the Id must be set")
	}

	res, err := database.DB.Exec(`
		UPDATE Majors
		SET Name = ($2)
		WHERE Id=($1);`, m.Id, m.Name)

	if err != nil {
		return err
	}

	num, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if num != 1 {
		return fmt.Errorf("[Student] Update: wrong number of affected rows [%d]", num)
	}

	return nil

}

func (m Majors) getAssociatedClasses() (*[]Class, error) {
	rows, err := database.DB.Query(`
		SELECT (Id) FROM classes 
		WHERE IdM = ($1);`, m.Id)

	if err != nil {
		return nil, err
	}

	var data []Class
	for rows.Next() {
		var result Class
		err := rows.Scan(&result.Id)
		if err != nil {
			return nil, err
		}
		data = append(data, result)
	}

	return &data, nil
}
