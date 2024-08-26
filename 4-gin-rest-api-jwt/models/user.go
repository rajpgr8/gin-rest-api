package models

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

var users = []User{}

func (e User) Save() {
	// later: add it to a database
	users = append(users, e)
}

func GetAllUsers() []User {
	return users
}
