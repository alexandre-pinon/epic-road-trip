package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/gin-gonic/gin"
)

func CheckTravelMode() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		mode := model.Transport(strings.ToUpper(ctx.Param("mode")))
		if mode != model.Air && mode != model.Ground {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, model.AppResponse{
				Success:   false,
				Message:   errors.New("invalid travel mode").Error(),
				Data:      struct{}{},
				ValErrors: []model.ValError{},
			})
		}
		ctx.Set("mode", mode)
		ctx.Next()
	}
}
