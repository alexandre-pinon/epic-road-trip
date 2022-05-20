package app

import (
	"log"

	"github.com/alexandre-pinon/epic-road-trip/config"
	"github.com/gin-gonic/gin"
)

func InitApp() {
	log.Print("Initializing app...")

	configs := config.GetConfig()
	db := config.ConnectDB(configs)
	defer config.DisconnectDB(db.Client())

	// repos := SetupRepositories(db)
	services := Services{}
	controllers := SetupControllers(&services)

	router := gin.Default()

	RegisterRoutes(router, controllers)

	router.Run()
}
