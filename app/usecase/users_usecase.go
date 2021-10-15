package usecase

import (
	"bitbucket.org/adhikag24/golang-api/app/domain"
	"bitbucket.org/adhikag24/golang-api/app/repository"
	"bitbucket.org/adhikag24/golang-api/app/utils"
)

type usersUseCase struct{}

var (
	UsersUseCase usersUseCase
)

func (u *usersUseCase) GetUser(userId string) (*domain.User, *utils.ApplicationError) {
	user, err := repository.Repository.GetUser(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *usersUseCase) GetAllUsers() ([]*domain.User, *utils.ApplicationError) {
	user, err := repository.Repository.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return user, nil
}
