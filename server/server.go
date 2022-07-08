package server


import (
	"github.com/gin-gonic/gin"
	"net/http"
	"mypackage/variable"
	"mypackage/router"
	"fmt"
)

func Serve(){
	var port = variable.GoDotEnvVariable("SERVER_PORT")
	port = ":"+port
	fmt.Println(port)
	r := setupRouter()
	_ = r.Run(port)
}

func setupRouter() *gin.Engine {
	r := gin.Default()
 
	r.GET("ping", func(c *gin.Context) {
	   c.JSON(http.StatusOK, "pong")
	})
	r.POST("/register", router.Register)
	r.GET("/search-user/:name", router.SearchUser)
	r.GET("/show-all-users", router.ShowAllUsers)
	r.POST("/top-up/:id/:amount", router.TopUp)
	return r
 }

