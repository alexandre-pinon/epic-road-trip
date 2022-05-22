package service

import (
	"net/http"
	"testing"

	"github.com/alexandre-pinon/epic-road-trip/mocks"
	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/stretchr/testify/suite"
)

type userServiceSuite struct {
	suite.Suite
	repo    *mocks.UserRepository
	service UserService
}

func (suite *userServiceSuite) SetupTest() {
	repo := new(mocks.UserRepository)
	service := NewUserService(repo)

	suite.repo = repo
	suite.service = service
}

func (suite *userServiceSuite) TestCreateUser_Positive() {
	user := model.User{
		Firstname: "yoimiya",
		Lastname:  "naganohara",
		Email:     "yoimiya.naganohara@gmail.com",
		Password:  "12345678",
		Phone:     "+33612345678",
		Trips:     []*model.RoadTrip{},
	}

	suite.repo.On("CreateUser", &user).Return(nil)

	err := suite.service.CreateUser(&user)
	suite.Nil(err, "err is a nil pointer so no error in this process")
	suite.repo.AssertExpectations(suite.T())
}

func (suite *userServiceSuite) TestCreateUser_NilPointer_Negative() {
	err := suite.service.CreateUser(nil)
	suite.Error(err.(*model.AppError).Err, "error when create tweet with nil pointer")
	suite.Assertions.Equal(err.(*model.AppError).StatusCode, http.StatusInternalServerError)
	suite.repo.AssertExpectations(suite.T())
}

func TestUserService(t *testing.T) {
	suite.Run(t, new(userServiceSuite))
}
