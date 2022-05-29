package service

import (
	"github.com/alexandre-pinon/epic-road-trip/repository"
)

type authService struct {
	userRepository repository.UserRepository
}

type AuthService interface {
}

func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{repo}
}
