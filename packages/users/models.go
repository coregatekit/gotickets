package users

import (
	"github.com/coregate/tickets-app/common"
)

type User struct {
	common.BaseModel
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
