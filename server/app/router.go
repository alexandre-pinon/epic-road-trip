package app

import (
	"github.com/alexandre-pinon/epic-road-trip/middleware"
	"github.com/alexandre-pinon/epic-road-trip/utils"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(router *gin.Engine, controllers *Controllers) {
	rootCtrl := controllers.RootController
	authMiddleware := controllers.AuthController.JWTMiddleware()
	authCtrl := controllers.AuthController
	userCtrl := controllers.UserController
	roadtripCtrl := controllers.RoadtripController

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	apiRoutes := router.Group("/api/v1")
	{
		authRoutes := apiRoutes.Group("/auth")
		{
			authRoutes.POST("/register", utils.ServeHTTP(userCtrl.CreateUser))
			authRoutes.POST("/login", authCtrl.HandleLogin)
			authRoutes.POST("/logout", authMiddleware, authCtrl.HandleLogout)
			authRoutes.POST("/refresh_token", authCtrl.HandleRefresh)
		}
		userRoutes := apiRoutes.Group("/user")
		{
			userRoutes.GET("/", utils.ServeHTTP(userCtrl.GetAllUsers))
			userRoutes.GET("/:id", middleware.CheckID(), utils.ServeHTTP(userCtrl.GetUserByID))

			userRoutes.Use(authMiddleware)

			userRoutes.POST("/", utils.ServeHTTP(userCtrl.CreateUser))
			userRoutes.PUT("/:id", middleware.CheckID(), utils.ServeHTTP(userCtrl.UpdateUser))
			userRoutes.DELETE("/:id", middleware.CheckID(), utils.ServeHTTP(userCtrl.DeleteUser))
		}
		roadtripRoutes := apiRoutes.Group("/roadtrip")
		{
			roadtripRoutes.POST("/", utils.ServeHTTP(roadtripCtrl.CreateRoadtrip))
			roadtripRoutes.DELETE("/:id", middleware.CheckID(), utils.ServeHTTP(roadtripCtrl.DeleteRoadtrip))

			roadtripRoutes.POST("/enjoy", utils.ServeHTTP(roadtripCtrl.Enjoy))
			roadtripRoutes.POST("/sleep", utils.ServeHTTP(roadtripCtrl.Sleep))
			roadtripRoutes.POST("/eat", utils.ServeHTTP(roadtripCtrl.Eat))
			roadtripRoutes.POST("/drink", utils.ServeHTTP(roadtripCtrl.Drink))
			roadtripRoutes.POST("/travel/:mode", middleware.CheckTravelMode(), utils.ServeHTTP(roadtripCtrl.Travel))
		}
		apiRoutes.GET("/", utils.ServeHTTP(rootCtrl.Healthcheck))
	}
}
