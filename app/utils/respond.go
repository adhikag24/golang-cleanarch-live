package utils

import "github.com/gin-gonic/gin"

func Respond(c *gin.Context, status int, body interface{}) {
	if c.GetHeader("Accept") == "application/json" {
		c.JSON(status, body)
		return
	}

	c.JSON(status, body)
}

func RespondError(c *gin.Context, err *ApplicationError) {
	if c.GetHeader("Accept") == "application/json" {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(err.StatusCode, err)
}
