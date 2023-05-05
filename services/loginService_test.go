package services_test

import (
	"errors"
	"gotest/repositories"
	"gotest/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	userRepo := repositories.NewUserRepositoryMock()

	// Arrange
	userRepo.On("Login", "email").Return(&repositories.User{
		ID:       1,
		Email:    "m1@email.com",
		Password: "$2a$12$ErV8S.A18n1nHftm5ph1Z./KhW/Kx41IVrliRSf1XuInN6TyPIAtC",
		Name:     "Mock 1",
	}, nil)

	userRepo.On("Login", "email2").Return(&repositories.User{}, errors.New("Not Found"))

	userRepo.On("Login", "email3").Return(&repositories.User{
		ID:       3,
		Email:    "m3@email.com",
		Password: "$2a$12$ErV8S.A18n1nHftm5ph1Z./KhW/Kx41IVrliRSf1XuInN6TyPIAtC",
		Name:     "Mock 3",
	}, nil)
	// ========================================================================

	// Act
	t.Run("Login success", func(t *testing.T) {

		s := services.NewLoginService(userRepo)

		res, err := s.Login("email", "1234")

		assert.Equal(t, nil, err)

		assert.Equal(t, "accessToken", res.AccessToken)
		assert.Equal(t, "refreshToken", res.RefreshToken)
	})

	t.Run("Login user not found", func(t *testing.T) {

		s := services.NewLoginService(userRepo)

		res, err := s.Login("email2", "1234")

		assert.Empty(t, res)
		assert.EqualError(t, err, "User Not Found")
	})

	t.Run("Login password not match", func(t *testing.T) {

		s := services.NewLoginService(userRepo)

		res, err := s.Login("email3", "pw")

		assert.Empty(t, res)
		assert.EqualError(t, err, "Cannot Login")
	})
}

func BenchmarkLogin(b *testing.B) {

	userRepo := repositories.NewUserRepositoryMock()

	userRepo.On("Login", "email").Return(&repositories.User{
		ID:       1,
		Email:    "m1@email.com",
		Password: "$2a$12$ErV8S.A18n1nHftm5ph1Z./KhW/Kx41IVrliRSf1XuInN6TyPIAtC",
		Name:     "Mock 1",
	}, nil)

	s := services.NewLoginService(userRepo)

	for i := 0; i < b.N; i++ {
		s.Login("email", "1234")
	}
}
