package service

import (
	"testing"

	"github.com/alexandre-pinon/epic-road-trip/mocks"
	"github.com/alexandre-pinon/epic-road-trip/model"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type authServiceSuite struct {
	suite.Suite
	repo        *mocks.UserRepository
	authService AuthService
}

func (suite *authServiceSuite) SetupTest() {
	repo := new(mocks.UserRepository)
	authService := NewAuthService(repo)

	suite.repo = repo
	suite.authService = authService
}

func (suite *authServiceSuite) TestPayloadFunc_Positive() {
	ID := primitive.NewObjectID()
	payload := &model.User{
		ID:             ID,
		Firstname:      "yoimiya",
		Lastname:       "naganohara",
		Email:          "yoimiya.naganohara@gmail.com",
		HashedPassword: "12345678",
		Phone:          "+33612345678",
	}
	claims := suite.authService.PayloadFunc(payload)

	suite.NotNil(claims)
	suite.NotEmpty(claims)
	suite.Equal(ID, claims[jwt.IdentityKey])
}

func (suite *authServiceSuite) TestPayloadFunc_InvalidData_Positive() {
	payload := struct{}{}
	claims := suite.authService.PayloadFunc(payload)

	suite.NotNil(claims)
	suite.Empty(claims)
}

func TestAuthService(t *testing.T) {
	suite.Run(t, new(authServiceSuite))
}
