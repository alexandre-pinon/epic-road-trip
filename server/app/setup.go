package app

import (
	"log"
	"os"

	"github.com/alexandre-pinon/epic-road-trip/config"
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

	router := gin.Default()

	RegisterRoutes(router, controllers)

	router.Run()
}
