package app

import "github.com/alexandre-pinon/epic-road-trip/controller"

type Controllers struct {
	RootController controller.RootController
}

func SetupControllers(services *Services) *Controllers {
	rootController := controller.NewRootController()

	return &Controllers{
		RootController: rootController,
	}
}
