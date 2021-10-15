package routes

import (
	"bitbucket.org/adhikag24/golang-api/app/controllers"
	"github.com/gin-gonic/gin"
)

type routes struct{}

var (
	Routes routes
)

func (r *routes) Router(router *gin.Engine) {
	router.POST("/user", controllers.DisplayUser)
	router.GET("/users", controllers.DisplayAllUsers)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "123123123"})
	})

}
