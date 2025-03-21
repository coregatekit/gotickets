package repos

import (
	"github.com/coregate/tickets-app/database"
	"github.com/coregate/tickets-app/pkg/users"
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

func (r *UsersRepository) GetUserByUsernameOrEmail(username, email string) (*users.User, error) {
	user := users.User{}

	query := r.dbConnection.Table(database.TableUsers)

	if username != "" {
		query = query.Where("username = ?", username)
	}
	if email != "" {
		query = query.Or("email = ?", email)
	}

	result := query.First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}

	return &user, nil
}

func (r *UsersRepository) CreateUser(name, username, email, password string) (*users.User, error) {
	user := &users.User{
		Name:     name,
		Username: username,
		Email:    email,
		Password: password,
	}

	result := r.dbConnection.Table(database.TableUsers).Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
