package app

import "github.com/alexandre-pinon/epic-road-trip/service"

type Services struct {
	UserService service.UserService
	AuthService service.AuthService
}

func SetupServices(repos *Repositories) *Services {
	return &Services{
		UserService: service.NewUserService(repos.UserRepository),
		AuthService: service.NewAuthService(repos.UserRepository),
	}
}
