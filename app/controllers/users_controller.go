package controllers

import (
	"net/http"

	"bitbucket.org/adhikag24/golang-api/app/usecase"
	"bitbucket.org/adhikag24/golang-api/app/utils"
	"github.com/gin-gonic/gin"
)

func DisplayUser(c *gin.Context) {
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

	user, apiErr := usecase.UsersUseCase.GetUser(userId)
	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}

	utils.Respond(c, http.StatusOK, user)
}

func DisplayAllUsers(c *gin.Context) {

	user, apiErr := usecase.UsersUseCase.GetAllUsers()

	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}

	utils.Respond(c, http.StatusOK, user)
}
