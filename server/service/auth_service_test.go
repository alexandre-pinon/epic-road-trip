package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alexandre-pinon/epic-road-trip/mocks"
	"github.com/alexandre-pinon/epic-road-trip/model"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
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
	suite.repo.AssertExpectations(suite.T())
}

func (suite *authServiceSuite) TestIdentityHandler_NotFound_Positive() {
	ID := primitive.NewObjectID()
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Set("JWT_PAYLOAD", jwt.MapClaims{jwt.IdentityKey: ID})

	suite.repo.On("GetUserByID", ID).Return(nil, mongo.ErrNoDocuments)

	result := suite.authService.IdentityHandler(ctx)
	suite.Nil(result)
	suite.repo.AssertExpectations(suite.T())
}

func (suite *authServiceSuite) TestIdentityHandler_EmptyContext_Positive() {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	result := suite.authService.IdentityHandler(ctx)
	suite.Nil(result)
}

func (suite *authServiceSuite) TestAuthenticator_InvalidJson_Negative() {
	rec := httptest.NewRecorder()
	ctx, eng := gin.CreateTestContext(rec)
	eng.POST("/test", func(ctx *gin.Context) {
		user, err := suite.authService.Authenticator(ctx)
		ctx.JSON(http.StatusBadRequest, &model.AppResponse{
			Message: err.Error(),
			Data:    user,
		})
	})

	requestBody, err := json.Marshal(nil)
	suite.NoError(err, "can not marshal struct to json")
	ctx.Request, _ = http.NewRequest(http.MethodPost, "/test", bytes.NewBuffer(requestBody))
	eng.ServeHTTP(rec, ctx.Request)

	responseBody := model.AppResponse{}
	json.NewDecoder(rec.Body).Decode(&responseBody)

	suite.Equal("missing email or password", responseBody.Message)
	suite.Nil(responseBody.Data)
}

func (suite *authServiceSuite) TestAuthenticator_IncorrectEmail_Negative() {
	rec := httptest.NewRecorder()
	ctx, eng := gin.CreateTestContext(rec)
	eng.POST("/test", func(ctx *gin.Context) {
		user, err := suite.authService.Authenticator(ctx)
		ctx.JSON(http.StatusBadRequest, &model.AppResponse{
			Message: err.Error(),
			Data:    user,
		})
	})

	userLogin := model.UserLogin{
		Email:    "yoimiya.naganohara@gmail.com",
		Password: "12345678",
	}
	suite.repo.On("GetUserByEmail", userLogin.Email).Return(nil, mongo.ErrNoDocuments)

	requestBody, err := json.Marshal(&userLogin)
	suite.NoError(err, "can not marshal struct to json")
	ctx.Request, _ = http.NewRequest(http.MethodPost, "/test", bytes.NewBuffer(requestBody))
	eng.ServeHTTP(rec, ctx.Request)

	responseBody := model.AppResponse{}
	json.NewDecoder(rec.Body).Decode(&responseBody)

	suite.Equal("incorrect email or password", responseBody.Message)
	suite.Nil(responseBody.Data)
	suite.repo.AssertExpectations(suite.T())
}

func (suite *authServiceSuite) TestAuthenticator_IncorrectPassword_Negative() {
	rec := httptest.NewRecorder()
	ctx, eng := gin.CreateTestContext(rec)
	eng.POST("/test", func(ctx *gin.Context) {
		user, err := suite.authService.Authenticator(ctx)
		ctx.JSON(http.StatusBadRequest, &model.AppResponse{
			Message: err.Error(),
			Data:    user,
		})
	})

	password := "12345678"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	suite.NoError(err, "no error hashing password")

	user := model.User{
		Firstname:      "yoimiya",
		Lastname:       "naganohara",
		Email:          "yoimiya.naganohara@gmail.com",
		HashedPassword: string(hashedPassword),
		Phone:          "+33612345678",
		Trips:          []*model.RoadTrip{},
	}
	userLogin := model.UserLogin{
		Email:    "yoimiya.naganohara@gmail.com",
		Password: password + "bad",
	}
	suite.repo.On("GetUserByEmail", userLogin.Email).Return(&user, nil)

	requestBody, err := json.Marshal(&userLogin)
	suite.NoError(err, "can not marshal struct to json")
	ctx.Request, _ = http.NewRequest(http.MethodPost, "/test", bytes.NewBuffer(requestBody))
	eng.ServeHTTP(rec, ctx.Request)

	responseBody := model.AppResponse{}
	json.NewDecoder(rec.Body).Decode(&responseBody)

	suite.Equal("incorrect email or password", responseBody.Message)
	suite.Nil(responseBody.Data)
	suite.repo.AssertExpectations(suite.T())
}

func (suite *authServiceSuite) TestAuthenticator_Positive() {
	rec := httptest.NewRecorder()
	ctx, eng := gin.CreateTestContext(rec)
	eng.POST("/test", func(ctx *gin.Context) {
		user, err := suite.authService.Authenticator(ctx)
		message := ""
		if err != nil {
			message = err.Error()
		}
		ctx.JSON(http.StatusOK, &model.AppResponse{
			Message: message,
			Data:    user,
		})
	})

	password := "12345678"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	suite.NoError(err, "no error hashing password")

	user := model.User{
		Firstname:      "yoimiya",
		Lastname:       "naganohara",
		Email:          "yoimiya.naganohara@gmail.com",
		HashedPassword: string(hashedPassword),
		Phone:          "+33612345678",
		Trips:          []*model.RoadTrip{},
	}
	userLogin := model.UserLogin{
		Email:    "yoimiya.naganohara@gmail.com",
		Password: password,
	}
	suite.repo.On("GetUserByEmail", userLogin.Email).Return(&user, nil)

	requestBody, err := json.Marshal(&userLogin)
	suite.NoError(err, "can not marshal struct to json")
	ctx.Request, _ = http.NewRequest(http.MethodPost, "/test", bytes.NewBuffer(requestBody))
	eng.ServeHTTP(rec, ctx.Request)

	responseBody := model.AppResponse{}
	json.NewDecoder(rec.Body).Decode(&responseBody)

	suite.Empty(responseBody.Message)
	suite.NotEmpty(responseBody.Data)
	suite.repo.AssertExpectations(suite.T())
}

func TestAuthService(t *testing.T) {
	gin.SetMode(gin.TestMode)
	suite.Run(t, new(authServiceSuite))
}
