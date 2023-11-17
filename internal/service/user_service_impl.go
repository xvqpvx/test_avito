package service

import (
	"context"
	"github.com/rs/zerolog/log"
	"test_avito/internal/data/request"
	"test_avito/internal/data/response"
	"test_avito/internal/helper"
	"test_avito/internal/model"
	"test_avito/internal/repos"
	"time"
)

type UserServiceImpl struct {
	UserRepository repos.UserRepo
}

func NewUserServiceImpl(userRepository repos.UserRepo) UserService {
	return &UserServiceImpl{UserRepository: userRepository}
}

func (u *UserServiceImpl) FindAll(ctx context.Context) []response.UserResponse {
	users := u.UserRepository.FindAll(ctx)

	var usersResp []response.UserResponse

	for _, value := range users {
		user := response.UserResponse{
			IdUser:    value.IdUser,
			Firstname: value.Firstname,
			Lastname:  value.Lastname,
		}
		usersResp = append(usersResp, user)
	}

	return usersResp
}

func (u *UserServiceImpl) FindById(ctx context.Context, idUser int) response.UserResponse {
	user, err := u.UserRepository.FindById(ctx, idUser)
	helper.PanicIfError(err)
	return response.UserResponse{
		IdUser:    user.IdUser,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
	}
}

func (u *UserServiceImpl) Save(ctx context.Context, request request.UserCreateRequest) {
	user := model.User{
		Firstname: request.Firstname,
		Lastname:  request.Lastname,
		Created:   time.Now().Format("2006-01-02 15:04:05"),
		Updated:   time.Now().Format("2006-01-02 15:04:05"),
	}
	u.UserRepository.Save(ctx, user)
	log.Info().Msg("New user created")
}

func (u *UserServiceImpl) Update(ctx context.Context, request request.UserUpdateRequest) {
	user, err := u.UserRepository.FindById(ctx, request.IdUser)
	helper.PanicIfError(err)

	user.Firstname, user.Lastname = request.Firstname, request.Lastname
	user.Updated = time.Now().Format("2006-01-02 15:04:05")
	u.UserRepository.Update(ctx, user)
	log.Info().Msg("User updated")
}

func (u *UserServiceImpl) Delete(ctx context.Context, idUser int) {
	user, err := u.UserRepository.FindById(ctx, idUser)
	helper.PanicIfError(err)

	u.UserRepository.Delete(ctx, user.IdUser)
	log.Info().Msg("User deleted")
}
