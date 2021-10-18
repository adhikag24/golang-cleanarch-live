package repository

import (
	"bitbucket.org/adhikag24/golang-api/app/domain"
	"bitbucket.org/adhikag24/golang-api/app/utils"
)

// import (
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

type Repo interface {
	GetAllUsers() ([]*domain.User, *utils.ApplicationError)
	GetUser(userID string) (*domain.User, *utils.ApplicationError)
}

func NewRepo() Repo {
	return &Repository{}
}
