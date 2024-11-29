package employee

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET
func GetEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Employee GET Method!",
	})
}

// GET
func GetEmployeebyID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"message": "Employee GET Method!",
	})
}

// POST
func PostEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Employee post Method!",
	})
}

// PUT
func PutEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Employee put Method!",
	})
}

// DELETE
func Deletemployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Employee DELETEE Method!",
	})
}
