package controller

import (
	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/gin-gonic/gin"
)

type rootController struct {
}

type RootController interface {
	Ok(ctx *gin.Context) (*model.AppResult, error)
}

func NewRootController() RootController {
	return &rootController{}
}

func (controller *rootController) Ok(ctx *gin.Context) (*model.AppResult, error) {
	return &model.AppResult{
		Data:       struct{}{},
		Message:    "Ok",
		StatusCode: 200,
	}, nil
}
