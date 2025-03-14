package auth

import (
	"github.com/coregate/tickets-app/packages/users"
)

type authService struct {
	usersRepository users.IUserRepository
}

type IAuthService interface {
	Register(data CreateUser) error
}

func NewAuthService(userRepo users.IUserRepository) IAuthService {
	return &authService{
		usersRepository: userRepo,
	}
}

func (s *authService) Register(data CreateUser) error {
	return nil
}
