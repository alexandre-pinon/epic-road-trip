package app

import (
	"log"
	"os"

	"github.com/alexandre-pinon/epic-road-trip/config"
	"github.com/alexandre-pinon/epic-road-trip/docs"
	"github.com/gin-gonic/gin"
)

func InitApp() {
	env := os.Getenv("GO_MODE")

	if err := config.Env(env).IsValid(); err != nil {
		log.Fatal(err)
	}

	log.Print("Initializing app...")

	cfg := config.GetConfig(config.Env(env))
	db := config.ConnectDB(cfg)
	defer config.DisconnectDB(cfg, db.Client())

	repositories := SetupRepositories(db)
	services := SetupServices(repositories)
	controllers := SetupControllers(cfg, services)

	log.Print("Initializing swagger...")

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Epic Road Trip API"
	docs.SwaggerInfo.Description = "This is the API of the BEST road trip planner of the market."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8000"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}

	router := gin.Default()

	RegisterRoutes(router, controllers)

	router.Run()
}
