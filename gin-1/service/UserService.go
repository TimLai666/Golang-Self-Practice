package service

import (
	"gin-1/pojo"
	"net/http"
	"strconv"

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

// DELETE User
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, user := range userList {

		userList = append(userList[:id], userList[id+1:]...)
		c.JSON(http.StatusOK, "User Deleted")
		return

	}
	c.JSON(http.StatusNotFound, "User Not Found")
}
