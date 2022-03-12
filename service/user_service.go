package service

import (
	"database/sql"
	"time"

	"github.com/anime454/go_hexagonal_architecture/logs"
	"github.com/anime454/go_hexagonal_architecture/repository"
	"github.com/anime454/go_hexagonal_architecture/responses"
	"github.com/google/uuid"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) userService {
	return userService{userRepo: repo}
}

func (us userService) Register(u User) (string, error) {
	user := repository.User{
		Id:           uuid.NewString(),
		Username:     u.Username,
		Password:     u.Password,
		FullName:     u.FullName,
		Email:        u.Email,
		Role:         u.Role,
		AutoDatetime: time.Now(),
	}
	userId, err := us.userRepo.Create(user)

	if err != nil {
		logs.Error(err)
		return "", responses.InternalServerError()
	}
	return userId, nil
}

func (us userService) GetAllUsers() ([]UserDetail, error) {
	allUsersRepo, err := us.userRepo.GetAll()
	if err.Error() == sql.ErrNoRows.Error() {
		logs.Info(dataNotFound)
		return nil, responses.DataNotFound()
	}

	if err != nil {
		logs.Error(err)
		return nil, responses.InternalServerError()
	}

	users := []UserDetail{}
	for _, userRepo := range allUsersRepo {
		tmp := UserDetail{
			Id:           userRepo.Id,
			Username:     userRepo.Username,
			FullName:     userRepo.FullName,
			Email:        userRepo.Email,
			Role:         userRepo.Role,
			AutoDatetime: userRepo.AutoDatetime,
		}
		users = append(users, tmp)
	}

	return users, nil
}

func (us userService) GetUserById(id string) (*UserDetail, error) {
	u, err := us.userRepo.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			logs.Info(dataNotFound)
			return nil, responses.DataNotFound()
		}
		logs.Error(err)
		return nil, responses.InternalServerError()
	}
	user := UserDetail{
		Id:           u.Id,
		Username:     u.Password,
		FullName:     u.FullName,
		Email:        u.Email,
		Role:         u.Role,
		AutoDatetime: u.AutoDatetime,
	}
	return &user, nil
}

func (us userService) UpdateUser(u UserDetail) (string, error) {
	uRepo := repository.UserDetail{
		Id:           u.Id,
		Username:     u.Username,
		FullName:     u.FullName,
		Email:        u.Email,
		Role:         u.Role,
		AutoDatetime: time.Now(),
	}
	userId, err := us.userRepo.Update(uRepo)
	if err != nil {
		if err == sql.ErrNoRows {
			logs.Info(dataNotFound)
			return "", responses.DataNotFound()
		}
		logs.Error(err)
		return "", responses.InternalServerError()
	}

	return userId, nil
}

func (us userService) DeleteUser(id string) error {
	err := us.userRepo.Delete(id)
	if err != nil {
		if err == sql.ErrNoRows {
			logs.Info(dataNotFound)
			return responses.DataNotFound()
		}
		logs.Error(err)
		return responses.InternalServerError()
	}
	return nil
}
