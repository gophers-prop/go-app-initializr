package main

import (
	"appname/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
	
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		user := new(controllers.UserController)
		v1.POST("/user", user.Create)
		v1.GET("/user/:id", user.Get)
		v1.GET("/user", user.Find)
		v1.PUT("/user/:id", user.Update)
		v1.DELETE("/user/:id", user.Delete)
	}

	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not Found")
	})

	router.Run()
}
