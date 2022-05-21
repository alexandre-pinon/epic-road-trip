package app

import (
	"log"
	"os"
	"strconv"

	"github.com/alexandre-pinon/epic-road-trip/config"
	"github.com/gin-gonic/gin"
)

func InitApp() {
	env, err := strconv.Atoi(os.Getenv("GO_MODE"))
	if err != nil {
		log.Fatalf("Invalid GO_MODE env variable, please specify either:\n0 -> Dev\n1 -> Prod\n2 -> Test")
	}

	log.Print("Initializing app...")

	cfg := config.GetConfig(config.Env(env))
	db := config.ConnectDB(cfg)
	defer config.DisconnectDB(cfg, db.Client())

	// repos := SetupRepositories(db)
	services := Services{}
	controllers := SetupControllers(&services)

	router := gin.Default()

	RegisterRoutes(router, controllers)

	router.Run()
}
