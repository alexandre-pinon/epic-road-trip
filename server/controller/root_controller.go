package controller

import (
	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/gin-gonic/gin"
)

type rootController struct {
}

type RootController interface {
	Ok(ctx *gin.Context) (*model.AppResult, *model.AppError)
}

func NewRootController() RootController {
	return &rootController{}
}

func (controller *rootController) Ok(ctx *gin.Context) (*model.AppResult, *model.AppError) {
	return &model.AppResult{
		StatusCode: 200,
		Message:    "Ok",
		Data:       struct{}{},
	}, nil
}
