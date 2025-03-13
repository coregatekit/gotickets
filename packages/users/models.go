package users

import (
	"github.com/coregate/tickets-app/database"
	"github.com/nrednav/cuid2"
	"gorm.io/gorm"
)

type User struct {
	ID        string `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Username  string `json:"username" gorm:"unique"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime"`
}

func (u *User) TableName() string {
	return database.TableUsers
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = cuid2.Generate()
	return nil
}
