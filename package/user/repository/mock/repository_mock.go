package mock

import (
	"template/package/user/model"

	"github.com/stretchr/testify/mock"
)

type userRepoMock struct {
	mock.Mock
}

func (_m *userRepoMock) GetUsers() ([]model.User, error) {
	return nil, nil
}

func (_m *userRepoMock) AddUser(model.User) error {
	return nil
}

func (_m *userRepoMock) EditUser(model.User) error {
	return nil
}
