package app

import (
	"github.com/alexandre-pinon/epic-road-trip/middleware"
	"github.com/alexandre-pinon/epic-road-trip/utils"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(router *gin.Engine, controllers *Controllers) {
	authMiddleware := controllers.AuthController.JWTMiddleware()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	apiRoutes := router.Group("/api/v1")
	{
		authRoutes := apiRoutes.Group("/auth")
		{
			authRoutes.POST("/register", utils.ServeHTTP(controllers.UserController.CreateUser))
			authRoutes.POST("/login", controllers.AuthController.HandleLogin)
			authRoutes.POST("/logout", authMiddleware, controllers.AuthController.HandleLogout)
			authRoutes.POST("/refresh_token", controllers.AuthController.HandleRefresh)
		}
		userRoutes := apiRoutes.Group("/user")
		{
			userRoutes.GET("/", utils.ServeHTTP(controllers.UserController.GetAllUsers))
			userRoutes.GET("/:id", middleware.CheckID(), utils.ServeHTTP(controllers.UserController.GetUserByID))

			userRoutes.Use(authMiddleware)

			userRoutes.POST("/", utils.ServeHTTP(controllers.UserController.CreateUser))
			userRoutes.PUT("/:id", middleware.CheckID(), utils.ServeHTTP(controllers.UserController.UpdateUser))
			userRoutes.DELETE("/:id", middleware.CheckID(), utils.ServeHTTP(controllers.UserController.DeleteUser))
		}
		roadtripRoutes := apiRoutes.Group("/roadtrip")
		{
			roadtripRoutes.POST("/enjoy", utils.ServeHTTP(controllers.RoadtripController.Enjoy))
			roadtripRoutes.POST("/travel/:mode", middleware.CheckTravelMode(), utils.ServeHTTP(controllers.RoadtripController.Travel))
		}
		apiRoutes.GET("/", utils.ServeHTTP(controllers.RootController.Healthcheck))
	}
}
