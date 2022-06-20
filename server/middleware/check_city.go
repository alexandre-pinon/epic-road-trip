package middleware

import (
	"errors"
	"net/http"

	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/gin-gonic/gin"
)

func CheckCity() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		city , err := ctx.Params.Get("city")
		if err == false {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, model.AppResponse{
				Success:   false,
				Message:   errors.New("no city fund").Error(),
				Data:      struct{}{},
				ValErrors: []model.ValError{},
			})
		}
		ctx.Set("city", city)
		ctx.Next()
	}
}