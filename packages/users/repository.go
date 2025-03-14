package users

type IUserRepository interface {
	CreateUser(name, username, email, password string) (*User, error)
}
