package service

import (
	"errors"
	"net/http"

	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/alexandre-pinon/epic-road-trip/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepository repository.UserRepository
}

type UserService interface {
	GetAllUsers() (*[]model.User, error)
	GetUserByID(id primitive.ObjectID) (*model.User, error)
	CreateUser(user *model.User) error
	HashPassword(user *model.UserFormData) error
	UpdateUser(id primitive.ObjectID, user *model.User) error
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

func (svc *userService) HashPassword(user *model.UserFormData) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		return err
	}

	user.HashedPassword = string(hashed)
	return nil
}

func (svc *userService) UpdateUser(id primitive.ObjectID, user *model.User) error {
	if user == nil {
		return &model.AppError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("user is nil pointer"),
		}
	}

	updatedResult, err := svc.userRepository.UpdateUser(id, user)
	if err != nil {
		return &model.AppError{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	if updatedResult.MatchedCount == 0 {
		return &model.AppError{
			StatusCode: http.StatusNotFound,
			Err:        errors.New("user not found"),
		}
	}

	return nil
}
