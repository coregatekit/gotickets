package encryptions_test

import (
	"crypto/rand"
	"errors"
	"testing"

	"github.com/coregate/tickets-app/configs"
	"github.com/coregate/tickets-app/packages/encryptions"
	"github.com/stretchr/testify/assert"
)

type errorReader struct{}

func (e *errorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("mocked error")
}

func TestHashPassword(t *testing.T) {
	configs := &configs.Configs{
		Argon: &configs.ArgonParams{
			Memory:      64 * 1024,
			Iterations:  3,
			Parallelism: 2,
			SaltLength:  16,
			KeyLength:   32,
		},
	}

	t.Run("should hash a password successfully", func(t *testing.T) {
		// Arrange
		password := "password"
		crypto := encryptions.NewEncryptionsService(configs)

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
		crypto := encryptions.NewEncryptionsService(configs)

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

// func TestComparePassword(t *testing.T) {
// 	configs := &configs.Configs{
// 		Argon: &configs.ArgonParams{
// 			Memory:      64 * 1024,
// 			Iterations:  3,
// 			Parallelism: 2,
// 			SaltLength:  16,
// 			KeyLength:   32,
// 		},
// 	}

// 	t.Run("should compare password successfully", func(t *testing.T) {
// 		// Arrange
// 		password := "password"
// 		crypto := encryptions.NewEncryptionsService(configs)
// 		hashedPassword, _ := crypto.HashPassword(password)

// 		// Act
// 		match, err := crypto.ComparePassword(hashedPassword, password)

// 		// Assert
// 		assert.NoError(t, err)
// 		assert.True(t, match)
// 	})

// 	t.Run("should return false when password does not match", func(t *testing.T) {
// 		// Arrange
// 		password := "password"
// 		crypto := encryptions.NewEncryptionsService(configs)
// 		hashedPassword, _ := crypto.HashPassword(password)

// 		// Act
// 		match, err := crypto.ComparePassword(hashedPassword, "wrong-password")

// 		// Assert
// 		assert.NoError(t, err)
// 		assert.False(t, match)
// 	})
// }
