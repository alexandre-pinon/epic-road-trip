package utils

import (
	"net/http"

	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/gin-gonic/gin"
)

type controllerFunc func(ctx *gin.Context) (*model.AppResult, *model.AppError)

func ServeHTTP(f controllerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		result, err := f(ctx)

		if result == nil {
			ctx.JSON(http.StatusInternalServerError, model.Response{
				Success: false,
				Message: http.StatusText(http.StatusInternalServerError),
				Data:    nil,
			})
			return
		}

		if err != nil {
			ctx.JSON(err.StatusCode, model.Response{
				Success: false,
				Message: err.Error(),
				Data:    result.Data,
			})
			return
		}

		ctx.JSON(result.StatusCode, model.Response{
			Success: true,
			Message: result.Message,
			Data:    result.Data,
		})
	}
}
