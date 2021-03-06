package service

import (
	"errors"
	"time"

	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/alexandre-pinon/epic-road-trip/repository"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	userRepository repository.UserRepository
}

type AuthService interface {
	PayloadFunc(data interface{}) jwt.MapClaims
	IdentityHandler(ctx *gin.Context) interface{}
	Authenticator(ctx *gin.Context) (interface{}, error)
	Unauthorized(ctx *gin.Context, code int, message string)
	LoginResponse(ctx *gin.Context, code int, token string, expire time.Time)
	LogoutResponse(ctx *gin.Context, code int)
	RefreshResponse(ctx *gin.Context, code int, message string, time time.Time)
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
	claims := jwt.ExtractClaims(ctx)
	ID, ok := claims[jwt.IdentityKey].(primitive.ObjectID)
	if !ok {
		return nil
	}

	user, err := svc.userRepository.GetUserByID(ID, false)
	if err != nil {
		return nil
	}

	return user
}

func (svc *authService) Authenticator(ctx *gin.Context) (interface{}, error) {
	var userLogin model.UserLogin
	if err := ctx.ShouldBindJSON(&userLogin); err != nil {
		return nil, errors.New("missing email or password")
	}

	user, err := svc.userRepository.GetUserByEmail(userLogin.Email)
	if err != nil {
		return nil, errors.New("incorrect email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(userLogin.Password)); err != nil {
		return nil, errors.New("incorrect email or password")
	}

	ctx.Set("userID", user.ID.Hex())

	return user, nil
}

func (svc *authService) Unauthorized(ctx *gin.Context, code int, message string) {
	ctx.JSON(code, &model.AppResponse{
		Success:   false,
		Message:   message,
		Data:      struct{}{},
		ValErrors: []model.ValError{},
	})
}

func (svc *authService) LoginResponse(ctx *gin.Context, code int, token string, expire time.Time) {
	userID, _ := ctx.Get("userID")
	dataID := struct {
		ID string `json:"id"`
	}{ID: userID.(string)}
	ctx.JSON(code, &model.AppResponse{
		Success:   true,
		Message:   "Login successfully",
		Data:      dataID,
		ValErrors: []model.ValError{},
	})
}

func (svc *authService) LogoutResponse(ctx *gin.Context, code int) {
	ctx.JSON(code, &model.AppResponse{
		Success:   true,
		Message:   "Logout successfully",
		Data:      struct{}{},
		ValErrors: []model.ValError{},
	})
}

func (svc *authService) RefreshResponse(ctx *gin.Context, code int, message string, time time.Time) {
	ctx.JSON(code, &model.AppResponse{
		Success:   true,
		Message:   "Token refreshed successfully",
		Data:      struct{}{},
		ValErrors: []model.ValError{},
	})
}
