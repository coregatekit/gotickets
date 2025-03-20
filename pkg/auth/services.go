package auth

import (
	"github.com/coregate/tickets-app/pkg/encrypt"
	"github.com/coregate/tickets-app/pkg/users"
	"github.com/pkg/errors"
)

type authService struct {
	usersRepository    users.IUserRepository
	encryptionsService encrypt.IEncryptionsService
}

type IAuthService interface {
	Register(data CreateUser) error
}

func NewAuthService(userRepo users.IUserRepository, encryptionsService encrypt.IEncryptionsService) IAuthService {
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

	hashedPassword, err := s.encryptionsService.HashPassword(data.Password)
	if err != nil {
		return errors.Wrap(err, "failed to hash password")
	}

	user, err := s.usersRepository.CreateUser(data.Name, data.Username, data.Email, hashedPassword)
	if err != nil {
		return errors.Wrap(err, "failed to create user")
	}

	_ = user

	return nil
}
