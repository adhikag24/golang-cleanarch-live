package repository

import (
	"bitbucket.org/adhikag24/golang-api/app/domain"
	"bitbucket.org/adhikag24/golang-api/app/utils"
)

type Repo interface {
	GetAllUsers() ([]*domain.User, *utils.ApplicationError)
	GetUser(userID string) (*domain.User, *utils.ApplicationError)
}

func NewRepo() Repo {
	return &Repository{}
}
