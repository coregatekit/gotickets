package encryptions_test

import (
	"testing"

	"github.com/coregate/tickets-app/packages/encryptions"
	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	t.Run("should hash a password successfully", func(t *testing.T) {
		// Arrange
		password := "password"
		crypto := encryptions.NewEncryptionsService()

		// Act
		hashedPassword, err := crypto.HashPassword(password)

		// Assert
		assert.NoError(t, err)
		assert.NotEmpty(t, hashedPassword)
	})
}
