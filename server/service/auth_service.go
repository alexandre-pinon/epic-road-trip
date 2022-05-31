package service

import (
	"errors"

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
	Authenticator(c *gin.Context) (interface{}, error)
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

	user, err := svc.userRepository.GetUserByID(ID)
	if err != nil {
		return nil
	}

	return user
}

func (svc *authService) Authenticator(c *gin.Context) (interface{}, error) {
	var userLogin model.UserLogin
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		return nil, errors.New("missing email or password")
	}

	user, err := svc.userRepository.GetUserByEmail(userLogin.Email)
	if err != nil {
		return nil, errors.New("incorrect email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(userLogin.Password)); err != nil {
		return nil, errors.New("incorrect email or password")
	}

	return user, nil
}
