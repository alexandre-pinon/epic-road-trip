package service

import (
	"net/http"
	"testing"

	"github.com/alexandre-pinon/epic-road-trip/mocks"
	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userServiceSuite struct {
	suite.Suite
	repo *mocks.UserRepository
	svc  UserService
}

func (suite *userServiceSuite) SetupTest() {
	repo := new(mocks.UserRepository)
	svc := NewUserService(repo)

	suite.repo = repo
	suite.svc = svc
}

func (suite *userServiceSuite) TestGetAllUsers_EmptySlice_Positive() {
	emptyUsers := []model.User(nil)
	suite.repo.On("GetAllUsers").Return(&emptyUsers, nil)
	users, err := suite.svc.GetAllUsers()
	suite.NoError(err, "no error when get all users")
	suite.Equal(0, len(*users), "users is a empty slice object")
}

func (suite *userServiceSuite) TestGetAllUsers_FilledSlice_Positive() {
	users := []model.User{
		{
			Firstname: "yoimiya",
			Lastname:  "naganohara",
			Email:     "yoimiya.naganohara@gmail.com",
			Password:  "12345678",
			Phone:     "+33612345678",
			Trips:     []*model.RoadTrip{},
		},
		{
			Firstname: "hu",
			Lastname:  "tao",
			Email:     "hu.tao@gmail.com",
			Password:  "23456789",
			Phone:     "+33623456789",
			Trips:     []*model.RoadTrip{},
		},
		{
			Firstname: "kokomi",
			Lastname:  "sangonomiya",
			Email:     "kokomi.sangonomiya@gmail.com",
			Password:  "87654321",
			Phone:     "+33687654321",
			Trips:     []*model.RoadTrip{},
		},
	}
	suite.repo.On("GetAllUsers").Return(&users, nil)
	result, err := suite.svc.GetAllUsers()
	suite.NoError(err, "no error when get all users")
	suite.Equal(len(users), len(*result), "users and result should have the same length")
	suite.Equal(users, *result, "result and users are the same")
}

func (suite *userServiceSuite) TestGetUserByID_NotFound_Negative() {
	id := primitive.NewObjectID()

	suite.repo.On("GetUserByID", id).Return(nil, mongo.ErrNoDocuments)

	result, err := suite.svc.GetUserByID(id)
	suite.Nil(result, "error is returned so result has to be nil")
	suite.Error(err, "error not found")
	suite.Equal("user not found", err.Error())
	suite.Equal(http.StatusNotFound, err.(*model.AppError).StatusCode)
	suite.repo.AssertExpectations(suite.T())
}

func (suite *userServiceSuite) TestGetUserByID_Exists_Positive() {
	id := primitive.NewObjectID()
	user := model.User{
		Firstname: "yoimiya",
		Lastname:  "naganohara",
		Email:     "yoimiya.naganohara@gmail.com",
		Password:  "12345678",
		Phone:     "+33612345678",
		Trips:     []*model.RoadTrip{},
	}

	suite.repo.On("GetUserByID", id).Return(&user, nil)

	result, err := suite.svc.GetUserByID(id)
	suite.Nil(err, "no error when return the user")
	suite.Equal(user, *result, "result and user should be equal")
}

func (suite *userServiceSuite) TestCreateUser_Positive() {
	id := &mongo.InsertOneResult{}
	user := model.User{
		Firstname: "yoimiya",
		Lastname:  "naganohara",
		Email:     "yoimiya.naganohara@gmail.com",
		Password:  "12345678",
		Phone:     "+33612345678",
		Trips:     []*model.RoadTrip{},
	}

	suite.repo.On("CreateUser", &user).Return(id, nil)

	err := suite.svc.CreateUser(&user)
	suite.NoError(err, "no error when create user with valid input")
	suite.repo.AssertExpectations(suite.T())
}

func (suite *userServiceSuite) TestCreateUser_NilPointer_Negative() {
	err := suite.svc.CreateUser(nil)
	suite.Error(err.(*model.AppError).Err, "error when create user with nil pointer")
	suite.Assertions.Equal(http.StatusInternalServerError, err.(*model.AppError).StatusCode)
	suite.repo.AssertExpectations(suite.T())
}

func TestUserService(t *testing.T) {
	suite.Run(t, new(userServiceSuite))
}
