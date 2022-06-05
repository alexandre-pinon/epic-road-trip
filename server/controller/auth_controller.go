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
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body model.RegisterRequest true "firstname, lastname, email, password, phone"
// @Success 200 {object} model.RegisterSuccess "Success"
// @Failure 400 {object} model.InvalidJsonBody "Invalid request body"
// @Failure 500 {object} model.InternalServerError "Internal server error"
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
// @Tags Auth
// @Accept json
// @Produce json
// @Param userLogin body model.LoginRequest true "email & password"
// @Success 200 {object} model.LoginSuccess "Success"
// @Failure 401 {object} model.LoginFailureCredentials "Missing/Incorrect credentials"
// @Failure 500 {object} model.InternalServerError "Internal server error"
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
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} model.RefreshSuccess "Success"
// @Failure 401 {object} model.Unauthorized "Missing/Expired token"
// @Failure 500 {object} model.InternalServerError "Internal server error"
// @Router /auth/refresh_token [post]
func (ctrl *authController) HandleRefresh(ctx *gin.Context) {
	ctrl.authMiddleware.RefreshHandler(ctx)
}

// Logout godoc
// @Summary Logout
// @Description Logout user by removing jwt cookie
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} model.LogoutSuccess "Success"
// @Failure 401 {object} model.Unauthorized "Missing/Expired token"
// @Failure 500 {object} model.InternalServerError "Internal server error"
// @Router /auth/logout [post]
func (ctrl *authController) HandleLogout(ctx *gin.Context) {
	ctrl.authMiddleware.LogoutHandler(ctx)
}
