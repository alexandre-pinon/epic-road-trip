package service

import (
	"errors"
	"net/http"

	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/alexandre-pinon/epic-road-trip/repository"
)

type userService struct {
	userRepository repository.UserRepository
}

type UserService interface {
	CreateUser(user *model.User) error
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (service *userService) CreateUser(user *model.User) error {
	return &model.AppError{
		Err:        errors.New("TODO: implement CreateUser"),
		StatusCode: http.StatusNotImplemented,
	}
}
