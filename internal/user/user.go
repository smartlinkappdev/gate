package user

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
}

func (u *User) CreateUser(user *User) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *User) UpdateUser(user *User) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *User) DeleteUser(user *User) error {
	//TODO implement me
	panic("implement me")
}

func (u *User) GetUser(user *User) (*User, error) {
	//TODO implement me
	panic("implement me")
}
