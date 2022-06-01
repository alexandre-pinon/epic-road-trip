package app

import (
	"github.com/alexandre-pinon/epic-road-trip/middleware"
	"github.com/alexandre-pinon/epic-road-trip/utils"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, controllers *Controllers) {
	authMiddleware := controllers.AuthController.JWTMiddleware()

	router.GET("/", utils.ServeHTTP(controllers.RootController.Ok))

	apiRoutes := router.Group("/api")
	{
		authRoutes := apiRoutes.Group("/auth")
		{
			authRoutes.POST("/register", utils.ServeHTTP(controllers.UserController.CreateUser))
			authRoutes.POST("/login", controllers.AuthController.HandleLogin)
			authRoutes.POST("/logout", authMiddleware, controllers.AuthController.HandleLogout)
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
	}
}
