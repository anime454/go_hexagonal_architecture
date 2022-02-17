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

type UserDetail struct {
	Id           string    `db:"id"`
	Username     string    `db:"username"`
	FullName     string    `db:"fullname"`
	Email        string    `db:"email"`
	Role         string    `db:"role"`
	AutoDatetime time.Time `db:"auto_datetime"`
}
type UserRepository interface {
	Create(User) (string, error)
	GetAll() ([]User, error)
	GetById(string) (*User, error)
	Update(UserDetail) (string, error)
	Delete(string) error
}
