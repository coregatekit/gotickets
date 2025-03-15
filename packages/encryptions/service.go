package encryptions

type encryptionsService struct{}

type IEncryptionsService interface {
	HashPassword(password string) (string, error)
}

func NewEncryptionsService() IEncryptionsService {
	return &encryptionsService{}
}

func (s *encryptionsService) HashPassword(password string) (string, error) {
	return "", nil
}
