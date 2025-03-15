package encryptions_test

import (
	"crypto/rand"
	"errors"
	"testing"

	"github.com/coregate/tickets-app/packages/encryptions"
	"github.com/stretchr/testify/assert"
)

type errorReader struct{}

func (e *errorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("mocked error")
}

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
		assert.Equal(t, 97, len(hashedPassword))

	})

	t.Run("should return error when cannot generate random bytes", func(t *testing.T) {
		// Arrange
		password := "password"
		crypto := encryptions.NewEncryptionsService()

		originalReader := rand.Reader
		defer func() {
			rand.Reader = originalReader
		}()
		rand.Reader = &errorReader{}

		// Act
		hashedPassword, err := crypto.HashPassword(password)

		// Assert
		assert.Error(t, err)
		assert.Empty(t, hashedPassword)
	})
}
