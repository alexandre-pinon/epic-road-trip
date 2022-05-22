package utils

import (
	"errors"

	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
			var ve validator.ValidationErrors
			if errors.As(err.Err, &ve) {
				errArr := make([]model.ValError, len(ve))

				for i, fe := range ve {
					errArr[i] = model.ValError{Field: fe.Field(), Message: model.GetValErrorMsg(fe)}
				}

				ctx.JSON(err.StatusCode, model.AppResponse{
					Success:   false,
					Message:   "invalid json request body",
					Data:      data,
					ValErrors: errArr,
				})
			} else {
				ctx.JSON(err.StatusCode, model.AppResponse{
					Success:   false,
					Message:   err.Error(),
					Data:      data,
					ValErrors: []model.ValError{},
				})
			}
			return
		}

		ctx.JSON(result.StatusCode, model.AppResponse{
			Success:   true,
			Message:   result.Message,
			Data:      data,
			ValErrors: []model.ValError{},
		})
	}
}
