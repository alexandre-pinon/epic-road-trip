package controller

import (
	"fmt"
	"net/http"

	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/alexandre-pinon/epic-road-trip/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userController struct {
	userService service.UserService
}

type UserController interface {
	GetAllUsers(ctx *gin.Context) (*model.AppResult, *model.AppError)
	GetUserByID(ctx *gin.Context) (*model.AppResult, *model.AppError)
	CreateUser(ctx *gin.Context) (*model.AppResult, *model.AppError)
	UpdateUser(ctx *gin.Context) (*model.AppResult, *model.AppError)
	DeleteUser(ctx *gin.Context) (*model.AppResult, *model.AppError)
}

func NewUserController(svc service.UserService) UserController {
	return &userController{svc}
}

func (ctrl *userController) GetAllUsers(ctx *gin.Context) (*model.AppResult, *model.AppError) {
	users, err := ctrl.userService.GetAllUsers()
	if err != nil {
		return nil, &model.AppError{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	var data *[]model.User
	if len(*users) != 0 {
		data = users
	} else {
		data = &[]model.User{}
	}

	return &model.AppResult{
		StatusCode: http.StatusOK,
		Message:    "Users retrieved successfully",
		Data:       data,
	}, nil
}

func (ctrl *userController) GetUserByID(ctx *gin.Context) (*model.AppResult, *model.AppError) {
	id, _ := ctx.Get("id")
	user, err := ctrl.userService.GetUserByID(id.(primitive.ObjectID))
	if err != nil {
		return nil, err.(*model.AppError)
	}

	return &model.AppResult{
		StatusCode: http.StatusOK,
		Message:    fmt.Sprintf("User %s retrieved successfully", id.(primitive.ObjectID).Hex()),
		Data:       user,
	}, nil
}

func (ctrl *userController) CreateUser(ctx *gin.Context) (*model.AppResult, *model.AppError) {
	var userFormData model.UserFormData

	if err := ctx.ShouldBindJSON(&userFormData); err != nil {
		return nil, &model.AppError{
			StatusCode: http.StatusBadRequest,
			Err:        err,
		}
	}

	if err := ctrl.userService.HashPassword(&userFormData); err != nil {
		return nil, &model.AppError{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	if err := ctrl.userService.CreateUser(&userFormData.User); err != nil {
		return nil, err.(*model.AppError)
	}

	return &model.AppResult{
		StatusCode: http.StatusCreated,
		Message:    "User created successfully",
		Data:       struct{}{},
	}, nil
}

func (ctrl *userController) UpdateUser(ctx *gin.Context) (*model.AppResult, *model.AppError) {
	var user model.User

	id, _ := ctx.Get("id")
	if err := ctx.ShouldBindJSON(&user); err != nil {
		return nil, &model.AppError{
			StatusCode: http.StatusBadRequest,
			Err:        err,
		}
	}

	if err := ctrl.userService.UpdateUser(id.(primitive.ObjectID), &user); err != nil {
		return nil, err.(*model.AppError)
	}

	return &model.AppResult{
		StatusCode: http.StatusOK,
		Message:    "User updated successfully",
		Data:       struct{}{},
	}, nil
}

func (ctrl *userController) DeleteUser(ctx *gin.Context) (*model.AppResult, *model.AppError) {
	id, _ := ctx.Get("id")
	if err := ctrl.userService.DeleteUser(id.(primitive.ObjectID)); err != nil {
		return nil, err.(*model.AppError)
	}

	return &model.AppResult{
		StatusCode: http.StatusOK,
		Message:    "User deleted successfully",
		Data:       struct{}{},
	}, nil
}
