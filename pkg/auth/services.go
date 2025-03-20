package auth

import (
	"github.com/coregate/tickets-app/pkg/encryptions"
	"github.com/coregate/tickets-app/pkg/users"
	"github.com/pkg/errors"
)

type authService struct {
	usersRepository    users.IUserRepository
	encryptionsService encryptions.IEncryptionsService
}

type IAuthService interface {
	Register(data CreateUser) error
}

func NewAuthService(userRepo users.IUserRepository, encryptionsService encryptions.IEncryptionsService) IAuthService {
	return &authService{
		usersRepository:    userRepo,
		encryptionsService: encryptionsService,
	}
}

func (s *authService) Register(data CreateUser) error {
	existingUser, err := s.usersRepository.GetUserByUsernameOrEmail(data.Username, data.Email)
	if err != nil {
		return errors.Wrap(err, "failed to get user by username or email")
	}

	if existingUser != nil {
		return errors.New("username or email already exists")
	}

	return nil
}
