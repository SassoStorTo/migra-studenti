package models

import (
	"errors"
	"fmt"
	"log"

	"github.com/SassoStorTo/studenti-italici/pkg/database"
)

type User struct {
	Id            int
	Email         string
	Name          string
	Hd            string
	picture       string
	VerifiedEmail bool
	IsAdmin       bool
	IsEditor      bool
}

func NewUser(email string, name string, hd string, picture string, verifiedEmail bool) *User {
	return &User{
		Id:            -1,
		Email:         email,
		Name:          name,
		Hd:            hd,
		picture:       picture,
		VerifiedEmail: verifiedEmail}
}

func (u User) Save() error {
	if u.Id != -1 {
		return errors.New("[User] Save: for saveing in the db the Id must be empty")
	}

	res, err := database.DB.Exec(`
		INSERT INTO Users
		(Email, Name, Hd, VerifiedEmail)
		VALUES
		(($1), ($2), ($3), ($4));`, u.Email, u.Name, u.Hd, u.VerifiedEmail)
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

func (u User) Update() error {
	if u.Id == -1 {
		return errors.New("[User] Save: for updateing in the db the Id must be set")
	}

	res, err := database.DB.Exec(`
		UPDATE Users
		SET Name = ($2),
			Email = ($3),
			Hd = ($4),
			VerifiedEmail = ($5),
			IsAdmin = ($6),
			IsEditor = ($7)
		WHERE Id = ($1);`, u.Id, u.Name,
		u.Email, u.Hd, u.VerifiedEmail,
		u.IsAdmin, u.IsEditor)
	if err != nil {
		return err
	}

	num, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if num != 1 {
		return fmt.Errorf("[User] Update: wrong number of affected rows [%d]", num)
	}

	return nil
}

func (u User) Delete() error {
	if u.Id == -1 {
		return errors.New("[User] Delete: for deleteings in the db the Id must be set")
	}

	res, err := database.DB.Exec(`
		DELETE FROM users
		WHERE Id = ($1);`, u.Id)
	if err != nil {
		return err
	}

	num, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if num != 1 {
		return fmt.Errorf("[User] Delete: wrong number of affected rows [%d]", num)
	}

	return nil
}

func GetUserById(id int) (*User, error) {
	rows, err := database.DB.Query(`SELECT Id, Name, Email, VerifiedEmail, IsAdmin, IsEditor
									FROM Users WHERE Id=$1;`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, fmt.Errorf("id User non in db")
	}
	var result User
	err = rows.Scan(&result.Id, &result.Name, &result.Email,
		&result.VerifiedEmail, &result.IsAdmin, &result.IsEditor)
	if err != nil {
		log.Panic("rotto mentre lettura azzzz")
	}

	return &result, nil
}

func GetUserByEmail(email string) (*User, error) {
	rows, err := database.DB.Query(`SELECT Id, Name, Email, VerifiedEmail, IsAdmin, IsEditor
									FROM Users WHERE Email=$1;`, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, fmt.Errorf("email User non in db")
	}
	var result User
	err = rows.Scan(&result.Id, &result.Name, &result.Email,
		&result.VerifiedEmail, &result.IsAdmin, &result.IsEditor)
	if err != nil {
		log.Panic("rotto mentre lettura azzzz")
	}

	return &result, nil
}
