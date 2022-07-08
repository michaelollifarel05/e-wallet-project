package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"mypackage/variable"
	"mypackage/database"
	"fmt"
	"strconv"

)


var ToDatabase = database.ConnectDB()

func Register(c *gin.Context) {
	var input variable.User
	// var account variable.Account

	fmt.Println(input.ID)
	if err := c.ShouldBindJSON(&input);
	err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	// var returnCheck = ValidateRegister(input)
	var returnCheck = 0
	if returnCheck == 0 {
		ToDatabase.Create(&input)
		// c.JSON(http.StatusOK, gin.H{"return":"register success", "status": "success", "code": 200})
		c.JSON(http.StatusOK, gin.H{"return": input})
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

func TopUp(c *gin.Context) {
	var input variable.Account
	ToDatabase.Select("balance").Where("user_id = ?", c.Param("id")).Find(&input)
	time, _ := strconv.Atoi(c.Param("amount"))
	var newBalance = input.Balance + time	
	ToDatabase.Model(&input).Where("user_id = ?", c.Param("id")).Update("balance",newBalance) 
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
