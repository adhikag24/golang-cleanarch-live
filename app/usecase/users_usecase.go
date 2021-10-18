package usecase

import (
	"bitbucket.org/adhikag24/golang-api/app/domain"
	"bitbucket.org/adhikag24/golang-api/app/repository"
	"bitbucket.org/adhikag24/golang-api/app/utils"
)

type UserUseCase struct {
	rpo repository.Repo
}

func (u *UserUseCase) GetUser(userId string) (*domain.User, *utils.ApplicationError) {
	user, err := u.rpo.GetUser(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserUseCase) GetAllUsers() ([]*domain.User, *utils.ApplicationError) {

	user, err := u.rpo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return user, nil
}
