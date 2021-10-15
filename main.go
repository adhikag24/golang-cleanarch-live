// package main

// import (
// 	"github.com/gin-gonic/gin"
// )

// type UserIDRequest struct {
// 	UserID string `json:"userid" bson:"userid"`
// }

// func HomePage(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"message": "Hello World",
// 	})
// }

// func PostHomePage(c *gin.Context) {
// 	requestBody := UserIDRequest{}
// 	c.Bind(&requestBody)

// 	c.JSON(200, gin.H{
// 		"message": string(requestBody.UserID),
// 	})
// }

// func QueryStrings(c *gin.Context) {

// }

// func main() {

// 	r := gin.Default()
// 	r.GET("/", HomePage)
// 	r.POST("/", PostHomePage)
// 	r.GET("/query", QueryStrings)

// 	r.Run()
// }

package main

import (
	"bitbucket.org/adhikag24/golang-api/app"
)

func main() {
	app.StartApp()
}
