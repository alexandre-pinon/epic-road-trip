package app

import (
	"net/http"

	"github.com/alexandre-pinon/epic-road-trip/controller"
	"github.com/gin-gonic/gin"
)

type appController func(ctx *gin.Context) (*controller.AppResult, error)

func serveHTTP(c appController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		result, err := c(ctx)

		if result == nil {
			ctx.JSON(http.StatusInternalServerError, controller.Response{
				Success: false,
				Message: http.StatusText(http.StatusInternalServerError),
				Data:    nil,
			})
			return
		}

		if err != nil {
			ctx.JSON(result.StatusCode, controller.Response{
				Success: false,
				Message: err.Error(),
				Data:    result.Data,
			})
			return
		}

		ctx.JSON(result.StatusCode, controller.Response{
			Success: true,
			Message: result.Message,
			Data:    result.Data,
		})
	}
}

func RegisterRoutes(router *gin.Engine, controllers *Controllers) {
	router.GET("/", serveHTTP(controllers.RootController.Ok))

	// apiRoutes := router.Group("/api")
	// {
	// 	userRoutes := apiRoutes.Group("/user")
	// 	{
	// 		userRoutes.GET("/")
	// 		userRoutes.GET("/:id")
	// 		userRoutes.POST("/")
	// 		userRoutes.PUT("/:id")
	// 		userRoutes.DELETE("/:id")
	// 	}
	// }
}
