package controllers

import (
	"bitbucket.org/adhikag24/golang-api/app/usecase"
	"github.com/gin-gonic/gin"
)

type Controllers interface {
	DisplayUser(c *gin.Context)
	DisplayAllUsers(c *gin.Context)
}

func NewControllers(us usecase.UC) Controllers {
	return &UserController{UC: us}
}
