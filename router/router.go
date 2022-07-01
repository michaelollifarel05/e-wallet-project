package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"mypackage/variable"
	"mypackage/database"
)


var ToDatabase = database.ConnectDB()

func Register(c *gin.Context) {
	var input variable.User

	if err := c.ShouldBindJSON(&input); 
	err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var returnCheck = ValidateRegister(input)
	if returnCheck == 0 {
		ToDatabase.Save(&input)
		c.JSON(http.StatusOK, gin.H{"return":"register success", "status": "success", "code": 200})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"return":"Failed to register","status": "User already registered","code": 400})		
	}
}

func SearchUser(c *gin.Context) {
	var input variable.User
	if err := ToDatabase.Where("name = ?", c.Param("name")).First(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, input)	
	
}

func ShowAllUsers(c *gin.Context) {
	var input []variable.User
	ToDatabase.Find(&input)

	c.JSON(http.StatusOK, gin.H{"data": input})
}

func ValidateRegister(person variable.User) (check int) {
	ToDatabase.Where("name = ?", person.Name).Find(&person)
	if person.ID == 0 {
		check = 0
	} else {
		check = 1
	}
	return
}
