package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

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

// Get all users godoc
// @Summary Get all users
// @Description Get all users
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} model.GetAllUserSuccess "Success"
// @Failure 500 {object} model.InternalServerError "Internal server error"
// @Router /user [get]
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

// Get user by ID godoc
// @Summary Get user by ID
// @Description Get user given a valid ID
// @Tags User
// @Accept json
// @Produce json
// @Param populate query bool false "Populate the user's roadtrips or not"
// @Param id path string true "User ID"
// @Success 200 {object} model.GetUserByIDSuccess "Success"
// @Failure 400 {object} model.InvalidID "Invalid ID"
// @Failure 404 {object} model.UserNotFound "User not found"
// @Failure 500 {object} model.InternalServerError "Internal server error"
// @Router /user/:id [get]
func (ctrl *userController) GetUserByID(ctx *gin.Context) (*model.AppResult, *model.AppError) {
	id, _ := ctx.Get("id")

	var populate bool
	populateParam, exists := ctx.GetQuery("populate")
	if exists {
		var err error
		populate, err = strconv.ParseBool(populateParam)
		if err != nil {
			return nil, &model.AppError{
				StatusCode: http.StatusBadRequest,
				Err:        errors.New("invalid query parameters"),
			}
		}
	}

	user, err := ctrl.userService.GetUserByID(id.(primitive.ObjectID), populate)
	if err != nil {
		return nil, err.(*model.AppError)
	}

	return &model.AppResult{
		StatusCode: http.StatusOK,
		Message:    fmt.Sprintf("User %s retrieved successfully", id.(primitive.ObjectID).Hex()),
		Data:       user,
	}, nil
}

// Create user godoc
// @Summary Create user
// @Description Create user user given valid firstname, lastname, email (unique), password, phone (unique)
// @Tags User
// @Accept json
// @Produce json
// @Param user body model.RegisterRequest true "firstname, lastname, email, password, phone"
// @Success 200 {object} model.RegisterSuccess "Success"
// @Failure 400 {object} model.InvalidJsonBody "Invalid request body"
// @Failure 401 {object} model.Unauthorized "Missing/Expired token"
// @Failure 500 {object} model.InternalServerError "Internal server error"
// @Router /user [post]
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

// Update user godoc
// @Summary Update user
// @Description Update user user given valid ID, firstname, lastname, email (unique), phone (unique)
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body model.UpdateUserRequest true "firstname, lastname, email, phone"
// @Success 200 {object} model.UpdateUserSuccess "Success"
// @Failure 400 {object} model.InvalidJsonBody "Invalid ID/body"
// @Failure 401 {object} model.Unauthorized "Missing/Expired token"
// @Failure 404 {object} model.UserNotFound "User not found"
// @Failure 500 {object} model.InternalServerError "Internal server error"
// @Router /user/:id [put]
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

// Delete user godoc
// @Summary Delete user
// @Description Delete user user given valid ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} model.DeleteUserSuccess "Success"
// @Failure 400 {object} model.InvalidID "Invalid ID"
// @Failure 401 {object} model.Unauthorized "Missing/Expired token"
// @Failure 404 {object} model.UserNotFound "User not found"
// @Failure 500 {object} model.InternalServerError "Internal server error"
// @Router /user/:id [delete]
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
