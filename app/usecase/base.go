package usecase

import (
	"bitbucket.org/adhikag24/golang-api/app/domain"
	"bitbucket.org/adhikag24/golang-api/app/repository"
	"bitbucket.org/adhikag24/golang-api/app/utils"
)

type UC interface {
	//user
	GetUser(userId string) (*domain.User, *utils.ApplicationError)
	//section
	GetAllUsers() ([]*domain.User, *utils.ApplicationError)
}

func NewUC(repo repository.Repo) UC {
	return &UserUseCase{rpo: repo}
}
