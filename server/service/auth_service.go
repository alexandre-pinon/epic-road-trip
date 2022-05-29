package service

import (
	"github.com/alexandre-pinon/epic-road-trip/repository"
	jwt "github.com/appleboy/gin-jwt/v2"
)

type authService struct {
	userRepository repository.UserRepository
}

type AuthService interface {
	PayloadFunc(data interface{}) jwt.MapClaims
}

func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{repo}
}

func (svc *authService) PayloadFunc(data interface{}) jwt.MapClaims {
	return nil
}
