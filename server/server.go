package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alexandre-pinon/epic-road-trip/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	_ = db.New(os.Getenv("MONGO_URI"), os.Getenv("DB_NAME"))

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

	server.Run()
}
