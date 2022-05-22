package controller

import (
	"net/http"

	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/alexandre-pinon/epic-road-trip/service"
	"github.com/gin-gonic/gin"
)

type userController struct {
	userService service.UserService
}

type UserController interface {
	GetAllUsers(ctx *gin.Context) (*model.AppResult, *model.AppError)
	CreateUser(ctx *gin.Context) (*model.AppResult, *model.AppError)
}

func NewUserController(svc service.UserService) UserController {
	return &userController{svc}
}

func (ctrl *userController) GetAllUsers(ctx *gin.Context) (*model.AppResult, *model.AppError) {
	users, err := ctrl.userService.GetAllUsers()
	if err != nil {
		return nil, &model.AppError{
			Err:        err,
			StatusCode: http.StatusInternalServerError,
		}
	}

	var data *[]model.User
	if len(*users) != 0 {
		data = users
	} else {
		data = &[]model.User{}
	}

	return &model.AppResult{
		Data:       data,
		Message:    "Users retrieved successfully",
		StatusCode: http.StatusOK,
	}, nil
}

func (ctrl *userController) CreateUser(ctx *gin.Context) (*model.AppResult, *model.AppError) {
	var user model.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		return nil, &model.AppError{
			Err:        err,
			StatusCode: http.StatusBadRequest,
		}
	}

	if err := ctrl.userService.CreateUser(&user); err != nil {
		return nil, &model.AppError{
			Err:        err,
			StatusCode: err.(*model.AppError).StatusCode,
		}
	}

	return &model.AppResult{
		Message:    "User created successfully",
		StatusCode: http.StatusCreated,
	}, nil
}
