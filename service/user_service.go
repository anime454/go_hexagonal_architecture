package service

import "github.com/anime454/go_hexagonal_architecture/repository"

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) userService {
	return userService{userRepo: repo}
}

func (us userService) Register(u User) (string, error) {
	user := repository.User{
		Id:           u.Id,
		Username:     u.Username,
		Password:     u.Password,
		FullName:     u.FullName,
		Email:        u.Email,
		Role:         u.Role,
		AutoDatetime: u.AutoDatetime,
	}
	userId, err := us.userRepo.Create(user)
	if err != nil {
		return "", err
	}
	return userId, nil
}

func (us userService) GetAllUsers() ([]UserDetail, error) {
	allUsersRepo, err := us.userRepo.GetAll()
	if err != nil {
		return nil, err
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

func (us userService) GetUserById(string) (*UserDetail, error) { return nil, nil }

func (us userService) UpdateUser(UserDetail) (string, error) { return "", nil }

func (us userService) DeleteUser(string) (string, error) { return "", nil }
