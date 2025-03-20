package users

type IUserRepository interface {
	GetUserByUsernameOrEmail(username, email string) (*User, error)
	CreateUser(name, username, email, password string) (*User, error)
}
