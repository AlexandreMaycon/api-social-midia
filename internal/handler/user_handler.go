package handler

import "github.com/gin-gonic/gin"

func CreateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func GetAllUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func GetUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func UpdateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func DeleteUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
