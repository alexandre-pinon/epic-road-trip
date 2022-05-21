package repository

import (
	"testing"

	"github.com/alexandre-pinon/epic-road-trip/config"
	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/alexandre-pinon/epic-road-trip/utils"
	"github.com/stretchr/testify/suite"
)

type userRepositorySuite struct {
	suite.Suite
	repository      UserRepository
	cleanupExecutor utils.DropCollectionExecutor
}

func (suite *userRepositorySuite) SetupTest() {
	configs := config.GetConfig(config.Test)
	db := config.ConnectDB(configs)
	repository := NewUserRepository(db)

	suite.repository = repository

	suite.cleanupExecutor = utils.NewDropCollectionExecutor(db)
}

func (suite *userRepositorySuite) TearDownTest() {
	defer config.DisconnectDB(suite.cleanupExecutor.DB.Client())
	defer suite.cleanupExecutor.DropCollection([]string{"user"})
}

func (suite *userRepositorySuite) TestCreateUser_Positive() {
	user := model.User{
		Firstname: "yoimiya",
		Lastname:  "naganohara",
		Email:     "yoimiya.naganohara@gmail.com",
		Password:  "12345678",
		Phone:     "+33612345678",
		Trips:     []*model.RoadTrip{},
	}

	err := suite.repository.CreateUser(&user)
	suite.NoError(err, "no error when create user with valid input")
}

func (suite *userRepositorySuite) TestCreateUser_NilPointer_Negative() {
	err := suite.repository.CreateUser(nil)
	suite.Error(err, "create error with nil input returns error")
}

func (suite *userRepositorySuite) TestCreateUser_EmptyFields_Positive() {
	var user model.User
	err := suite.repository.CreateUser(&user)
	suite.NoError(err, "no error when create user with empty fields")
}

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(userRepositorySuite))
}