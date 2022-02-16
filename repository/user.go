package repository

import "time"

type User struct {
	Id           string    `db:"id"`
	Username     string    `db:"username"`
	Password     string    `db:"password"`
	FullName     string    `db:"fullname"`
	Email        string    `db:"email"`
	Role         string    `db:"role"`
	AutoDatetime time.Time `db:"auto_datetime"`
}

// type User struct {
// 	Id       string `json:"id"`
// 	Username string `json:"username"`
// 	FullName string `json:"fullname"`
// 	Email    string `json:"email"`
// 	Role     string `json:"role"`
// }

type UserRepository interface {
	Create(User) (string, error)
	GetAll() ([]User, error)
	GetById(string) (*User, error)
	Update(User) (*User, error)
	Delete(string) error
}
