package controller

import (
	"log"
	"time"

	"github.com/alexandre-pinon/epic-road-trip/config"
	"github.com/alexandre-pinon/epic-road-trip/service"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type authController struct {
	authService    service.AuthService
	authMiddleware jwt.GinJWTMiddleware
}

type AuthController interface {
	HandleLogin(ctx *gin.Context)
	JWTMiddleware() gin.HandlerFunc
}

func NewAuthController(cfg *config.Config, svc service.AuthService) AuthController {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           cfg.AppName,
		Key:             []byte(cfg.SecretKey),
		Timeout:         time.Minute * 15,
		MaxRefresh:      time.Hour * 1,
		IdentityKey:     jwt.IdentityKey,
		PayloadFunc:     svc.PayloadFunc,
		IdentityHandler: svc.IdentityHandler,
		Authenticator:   svc.Authenticator,
		Unauthorized:    svc.Unauthorized,
		LoginResponse:   svc.LoginResponse,
		TokenLookup:     "cookie:jwt",
		SendCookie:      true,
		CookieHTTPOnly:  true,
	})
	if err != nil {
		log.Fatal("Error initializing auth")
	}

	return &authController{svc, *authMiddleware}
}

func (ctrl *authController) HandleLogin(ctx *gin.Context) {
	ctrl.authMiddleware.LoginHandler(ctx)
}

func (ctrl *authController) JWTMiddleware() gin.HandlerFunc {
	return ctrl.authMiddleware.MiddlewareFunc()
}
