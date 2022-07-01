package main

import (
	"fmt"
	"mypackage/database"
	"mypackage/server"
	// "mypackage/router"
	// "github.com/gin-gonic/gin"
	// "gorm-test/controllers"
	// "net/http"
	// "mypackage/variable"
)

// var Db *gorm.DB

func main() {
	database.ConnectDB()
	database.DbMigrate()
	fmt.Println("Server serve at")
	server.Serve()
	// userRepo := database.InitDb()
	// r.POST("/users", router.userRepo.Test)
	// var persona variable.User
	// _ := persona
}
