package service

import (
	"net/http/httptest"
	"testing"

	"github.com/alexandre-pinon/epic-road-trip/mocks"
	"github.com/alexandre-pinon/epic-road-trip/model"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type authServiceSuite struct {
	suite.Suite
	repo        *mocks.UserRepository
	authService AuthService
}

func (suite *authServiceSuite) SetupTest() {
	gin.SetMode(gin.ReleaseMode)
	repo := new(mocks.UserRepository)
	authService := NewAuthService(repo)

	suite.repo = repo
	suite.authService = authService
}

func (suite *authServiceSuite) TestPayloadFunc_Positive() {
	ID := primitive.NewObjectID()
	payload := model.User{
		ID:             ID,
		Firstname:      "yoimiya",
		Lastname:       "naganohara",
		Email:          "yoimiya.naganohara@gmail.com",
		HashedPassword: "12345678",
		Phone:          "+33612345678",
	}
	claims := suite.authService.PayloadFunc(&payload)

	suite.NotEmpty(claims)
	suite.Equal(ID, claims[jwt.IdentityKey])
}

func (suite *authServiceSuite) TestPayloadFunc_InvalidData_Positive() {
	payload := struct{}{}
	claims := suite.authService.PayloadFunc(&payload)

	suite.NotNil(claims)
	suite.Empty(claims)
}

func (suite *authServiceSuite) TestIdentityHandler_Positive() {
	ID := primitive.NewObjectID()
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Set("JWT_PAYLOAD", jwt.MapClaims{jwt.IdentityKey: ID})

	user := model.User{
		Firstname:      "yoimiya",
		Lastname:       "naganohara",
		Email:          "yoimiya.naganohara@gmail.com",
		HashedPassword: "12345678",
		Phone:          "+33612345678",
		Trips:          []*model.RoadTrip{},
	}

	suite.repo.On("GetUserByID", ID).Return(&user, nil)

	result, ok := suite.authService.IdentityHandler(ctx).(*model.User)
	suite.Require().True(ok)
	suite.Equal(user, *result)
}

func (suite *authServiceSuite) TestIdentityHandler_NotFound_Positive() {
	ID := primitive.NewObjectID()
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Set("JWT_PAYLOAD", jwt.MapClaims{jwt.IdentityKey: ID})

	suite.repo.On("GetUserByID", ID).Return(nil, mongo.ErrNoDocuments)

	result := suite.authService.IdentityHandler(ctx)
	suite.Nil(result)
}

func (suite *authServiceSuite) TestIdentityHandler_EmptyContext_Positive() {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	result := suite.authService.IdentityHandler(ctx)
	suite.Nil(result)
}

func TestAuthService(t *testing.T) {
	suite.Run(t, new(authServiceSuite))
}
