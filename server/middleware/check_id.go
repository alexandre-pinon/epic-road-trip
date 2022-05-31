package middleware

import (
	"errors"
	"net/http"

	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CheckID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, model.AppResponse{
				Success:   false,
				Message:   errors.New("invalid id").Error(),
				Data:      struct{}{},
				ValErrors: []model.ValError{},
			})
		}
		ctx.Set("id", id)
		ctx.Next()
	}
}
