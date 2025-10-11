package domain

type UserRepository interface {
	InsertUser(u *User) (*User, error)
	GetAllUsers() ([]User, error)
	GetUser(id int64) (*User, error)
	UpdateUser(u *User) (*User, error)
	DeleteUser(id int64) (int64, error)
}

type UserService interface {
	CreateUser(u *User) (*User, error)
	GetAllUsers() ([]User, error)
	GetUser(id int64) (*User, error)
	UpdateUser(u *User) (*User, error)
	DeleteUser(id int64) (int64, error)
}
