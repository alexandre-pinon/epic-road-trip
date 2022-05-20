package controller

import "github.com/gin-gonic/gin"

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type AppResult struct {
	Data       interface{}
	Message    string
	StatusCode int
}

type rootController struct {
}

type RootController interface {
	Ok(ctx *gin.Context) (*AppResult, error)
}

func NewRootController() RootController {
	return &rootController{}
}

func (controller *rootController) Ok(ctx *gin.Context) (*AppResult, error) {
	return &AppResult{
		Data:       struct{}{},
		Message:    "Ok",
		StatusCode: 200,
	}, nil
}
