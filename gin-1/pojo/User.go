package pojo

import "gin-1/database"

type User struct {
	Id       int    `json:"UserId" binding:"omitempty"`
	Name     string `json:"UserName" binding:"required,gt=5"`
	Password string `json:"UserPassword" binding:"min=4,max=20,userpassword"`
	Email    string `json:"UserEmail" binding:"email"`
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
