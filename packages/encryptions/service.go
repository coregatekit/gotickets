package encryptions

import (
	"crypto/rand"

	"golang.org/x/crypto/argon2"
)

type encryptionsService struct{}

type IEncryptionsService interface {
	HashPassword(password string) ([]byte, error)
}

func NewEncryptionsService() IEncryptionsService {
	return &encryptionsService{}
}

func (s *encryptionsService) HashPassword(password string) ([]byte, error) {
	p := Params{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}

	salt, err := s.generateRandomBytes(p.SaltLength)
	if err != nil {
		return nil, err
	}

	hash := argon2.IDKey([]byte(password), salt, p.Iterations, p.Memory, p.Parallelism, p.KeyLength)
	return hash, nil
}

func (s *encryptionsService) generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
