package app

import (
	"github.com/alexandre-pinon/epic-road-trip/utils"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, controllers *Controllers) {
	router.GET("/", utils.ServeHTTP(controllers.RootController.Ok))

	apiRoutes := router.Group("/api")
	{
		userRoutes := apiRoutes.Group("/user")
		{
			userRoutes.GET("/", utils.ServeHTTP(controllers.UserController.GetAllUsers))
			userRoutes.GET("/:id", utils.ServeHTTP(controllers.UserController.GetUserByID))
			userRoutes.POST("/", utils.ServeHTTP(controllers.UserController.CreateUser))
		}
	}
}
