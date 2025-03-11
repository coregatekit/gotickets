package repositoriese

import "gorm.io/gorm"

type usersRepository struct {
	dbConnection *gorm.DB
}

func NewUsersRepository(dbConnection *gorm.DB) *usersRepository {
	return &usersRepository{
		dbConnection: dbConnection,
	}
}
