package encryptions

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/coregate/tickets-app/configs"
	"golang.org/x/crypto/argon2"
)

type encryptionsService struct {
	configs *configs.Configs
}

type IEncryptionsService interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword, password string) (bool, error)
}

var (
	ErrorInvalidHash               = errors.New("the encoded hash is not in the correct format")
	ErrorIncompatibleVersion       = errors.New("incompatible version of argon2")
	ErrorMismatchedHashAndPassword = errors.New("the hash and password do not match")
)

func NewEncryptionsService(configs *configs.Configs) IEncryptionsService {
	return &encryptionsService{
		configs: configs,
	}
}

func (s *encryptionsService) HashPassword(password string) (string, error) {
	p := s.configs.Argon

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

func (s *encryptionsService) decodeHash(encodedHash string) (*configs.ArgonParams, []byte, []byte, error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, ErrorInvalidHash
	}

	var version int
	_, err := fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}

	p := &configs.ArgonParams{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.Memory, &p.Iterations, &p.Parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	sal, err := base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.SaltLength = uint32(len(sal))

	hash, err := base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.KeyLength = uint32(len(hash))

	return p, sal, hash, nil
}
