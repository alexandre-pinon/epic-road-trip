package controller

import (
	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/gin-gonic/gin"
)

type rootController struct {
}

type RootController interface {
	Healthcheck(ctx *gin.Context) (*model.AppResult, *model.AppError)
}

func NewRootController() RootController {
	return &rootController{}
}

// Healthcheck godoc
// @Summary healthcheck
// @Description allows healthcheck
// @Tags healthcheck
// @Accept json
// @Produce json
// @Success 200 {object} model.AppResponse
// @Router / [get]
func (controller *rootController) Healthcheck(ctx *gin.Context) (*model.AppResult, *model.AppError) {
	return &model.AppResult{
		StatusCode: 200,
		Message:    "Ok",
		Data:       struct{}{},
	}, nil
}
