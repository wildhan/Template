package usecase

import (
	"template/package/user/model"
	"template/package/user/repository"
)

type userUsecase struct {
	repo repository.UserRepo
}

func NewUserUsecase(repo repository.UserRepo) UserUsecase {
	return &userUsecase{repo}
}

type UserUsecase interface {
	GetUsers() ([]model.User, error)
	AddUser(model.User) error
	EditUser(model.User) error
}

func (uc *userUsecase) GetUsers() ([]model.User, error) {
	return uc.repo.GetUsers()
}

func (uc *userUsecase) AddUser(user model.User) error {
	return uc.repo.AddUser(user)
}

func (uc *userUsecase) EditUser(user model.User) error {
	return uc.repo.EditUser(user)
}
