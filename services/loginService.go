package services

import (
	"errors"
	"gotest/models"
	"gotest/repositories"

	"golang.org/x/crypto/bcrypt"
)

var (
	errNotFound    = errors.New("User Not Found")
	errCannotLogin = errors.New("Cannot Login")
)

type LoginService struct {
	userRepo repositories.UserRepository
}

func NewLoginService(userRepo repositories.UserRepository) *LoginService {
	return &LoginService{userRepo}
}

func (obj *LoginService) Login(email string, password string) (models.LoginResponse, error) {

	user, err := obj.userRepo.Login(email, password)

	if err != nil {
		return models.LoginResponse{}, errNotFound
	}

	// check user password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return models.LoginResponse{}, errCannotLogin
	}

	return models.LoginResponse{
		AccessToken:  "accessToken",
		RefreshToken: "refreshToken",
	}, nil
}
