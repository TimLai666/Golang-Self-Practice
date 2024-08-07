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

func FindUserById(id string) User {
	var user User
	database.DBconnect.Where("id = ?", id).First(&user)
	return user
}

func AddUser(user User) User {
	database.DBconnect.Create(&user)
	return user
}

func DeleteUser(id string) bool {
	user := User{}
	result := database.DBconnect.Where("id = ?", id).Delete(&user).RowsAffected
	return result != 0
}

func UpdateUser(userId string, user User) User {
	database.DBconnect.Model(&user).Where("id = ?", userId).Updates(&user)
	return user
}
