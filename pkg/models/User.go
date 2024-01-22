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
	VerifiedEmail bool
	IsAdmin       bool
	IsEditor      bool
}

func NewUser(email string, name string, hd string, verifiedEmail bool) *User {
	return &User{
		Id:            -1,
		Email:         email,
		Name:          name,
		Hd:            hd,
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
