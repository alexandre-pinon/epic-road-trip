package main

import (
	"example/hello/controller"
	"example/hello/service"
	"github.com/gin-gonic/gin"
)

var(
	activityService service.ActivityService = service.NewActivity()
	activityController controller.ActivityController = controller.NewActivity(activityService)

	userService service.UserService = service.NewUser()
	userController controller.UserController = controller.NewUser(userService)
)
 

func main() {
    server := gin.Default()

	server.GET("/activities", func (ctx *gin.Context)  {
		ctx.JSON(200 , activityController.FindAllActivities())
	})

	server.POST("/activity", func (ctx *gin.Context)  {
		ctx.JSON(200 , activityController.SaveActivity(ctx))
	})

	server.GET("/users", func (ctx *gin.Context)  {
		ctx.JSON(200 , userController.FindAllUsers())
	})

	server.POST("/user", func (ctx *gin.Context)  {
		ctx.JSON(200 , userController.SaveUser(ctx))
	})


	 server.Run(":8080")
} 
