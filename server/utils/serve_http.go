package utils

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
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
			var me mongo.WriteException

			log.Printf("%T: %v", err.Err, err.Err)

			switch {
			case errors.As(err.Err, &ve):
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

			case errors.As(err.Err, &me):
				errArr := make([]model.ValError, len(me.WriteErrors))

				for i, we := range me.WriteErrors {
					if mongo.IsDuplicateKeyError(we) {
						sMsg := strings.Split(we.Message, "dup key: { ")
						msg := strings.Split(sMsg[len(sMsg)-1], ":")[0]
						field := strings.ToTitle(msg[:1]) + msg[1:]
						message := fmt.Sprintf("%s is already taken", field)
						errArr[i] = model.ValError{Field: field, Message: message}
					} else {
						errArr[i] = model.ValError{Field: fmt.Sprint(we.Code), Message: we.Message}
					}
				}

				ctx.JSON(http.StatusBadRequest, model.AppResponse{
					Success:   false,
					Message:   "invalid json request body",
					Data:      data,
					ValErrors: errArr,
				})

			default:
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
