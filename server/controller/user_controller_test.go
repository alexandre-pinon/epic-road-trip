package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alexandre-pinon/epic-road-trip/middleware"
	"github.com/alexandre-pinon/epic-road-trip/mocks"
	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/alexandre-pinon/epic-road-trip/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

	router := gin.New()
	apiRoutes := router.Group("/api")
	{
		userRoutes := apiRoutes.Group("/user")
		{
			userRoutes.GET("/", utils.ServeHTTP(ctrl.GetAllUsers))
			userRoutes.GET("/:id", middleware.CheckID(), utils.ServeHTTP(ctrl.GetUserByID))
			userRoutes.POST("/", utils.ServeHTTP(ctrl.CreateUser))
			userRoutes.PUT("/:id", middleware.CheckID(), utils.ServeHTTP(ctrl.UpdateUser))
			userRoutes.DELETE("/:id", middleware.CheckID(), utils.ServeHTTP(ctrl.DeleteUser))
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

func (suite *userControllerSuite) TestGetAllUsers_EmptySlice_Positive() {
	emptyUsers := []model.User(nil)

	suite.svc.On("GetAllUsers").Return(&emptyUsers, nil)

	response, err := http.Get(fmt.Sprintf("%s/api/user", suite.testServer.URL))
	suite.NoError(err, "no error when calling this endpoint")
	defer response.Body.Close()

	responseBody := model.AppResponse{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	suite.Equal(http.StatusOK, response.StatusCode)
	suite.Equal("Users retrieved successfully", responseBody.Message)
	suite.Empty(responseBody.Data, "no users should be retrieved")
	suite.svc.AssertExpectations(suite.T())
}

func (suite *userControllerSuite) TestGetAllUsers_FilledSlice_Positive() {
	users := []model.User{
		{
			Firstname: "yoimiya",
			Lastname:  "naganohara",
			Email:     "yoimiya.naganohara@gmail.com",
			Phone:     "+33612345678",
			Trips:     []*model.Roadtrip{},
		},
		{
			Firstname: "hu",
			Lastname:  "tao",
			Email:     "hu.tao@gmail.com",
			Phone:     "+33623456789",
			Trips:     []*model.Roadtrip{},
		},
		{
			Firstname: "kokomi",
			Lastname:  "sangonomiya",
			Email:     "kokomi.sangonomiya@gmail.com",
			Phone:     "+33687654321",
			Trips:     []*model.Roadtrip{},
		},
	}

	suite.svc.On("GetAllUsers").Return(&users, nil)

	response, err := http.Get(fmt.Sprintf("%s/api/user", suite.testServer.URL))
	suite.NoError(err, "no error when calling this endpoint")
	defer response.Body.Close()

	responseBody := model.AppResponse{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	suite.Equal(http.StatusOK, response.StatusCode)
	suite.Equal("Users retrieved successfully", responseBody.Message)
	suite.NotEmpty(responseBody.Data, "users should be retrieved")
	suite.svc.AssertExpectations(suite.T())
}

func (suite *userControllerSuite) TestGetUserByID_Exists_Positive() {
	id := primitive.NewObjectID()
	user := model.User{
		Firstname: "yoimiya",
		Lastname:  "naganohara",
		Email:     "yoimiya.naganohara@gmail.com",
		Phone:     "+33612345678",
		Trips:     []*model.Roadtrip{},
	}

	suite.svc.On("GetUserByID", id).Return(&user, nil)
	response, err := http.Get(fmt.Sprintf("%s/api/user/%s", suite.testServer.URL, id.Hex()))
	suite.NoError(err, "no error when calling this endpoint")
	defer response.Body.Close()

	responseBody := model.AppResponse{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	suite.Equal(http.StatusOK, response.StatusCode)
	suite.Equal(fmt.Sprintf("User %s retrieved successfully", id.Hex()), responseBody.Message)
	suite.NotEmpty(responseBody.Data, "user should be retrieved")
	suite.svc.AssertExpectations(suite.T())
}

func (suite *userControllerSuite) TestGetUserByID_InvalidID_Negative() {
	id := primitive.NewObjectID()

	response, err := http.Get(fmt.Sprintf("%s/api/user/%s", suite.testServer.URL, id.Hex()+"bad"))
	suite.NoError(err, "no error when calling this endpoint")
	defer response.Body.Close()

	responseBody := model.AppResponse{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	suite.Equal(http.StatusBadRequest, response.StatusCode)
	suite.Equal("invalid id", responseBody.Message)
	suite.Empty(responseBody.Data, "user should not be retrieved")
}

func (suite *userControllerSuite) TestCreateUser_Positive() {
	userFormData := model.UserFormData{
		User: model.User{
			Firstname: "yoimiya",
			Lastname:  "naganohara",
			Email:     "yoimiya.naganohara@gmail.com",
			Phone:     "+33612345678",
		},
		Password: "12345678",
	}

	suite.svc.On("HashPassword", &userFormData).Return(nil)
	suite.svc.On("CreateUser", &userFormData.User).Return(nil)

	requestBody, err := json.Marshal(&userFormData)
	suite.NoError(err, "can not marshal struct to json")

	response, err := http.Post(
		fmt.Sprintf("%s/api/user", suite.testServer.URL),
		gin.MIMEJSON,
		bytes.NewBuffer(requestBody),
	)
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	responseBody := model.AppResponse{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	suite.Equal(http.StatusCreated, response.StatusCode)
	suite.Equal("User created successfully", responseBody.Message)
	suite.svc.AssertExpectations(suite.T())
}

func (suite *userControllerSuite) TestCreateUser_NilBody_Negative() {
	response, err := http.Post(
		fmt.Sprintf("%s/api/user", suite.testServer.URL),
		gin.MIMEJSON,
		bytes.NewBuffer(nil),
	)
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	responseBody := model.AppResponse{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	suite.Equal(http.StatusBadRequest, response.StatusCode)
	suite.Equal("EOF", responseBody.Message)
	suite.Empty(responseBody.Data)
	suite.Empty(responseBody.ValErrors)
}

func (suite *userControllerSuite) TestCreateUser_InvalidJSON_Negative() {
	userFormData := model.UserFormData{
		User: model.User{
			Firstname: "y",
			Lastname:  "naganoharanaganoharanaganoharanaganoharanaganoharanaganoharanaganoharanaganoharanaganohara",
			Email:     "bademail.com",
			Phone:     "-336123456789",
			Trips:     []*model.Roadtrip{},
		},
		Password: "root",
	}

	requestBody, err := json.Marshal(&userFormData)
	suite.NoError(err, "can not marshal struct to json")

	response, err := http.Post(
		fmt.Sprintf("%s/api/user", suite.testServer.URL),
		gin.MIMEJSON,
		bytes.NewBuffer(requestBody),
	)
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	responseBody := model.AppResponse{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	suite.Equal(http.StatusBadRequest, response.StatusCode)
	suite.Equal("invalid json request body", responseBody.Message)
	suite.Empty(responseBody.Data)
	suite.Equal("Firstname", responseBody.ValErrors[0].Field)
	suite.Equal("Should be at least 2 characters", responseBody.ValErrors[0].Message)
	suite.Equal("Lastname", responseBody.ValErrors[1].Field)
	suite.Equal("Should be less than 50 characters", responseBody.ValErrors[1].Message)
	suite.Equal("Email", responseBody.ValErrors[2].Field)
	suite.Equal("Invalid email format", responseBody.ValErrors[2].Message)
	suite.Equal("Phone", responseBody.ValErrors[3].Field)
	suite.Equal("Invalid phone format", responseBody.ValErrors[3].Message)
	suite.Equal("Password", responseBody.ValErrors[4].Field)
	suite.Equal("Should be at least 8 characters", responseBody.ValErrors[4].Message)
}

func (suite *userControllerSuite) TestCreateUser_DupKey_Negative() {
	userFormData := model.UserFormData{
		User: model.User{
			Firstname: "yoimiya",
			Lastname:  "naganohara",
			Email:     "yoimiya.naganohara@gmail.com",
			Phone:     "+33612345678",
		},
		Password: "12345678",
	}

	suite.svc.On("HashPassword", &userFormData).Return(nil)
	suite.svc.On("CreateUser", &userFormData.User).Return(&model.AppError{
		StatusCode: http.StatusInternalServerError,
		Err: mongo.WriteException{
			WriteErrors: mongo.WriteErrors{{
				Code:    11000,
				Message: "E11000 duplicate key error collection: dev-epic-road-trip-db.user index: email_1 dup key: { email: \"yoimiya.naganohara@gmail.com\" }"},
			},
		},
	})

	requestBody, err := json.Marshal(&userFormData)
	suite.NoError(err, "can not marshal struct to json")

	response, err := http.Post(
		fmt.Sprintf("%s/api/user", suite.testServer.URL),
		gin.MIMEJSON,
		bytes.NewBuffer(requestBody),
	)
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	responseBody := model.AppResponse{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	suite.Equal(http.StatusBadRequest, response.StatusCode)
	suite.Equal("invalid json request body", responseBody.Message)
	suite.Empty(responseBody.Data)
	suite.Equal("Email is already taken", responseBody.ValErrors[0].Message)
	suite.svc.AssertExpectations(suite.T())
}

func (suite *userControllerSuite) TestUpdateUser_Positive() {
	id := primitive.NewObjectID()
	user := model.User{
		Firstname: "yoiyoi",
		Lastname:  "miya",
		Email:     "yoiyoi.miya@gmail.com",
		Phone:     "+33612345678",
	}

	suite.svc.On("UpdateUser", id, &user).Return(nil)

	requestBody, err := json.Marshal(&user)
	suite.NoError(err, "can not marshal struct to json")

	request, err := http.NewRequest(
		http.MethodPut,
		fmt.Sprintf("%s/api/user/%s", suite.testServer.URL, id.Hex()),
		bytes.NewBuffer(requestBody),
	)
	suite.NoError(err, "no error when creating the request")
	request.Header.Add("Content-type", gin.MIMEJSON)

	response, err := http.DefaultClient.Do(request)
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	responseBody := model.AppResponse{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	suite.Equal(http.StatusOK, response.StatusCode)
	suite.Equal("User updated successfully", responseBody.Message)
	suite.svc.AssertExpectations(suite.T())
}

func (suite *userControllerSuite) TestUpdateUser_InvalidID_Negative() {
	id := primitive.NewObjectID()
	user := model.User{
		Firstname: "yoiyoi",
		Lastname:  "miya",
		Email:     "yoiyoi.miya@gmail.com",
		Phone:     "+33612345678",
	}

	requestBody, err := json.Marshal(&user)
	suite.NoError(err, "can not marshal struct to json")

	request, err := http.NewRequest(
		http.MethodPut,
		fmt.Sprintf("%s/api/user/%s", suite.testServer.URL, id.Hex()+"bad"),
		bytes.NewBuffer(requestBody),
	)
	suite.NoError(err, "no error when creating the request")
	request.Header.Add("Content-type", gin.MIMEJSON)

	response, err := http.DefaultClient.Do(request)
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	responseBody := model.AppResponse{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	suite.Equal(http.StatusBadRequest, response.StatusCode)
	suite.Equal("invalid id", responseBody.Message)
	suite.Empty(responseBody.Data, "user should not be retrieved")
}

func (suite *userControllerSuite) TestUpdateUser_InvalidJSON_Negative() {
	id := primitive.NewObjectID()
	user := model.User{
		Firstname: "y",
		Lastname:  "naganoharanaganoharanaganoharanaganoharanaganoharanaganoharanaganoharanaganoharanaganohara",
		Email:     "bademail.com",
		Phone:     "-336123456789",
	}

	requestBody, err := json.Marshal(&user)
	suite.NoError(err, "can not marshal struct to json")

	request, err := http.NewRequest(
		http.MethodPut,
		fmt.Sprintf("%s/api/user/%s", suite.testServer.URL, id.Hex()),
		bytes.NewBuffer(requestBody),
	)
	suite.NoError(err, "no error when creating the request")
	request.Header.Add("Content-type", gin.MIMEJSON)

	response, err := http.DefaultClient.Do(request)
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	responseBody := model.AppResponse{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	suite.Equal(http.StatusBadRequest, response.StatusCode)
	suite.Equal("invalid json request body", responseBody.Message)
	suite.Empty(responseBody.Data)
	suite.Require().NotEmpty(responseBody.ValErrors)
	suite.Equal("Firstname", responseBody.ValErrors[0].Field)
	suite.Equal("Should be at least 2 characters", responseBody.ValErrors[0].Message)
	suite.Equal("Lastname", responseBody.ValErrors[1].Field)
	suite.Equal("Should be less than 50 characters", responseBody.ValErrors[1].Message)
	suite.Equal("Email", responseBody.ValErrors[2].Field)
	suite.Equal("Invalid email format", responseBody.ValErrors[2].Message)
	suite.Equal("Phone", responseBody.ValErrors[3].Field)
	suite.Equal("Invalid phone format", responseBody.ValErrors[3].Message)
}

func (suite *userControllerSuite) TestDeleteUser_Positive() {
	id := primitive.NewObjectID()

	suite.svc.On("DeleteUser", id).Return(nil)

	request, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("%s/api/user/%s", suite.testServer.URL, id.Hex()),
		bytes.NewBuffer(nil),
	)
	suite.NoError(err, "no error when creating the request")

	response, err := http.DefaultClient.Do(request)
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	responseBody := model.AppResponse{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	suite.Equal(http.StatusOK, response.StatusCode)
	suite.Equal("User deleted successfully", responseBody.Message)
	suite.svc.AssertExpectations(suite.T())
}

func (suite *userControllerSuite) TestDeleteUser_InvalidID_Negative() {
	id := primitive.NewObjectID()

	request, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("%s/api/user/%s", suite.testServer.URL, id.Hex()+"bad"),
		bytes.NewBuffer(nil),
	)
	suite.NoError(err, "no error when creating the request")

	response, err := http.DefaultClient.Do(request)
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	responseBody := model.AppResponse{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	suite.Equal(http.StatusBadRequest, response.StatusCode)
	suite.Equal("invalid id", responseBody.Message)
	suite.Empty(responseBody.Data, "user should not be retrieved")
}

func TestUserController(t *testing.T) {
	gin.SetMode(gin.TestMode)
	suite.Run(t, new(userControllerSuite))
}
