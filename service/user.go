package service

import "time"

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type UserDetail struct {
	Id           string    `json:"id"`
	Username     string    `json:"username"`
	FullName     string    `json:"fullname"`
	Email        string    `json:"email"`
	Role         string    `json:"role"`
	AutoDatetime time.Time `json:"auto_datetime"`
}

type UserService interface {
	Register(User) (string, error)
	GetAllUsers() ([]UserDetail, error)
	GetUserById(string) (*UserDetail, error)
	UpdateUser(UserDetail) (string, error)
	DeleteUser(string) error
}
