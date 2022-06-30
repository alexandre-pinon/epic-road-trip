package app

import (
	"github.com/alexandre-pinon/epic-road-trip/config"
	"github.com/alexandre-pinon/epic-road-trip/controller"
)

type Controllers struct {
	RootController     controller.RootController
	UserController     controller.UserController
	AuthController     controller.AuthController
	RoadtripController controller.RoadTripController
}

func SetupControllers(cfg *config.Config, services *Services) *Controllers {
	return &Controllers{
		RootController:     controller.NewRootController(),
		UserController:     controller.NewUserController(services.UserService),
		AuthController:     controller.NewAuthController(cfg, services.AuthService),
		RoadtripController: controller.NewRoadTripController(cfg, services.GoogleService, services.AmadeusService),
	}
}
