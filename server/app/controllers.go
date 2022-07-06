package app

import (
	"github.com/alexandre-pinon/epic-road-trip/config"
	"github.com/alexandre-pinon/epic-road-trip/controller"
)

type Controllers struct {
	RootController     controller.RootController
	UserController     controller.UserController
	AuthController     controller.AuthController
	RoadtripController controller.RoadtripController
}

func SetupControllers(cfg *config.Config, services *Services, repositories *Repositories) *Controllers {
	return &Controllers{
		RootController:     controller.NewRootController(),
		UserController:     controller.NewUserController(services.UserService),
		AuthController:     controller.NewAuthController(cfg, services.AuthService),
		RoadtripController: controller.NewRoadtripController(cfg, services.UserService, services.GoogleService, services.AmadeusService, repositories.TripStepRepository),
	}
}
