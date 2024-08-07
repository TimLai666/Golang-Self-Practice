package service

import (
	"gin-1/pojo"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET User
func FindAllUsers(c *gin.Context) {
	// c.JSON(http.StatusOK, userList)
	users := pojo.FindAllUsers()
	c.JSON(http.StatusOK, users)
}

// GET User by ID
func FindUserById(c *gin.Context) {
	user := pojo.FindUserById(c.Param("id"))
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "User Not Found")
		return
	}
	log.Println(user)
	c.JSON(http.StatusOK, user)
}

// POST User
func PostUser(c *gin.Context) {
	user := pojo.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUser := pojo.AddUser(user)
	c.JSON(http.StatusOK, newUser)
}

// DELETE User
func DeleteUser(c *gin.Context) {
	userId := c.Param("id")
	result := pojo.DeleteUser(userId)
	if !result {
		c.JSON(http.StatusNotFound, "User Not Found")
		return
	}
	c.JSON(http.StatusOK, "User Deleted")
}

// PUT User
func PutUser(c *gin.Context) {
	userId := c.Param("id")
	user := pojo.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUser := pojo.UpdateUser(userId, user)
	if newUser.Id == 0 {
		c.JSON(http.StatusNotFound, "User Not Found")
		return
	}
	c.JSON(http.StatusOK, newUser)
}
