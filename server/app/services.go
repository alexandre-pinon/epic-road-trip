package app

import (
	"github.com/alexandre-pinon/epic-road-trip/config"
	"github.com/alexandre-pinon/epic-road-trip/service"
)

type Services struct {
	UserService   service.UserService
	AuthService   service.AuthService
	GoogleService service.GoogleService
}

func SetupServices(cfg *config.Config, repos *Repositories) *Services {
	return &Services{
		UserService:   service.NewUserService(repos.UserRepository),
		AuthService:   service.NewAuthService(repos.UserRepository),
		GoogleService: service.NewGoogleService(cfg),
	}
}
