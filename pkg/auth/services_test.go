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
		mockUserRepo.On("CreateUser", newUser.Name, newUser.Username, newUser.Email, "hashed_password").
			Return(&users.User{
				Email:    newUser.Email,
				Name:     newUser.Name,
				Username: newUser.Username,
				Password: "hashed_password",
			}, nil)

		mockEncryptionSvc := new(fakes.IEncryptionsService)
		mockEncryptionSvc.On("HashPassword", newUser.Password).Return("hashed_password", nil)

		service := auth.NewAuthService(mockUserRepo, mockEncryptionSvc)

		// Act
		err := service.Register(newUser)

		// Assert
		assert.NoError(t, err)
		mockUserRepo.AssertCalled(t, "GetUserByUsernameOrEmail", newUser.Username, newUser.Email)
		mockEncryptionSvc.AssertCalled(t, "HashPassword", newUser.Password)
		mockUserRepo.AssertCalled(t, "CreateUser", newUser.Name, newUser.Username, newUser.Email, "hashed_password")
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

	t.Run("should return error when failed to hash password", func(t *testing.T) {
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
		mockEncryptionSvc.On("HashPassword", newUser.Password).Return("", errors.New("some error"))

		service := auth.NewAuthService(mockUserRepo, mockEncryptionSvc)

		// Act
		err := service.Register(newUser)

		// Assert
		assert.Error(t, err)
		mockUserRepo.AssertCalled(t, "GetUserByUsernameOrEmail", newUser.Username, newUser.Email)
		mockEncryptionSvc.AssertCalled(t, "HashPassword", newUser.Password)
	})

	t.Run("should return error when failed to create user", func(t *testing.T) {
		// Arrange
		newUser := auth.CreateUser{
			Email:    "aerichan@coregate.dev",
			Name:     "Uchinaga Aeri",
			Username: "aerichandesu",
			Password: "password",
		}

		mockUserRepo := new(fakes.IUserRepository)
		mockUserRepo.On("GetUserByUsernameOrEmail", newUser.Username, newUser.Email).Return(nil, nil)
		mockUserRepo.On("CreateUser", newUser.Name, newUser.Username, newUser.Email, "hashed_password").Return(nil, errors.New("some error"))

		mockEncryptionSvc := new(fakes.IEncryptionsService)
		mockEncryptionSvc.On("HashPassword", newUser.Password).Return("hashed_password", nil)

		service := auth.NewAuthService(mockUserRepo, mockEncryptionSvc)

		// Act
		err := service.Register(newUser)

		// Assert
		assert.Error(t, err)
		mockUserRepo.AssertCalled(t, "GetUserByUsernameOrEmail", newUser.Username, newUser.Email)
		mockEncryptionSvc.AssertCalled(t, "HashPassword", newUser.Password)
		mockUserRepo.AssertCalled(t, "CreateUser", newUser.Name, newUser.Username, newUser.Email, "hashed_password")
	})
}

func TestLogin(t *testing.T) {
	t.Run("should login user successfully", func(t *testing.T) {
		// Arrange
		loginRequest := auth.LoginRequest{
			Username: "aerichandesu",
			Password: "password",
		}

		mockUserRepo := new(fakes.IUserRepository)
		mockUserRepo.On("GetUserByUsernameOrEmail", loginRequest.Username, "").
			Return(&users.User{
				Email:    "aerichan@coregate.dev",
				Name:     "Uchinaga Aeri",
				Username: "aerichandesu",
				Password: "hashed_password",
			}, nil)

		mockEncryptionSvc := new(fakes.IEncryptionsService)
		mockEncryptionSvc.On("ComparePassword", loginRequest.Password, "hashed_password").Return(true, nil)

		service := auth.NewAuthService(mockUserRepo, mockEncryptionSvc)

		// Act
		result, err := service.Login(loginRequest)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.NotEmpty(t, result.Token)
		assert.NotEmpty(t, result.RefreshToken)
	})
}
