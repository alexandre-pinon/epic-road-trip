package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	routes := server.Group("/api")
	{
		userRoutes := routes.Group("/user")
		{
			userRoutes.GET("/", func(c *gin.Context) {
				c.JSON(200, "TODO: implement get all users")
			})

			userRoutes.GET("/:id", func(c *gin.Context) {
				id := c.Param("id")
				c.JSON(200, fmt.Sprintf("TODO: implement get user %s", id))
			})

			userRoutes.POST("/", func(c *gin.Context) {
				c.JSON(200, "TODO: implement create user")
			})

			userRoutes.PUT("/:id", func(c *gin.Context) {
				id := c.Param("id")
				c.JSON(200, fmt.Sprintf("TODO: implement update user %s", id))
			})

			userRoutes.DELETE("/:id", func(c *gin.Context) {
				id := c.Param("id")
				c.JSON(200, fmt.Sprintf("TODO: implement delete user %s", id))
			})
		}
	}

	server.Run(":8080")
}
