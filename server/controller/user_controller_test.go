package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alexandre-pinon/epic-road-trip/mocks"
	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/alexandre-pinon/epic-road-trip/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type userControllerSuite struct {
	suite.Suite
	svc        *mocks.UserService
	ctrl       UserController
	testServer *httptest.Server
}

func (suite *userControllerSuite) SetupTest() {
	svc := new(mocks.UserService)
	ctrl := NewUserController(svc)

	router := gin.Default()
	apiRoutes := router.Group("/api")
	{
		userRoutes := apiRoutes.Group("/user")
		{
			userRoutes.POST("/", utils.ServeHTTP(ctrl.CreateUser))
		}
	}
	testServer := httptest.NewServer(router)

	suite.testServer = testServer
	suite.svc = svc
	suite.ctrl = ctrl
}

func (suite *userControllerSuite) TearDownTest() {
	defer suite.testServer.Close()
}

func (suite *userControllerSuite) TestCreateUser_Positive() {
	user := model.User{
		Firstname: "yoimiya",
		Lastname:  "naganohara",
		Email:     "yoimiya.naganohara@gmail.com",
		Password:  "12345678",
		Phone:     "+33612345678",
		Trips:     []*model.RoadTrip{},
	}

	suite.svc.On("CreateUser", &user).Return(nil)

	requestBody, err := json.Marshal(&user)
	suite.NoError(err, "can not marshal struct to json")

	response, err := http.Post(
		fmt.Sprintf("%s/api/user", suite.testServer.URL),
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	responseBody := model.Response{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	suite.Equal(http.StatusCreated, response.StatusCode)
	suite.Equal("User created successfully", responseBody.Message)
	suite.svc.AssertExpectations(suite.T())
}

func TestUserHandler(t *testing.T) {
	suite.Run(t, new(userControllerSuite))
}
