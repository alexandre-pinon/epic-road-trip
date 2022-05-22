package service

import (
	"errors"
	"net/http"

	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/alexandre-pinon/epic-road-trip/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userService struct {
	userRepository repository.UserRepository
}

type UserService interface {
	GetAllUsers() (*[]model.User, error)
	GetUserByID(id primitive.ObjectID) (*model.User, error)
	CreateUser(user *model.User) error
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (svc *userService) GetAllUsers() (*[]model.User, error) {
	return svc.userRepository.GetAllUsers()
}

func (svc *userService) GetUserByID(id primitive.ObjectID) (*model.User, error) {
	user, _ := svc.userRepository.GetUserByID(id)
	if user == nil {
		return nil, &model.AppError{
			StatusCode: http.StatusNotFound,
			Err:        errors.New("user not found"),
		}
	}

	return user, nil
}

func (svc *userService) CreateUser(user *model.User) error {
	if user == nil {
		return &model.AppError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("user is nil pointer"),
		}
	}

	if _, err := svc.userRepository.CreateUser(user); err != nil {
		return &model.AppError{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return nil
}
