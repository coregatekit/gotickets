package auth

import (
	"github.com/coregate/tickets-app/packages/users"
)

type AuthService struct {
	usersRepository users.IUserRepository
}

type IAuthService interface {
	Register(data CreateUser) error
}

func NewAuthService(userRepo users.IUserRepository) *AuthService {
	return &AuthService{
		usersRepository: userRepo,
	}
}
