package user

type InterfaceUser interface {
	CreateUser(user *User) (*User, error)
	UpdateUser(user *User) (*User, error)
	DeleteUser(user *User) error
	GetUser(user *User) (*User, error)
}
