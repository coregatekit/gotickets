package repositoriese

import (
	"github.com/coregate/tickets-app/database"
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

func (r *UsersRepository) CreateUser(data users.CreateUser) (*users.User, error) {
	user := &users.User{
		ID:       data.ID,
		Name:     data.Name,
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
	}

	result := r.dbConnection.Table(database.TableUsers).Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
