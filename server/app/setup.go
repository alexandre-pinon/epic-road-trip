package app

import (
	"log"

	"github.com/alexandre-pinon/epic-road-trip/config"
	"github.com/alexandre-pinon/epic-road-trip/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitApp() {
	log.Print("Initializing app...")

	cfg := config.GetConfig("")
	db := config.ConnectDB(cfg)
	defer config.DisconnectDB(cfg, db.Client())

	repositories := SetupRepositories(db)
	services := SetupServices(cfg, repositories)
	controllers := SetupControllers(cfg, services, repositories)

	log.Print("Initializing swagger...")

	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AllowCredentials = true
	router.Use(cors.New(corsConfig))
	router.StaticFile("/docs.html", "docs/index.html")

	RegisterRoutes(router, controllers)

	router.Run()
}

func InitDocs() {
	docs.SwaggerInfo.Title = "Epic Road Trip API"
	docs.SwaggerInfo.Description = "This is the API of the BEST road trip planner of the market."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8000"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}
}
