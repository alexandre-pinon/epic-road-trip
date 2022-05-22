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
	GetAllUsers() (*[]model.User, error)
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (svc *userService) CreateUser(user *model.User) error {
	if user == nil {
		return &model.AppError{
			Err:        errors.New("user is nil pointer"),
			StatusCode: http.StatusInternalServerError,
		}
	}

	if err := svc.userRepository.CreateUser(user); err != nil {
		return &model.AppError{
			Err:        err,
			StatusCode: http.StatusInternalServerError,
		}
	}

	return nil
}

func (svc *userService) GetAllUsers() (*[]model.User, error) {
	return &[]model.User{}, &model.AppError{
		Err:        errors.New("TODO: implement GetAllUsers"),
		StatusCode: http.StatusInternalServerError,
	}
}
