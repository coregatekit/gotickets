package auth_test

import (
	"errors"
	"testing"

	"github.com/coregate/tickets-app/pkg/auth"
	"github.com/coregate/tickets-app/pkg/users"
	"github.com/coregate/tickets-app/tests/fakes"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	t.Run("should register a new user successfully", func(t *testing.T) {
		// Arrange
		newUser := auth.CreateUser{
			Email:    "aerichan@coregate.dev",
			Name:     "Uchinaga Aeri",
			Username: "aerichandesu",
			Password: "password",
		}
		mockUserRepo := new(fakes.IUserRepository)
		mockUserRepo.On("GetUserByUsernameOrEmail", newUser.Username, newUser.Email).Return(nil, nil)

		mockEncryptionSvc := new(fakes.IEncryptionsService)

		service := auth.NewAuthService(mockUserRepo, mockEncryptionSvc)

		// Act
		err := service.Register(newUser)

		// Assert
		assert.NoError(t, err)
		mockUserRepo.AssertCalled(t, "GetUserByUsernameOrEmail", newUser.Username, newUser.Email)
	})

	t.Run("should return error when an error occured while getting user by username or email", func(t *testing.T) {
		// Arrange
		newUser := auth.CreateUser{
			Email:    "aerichan@coregate.dev",
			Name:     "Uchinaga Aeri",
			Username: "aerichandesu",
			Password: "password",
		}
		mockUserRepo := new(fakes.IUserRepository)
		mockUserRepo.On("GetUserByUsernameOrEmail", newUser.Username, newUser.Email).Return(nil, errors.New("some error"))

		mockEncryptionSvc := new(fakes.IEncryptionsService)

		service := auth.NewAuthService(mockUserRepo, mockEncryptionSvc)

		// Act
		err := service.Register(newUser)

		// Assert
		assert.Error(t, err)
		mockUserRepo.AssertCalled(t, "GetUserByUsernameOrEmail", newUser.Username, newUser.Email)
	})

	t.Run("should rerturn error when username or email already exists", func(t *testing.T) {
		// Arrange
		newUser := auth.CreateUser{
			Email:    "aerichan@coregate.dev",
			Name:     "Uchinaga Aeri",
			Username: "aerichandesu",
			Password: "password",
		}
		mockUserRepo := new(fakes.IUserRepository)
		mockUserRepo.On("GetUserByUsernameOrEmail", newUser.Username, newUser.Email).Return(&users.User{
			Email:    "aerichan@coregate.dev",
			Name:     "Aeri",
			Username: "aerichandesu",
		}, nil)

		mockEncryptionSvc := new(fakes.IEncryptionsService)

		service := auth.NewAuthService(mockUserRepo, mockEncryptionSvc)

		// Act
		err := service.Register(newUser)

		// Assert
		assert.Error(t, err)
		mockUserRepo.AssertCalled(t, "GetUserByUsernameOrEmail", newUser.Username, newUser.Email)

	})
}
