package repositoriese

import (
	"github.com/coregate/tickets-app/packages/users"
	"gorm.io/gorm"
)

type UsersRepository struct {
	dbConnection *gorm.DB
}

func NewUsersRepository(dbConnection *gorm.DB) *UsersRepository {
	return &UsersRepository{
		dbConnection: dbConnection,
	}
}

func (r *UsersRepository) CreateUser(name, username, email, password string) (*users.User, error) {
	user := &users.User{
		Name:     name,
		Username: username,
		Email:    email,
		Password: password,
	}

	result := r.dbConnection.Table("auth.users").Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
