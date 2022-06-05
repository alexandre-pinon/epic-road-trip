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
	HandleRefresh(ctx *gin.Context)
	HandleLogout(ctx *gin.Context)
}

// Register godoc
// @Summary Register
// @Description Register user given valid firstname, lastname, email (unique), password, phone (unique)
// @Tags auth
// @Accept json
// @Produce json
// @Param user body model.RegisterRequest true "firstname, lastname, email, password, phone"
// @Success 200 {object} model.RegisterSuccess "Successful register"
// @Failure 400 {object} model.RegisterFailureInvalid "Invalid request body"
// @Router /auth/register [post]
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
		LogoutResponse:  svc.LogoutResponse,
		RefreshResponse: svc.RefreshResponse,
		TokenLookup:     "cookie:jwt",
		SendCookie:      true,
		CookieHTTPOnly:  true,
	})
	if err != nil {
		log.Fatal("Error initializing auth")
	}

	return &authController{svc, *authMiddleware}
}

// Login godoc
// @Summary Login
// @Description Login user given valid email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param userLogin body model.LoginRequest true "email & password"
// @Success 200 {object} model.LoginSuccess "Successful login"
// @Failure 401 {object} model.LoginFailureCredentials "Missing/Incorrect credentials"
// @Router /auth/login [post]
func (ctrl *authController) HandleLogin(ctx *gin.Context) {
	ctrl.authMiddleware.LoginHandler(ctx)
}

func (ctrl *authController) JWTMiddleware() gin.HandlerFunc {
	return ctrl.authMiddleware.MiddlewareFunc()
}

// Refresh godoc
// @Summary Refresh
// @Description Refresh user's access token given valid refresh token
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} model.RefreshSuccess "Token refreshed successfully"
// @Failure 401 {object} model.RefreshFailure "Missing/Expired token"
// @Router /auth/refresh_token [post]
func (ctrl *authController) HandleRefresh(ctx *gin.Context) {
	ctrl.authMiddleware.RefreshHandler(ctx)
}

// Logout godoc
// @Summary Logout
// @Description Logout user by removing jwt cookie
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} model.LogoutSuccess "Successful register"
// @Failure 401 {object} model.LogoutFailure "Missing/Expired token"
// @Router /auth/logout [post]
func (ctrl *authController) HandleLogout(ctx *gin.Context) {
	ctrl.authMiddleware.LogoutHandler(ctx)
}
