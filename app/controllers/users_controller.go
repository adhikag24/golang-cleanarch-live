package controllers

import (
	"net/http"

	"bitbucket.org/adhikag24/golang-api/app/usecase"
	"bitbucket.org/adhikag24/golang-api/app/utils"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UC usecase.UC
}

func (u *UserController) DisplayUser(c *gin.Context) {
	userId := c.PostForm("userid")
	if userId == "" {
		apiErr := &utils.ApplicationError{
			Message:    "userid is required",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	user, apiErr := u.UC.GetUser(userId)
	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}

	utils.Respond(c, http.StatusOK, user)
}

func (u *UserController) DisplayAllUsers(c *gin.Context) {
	user, apiErr := u.UC.GetAllUsers()

	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}

	utils.Respond(c, http.StatusOK, user)
}
