package auth_test

import (
	"testing"

	"github.com/coregate/tickets-app/packages/auth"
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
		mockEncryptionSvc := new(fakes.IEncryptionsService)

		service := auth.NewAuthService(mockUserRepo, mockEncryptionSvc)

		// Act
		err := service.Register(newUser)

		// Assert
		assert.NoError(t, err)
	})
}
