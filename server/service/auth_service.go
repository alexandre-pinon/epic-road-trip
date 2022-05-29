package service

import (
	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/alexandre-pinon/epic-road-trip/repository"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type authService struct {
	userRepository repository.UserRepository
}

type AuthService interface {
	PayloadFunc(data interface{}) jwt.MapClaims
	IdentityHandler(ctx *gin.Context) interface{}
}

func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{repo}
}

func (svc *authService) PayloadFunc(data interface{}) jwt.MapClaims {
	if user, ok := data.(*model.User); ok {
		return jwt.MapClaims{
			jwt.IdentityKey: user.ID,
		}
	}
	return jwt.MapClaims{}
}

func (svc *authService) IdentityHandler(ctx *gin.Context) interface{} {
	return &struct{}{}
}
