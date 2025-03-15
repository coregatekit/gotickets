package encryptions

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

type encryptionsService struct{}

type IEncryptionsService interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword, password string) (bool, error)
}

func NewEncryptionsService() IEncryptionsService {
	return &encryptionsService{}
}

func (s *encryptionsService) HashPassword(password string) (string, error) {
	p := Params{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}

	salt, err := s.generateRandomBytes(p.SaltLength)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, p.Iterations, p.Memory, p.Parallelism, p.KeyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, p.Memory, p.Iterations, p.Parallelism, b64Salt, b64Hash)
	return encodedHash, nil
}

func (s *encryptionsService) ComparePassword(hashedPassword, password string) (bool, error) {
	return true, nil
}

func (s *encryptionsService) generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
