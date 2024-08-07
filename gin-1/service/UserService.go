package service

import (
	"gin-1/pojo"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var userList []pojo.User

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
	userList = append(userList, user)
	c.JSON(http.StatusOK, "User Added")
}

// DELETE User
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, user := range userList {
		if user.Id == id {
			userList = append(userList[:id], userList[id+1:]...)
			c.JSON(http.StatusOK, "User Deleted")
			return
		}
	}
	c.JSON(http.StatusNotFound, "User Not Found")
}

// PUT User
func PutUser(c *gin.Context) {
	beforeUser := pojo.User{}
	err := c.BindJSON(&beforeUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId, _ := strconv.Atoi(c.Param("id"))
	for i, user := range userList {
		if user.Id == userId {
			userList[i] = beforeUser
			c.JSON(http.StatusOK, "User Updated")
			return
		}
	}
	c.JSON(http.StatusNotFound, "User Not Found")
}
