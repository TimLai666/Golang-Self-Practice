package pojo

import "gin-1/database"

type User struct {
	Id       int    `json:"UserId"`
	Name     string `json:"UserName"`
	Password string `json:"UserPassword"`
	Email    string `json:"UserEmail"`
}

func FindAllUsers() []User {
	var users []User
	database.DBconnect.Find(&users)
	return users
}

func FindUserById(id int) User {
	var user User
	database.DBconnect.Where("id = ?", id).First(&user)
	return user
}
