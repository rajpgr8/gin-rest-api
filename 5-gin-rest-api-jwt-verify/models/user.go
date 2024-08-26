package models

import (
	"gin-rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

var users = []User{}

func (u User) Save() error {
	// later: add it to a database
	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	u.Password = hashedPassword
	users = append(users, u)

	return err
}

func (u User) ValidateCredentials() error {

	// fetch password for the u.Email from the database
	//query := "SELECT id, password FROM users WHERE email = ?"
	//row := db.DB.QueryRow(query, u.Email)
	//
	//var retrievedPassword string
	//err := row.Scan(&u.ID, &retrievedPassword)
	//
	//if err != nil {
	//	return errors.New("Credentials invalid")
	//}
	//
	//passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	//if !passwordIsValid {
	//	return errors.New("Credentials invalid")
	//}
	//
	return nil
}
