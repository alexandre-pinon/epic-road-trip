package utils

import (
	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/gin-gonic/gin"
)

type controllerFunc func(ctx *gin.Context) (*model.AppResult, *model.AppError)

func ServeHTTP(f controllerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		result, err := f(ctx)

		var data interface{}
		if result != nil {
			data = result.Data
		} else {
			data = struct{}{}
		}

		if err != nil {
			ctx.JSON(err.StatusCode, model.Response{
				Success: false,
				Message: err.Error(),
				Data:    data,
			})
			return
		}

		ctx.JSON(result.StatusCode, model.Response{
			Success: true,
			Message: result.Message,
			Data:    data,
		})
	}
}
