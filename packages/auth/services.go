package auth

import (
	"github.com/coregate/tickets-app/packages/encryptions"
	"github.com/coregate/tickets-app/packages/users"
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
	return nil
}
