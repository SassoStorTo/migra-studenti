package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/SassoStorTo/studenti-italici/pkg/database"
)

type User struct {
	Id            int    `json:"id"`
	Email         string `json:"email"`
	Name          string `json:"name"`
	Hd            string `json:"hd"`
	Picture       string `json:"picture"`
	VerifiedEmail bool   `json:"verified_email"`
	IsAdmin       bool   `json:"is_admin"`
	IsEditor      bool   `json:"is_editor"`
}

func NewUser(email string, name string, hd string, picture string, verifiedEmail bool) *User {
	return &User{
		Id:            -1,
		Email:         email,
		Name:          name,
		Hd:            hd,
		Picture:       picture,
		VerifiedEmail: verifiedEmail}
}

func (u User) Save() error {
	if u.Id != -1 {
		return errors.New("[User] Save: for saveing in the db the Id must be empty")
	}

	res, err := database.DB.Exec(`
		INSERT INTO Users
		(Email, Name, Hd, VerifiedEmail, IsAdmin, IsEditor, Picture)
		VALUES
		(($1), ($2), ($3), ($4), ($5), ($6), ($7));`, u.Email, u.Name, u.Hd, u.VerifiedEmail, u.IsAdmin, u.IsEditor, u.Picture)
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
			IsEditor = ($7),
			Picture = ($8)
		WHERE Id = ($1);`, u.Id, u.Name,
		u.Email, u.Hd, u.VerifiedEmail,
		u.IsAdmin, u.IsEditor, u.Picture)
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

func (i User) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}

func GetUserById(id int) (*User, error) {
	rows, err := database.DB.Query(`SELECT Id, Name, Email, VerifiedEmail, IsAdmin, IsEditor, Picture
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
		&result.VerifiedEmail, &result.IsAdmin, &result.IsEditor,
		&result.Picture)
	if err != nil {
		log.Panic("rotto mentre lettura azzzz")
	}
	return &result, nil
}

func GetUserByEmail(email string) (*User, error) {
	rows, err := database.DB.Query(`SELECT Id, Name, Email, VerifiedEmail, IsAdmin, IsEditor, Picture
									FROM Users WHERE Email=($1);`, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, fmt.Errorf("email User non in db")
	}
	var result User
	err = rows.Scan(&result.Id, &result.Name, &result.Email,
		&result.VerifiedEmail, &result.IsAdmin, &result.IsEditor,
		&result.Picture)
	if err != nil {
		log.Panic("rotto mentre lettura azzzz")
	}

	return &result, nil
}

func GetAllUsers() ([]User, error) {
	rows, err := database.DB.Query(`SELECT Id, Name, Email, VerifiedEmail, IsAdmin, IsEditor, Picture
									FROM Users;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []User
	for rows.Next() {
		var u User
		err = rows.Scan(&u.Id, &u.Name, &u.Email,
			&u.VerifiedEmail, &u.IsAdmin, &u.IsEditor,
			&u.Picture)
		if err != nil {
			log.Panic("rotto mentre lettura azzzz")
		}
		result = append(result, u)
	}

	return result, nil
}
