package app

import "github.com/alexandre-pinon/epic-road-trip/controller"

type Controllers struct {
	RootController controller.RootController
	UserController controller.UserController
}

func SetupControllers(services *Services) *Controllers {
	return &Controllers{
		RootController: controller.NewRootController(),
		UserController: controller.NewUserController(services.UserService),
	}
}
