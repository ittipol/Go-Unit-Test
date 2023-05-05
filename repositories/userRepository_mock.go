package repositories

import "github.com/stretchr/testify/mock"

type userRepositoryMock struct {
	mock.Mock
}

func NewUserRepositoryMock() *userRepositoryMock {
	return &userRepositoryMock{}
}

func (m *userRepositoryMock) Login(email string, password string) (*User, error) {
	args := m.Called(email)

	// Return mock data
	return args.Get(0).(*User), args.Error(1)
}
