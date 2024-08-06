package service

import (
	"gin-1/pojo"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userList []pojo.User

// GET User
func FindAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, userList)
}

// POST User
func PostUser(c *gin.Context) {
	user := pojo.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userList = append(userList, user)
	c.JSON(http.StatusOK, "User Added")
}
