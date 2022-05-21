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
	cfg             config.Config
	repository      UserRepository
	cleanupExecutor utils.DropCollectionExecutor
}

func (suite *userRepositorySuite) SetupTest() {
	db := config.ConnectDB(&suite.cfg)
	repository := NewUserRepository(db)

	suite.repository = repository

	suite.cleanupExecutor = utils.NewDropCollectionExecutor(db)
}

func (suite *userRepositorySuite) TearDownTest() {
	defer config.DisconnectDB(&suite.cfg, suite.cleanupExecutor.DB.Client())
	defer suite.cleanupExecutor.DropCollection([]string{"user"})
}

func (suite *userRepositorySuite) TestGetAllUsers_EmptySlice_Positive() {
	users, err := suite.repository.GetAllUsers()
	suite.NoError(err, "no error when get all users when the table is empty")
	suite.Equal(len(*users), 0, "length of users should be 0, since it is empty slice")
	suite.Equal(*users, []model.User(nil), "users is an empty slice")
}

func (suite *userRepositorySuite) TestGetAllUsers_FilledRecords_Positive() {
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
	err = suite.repository.CreateUser(&user)
	suite.NoError(err, "no error when create user with valid input")
	err = suite.repository.CreateUser(&user)
	suite.NoError(err, "no error when create user with valid input")

	users, err := suite.repository.GetAllUsers()
	suite.NoError(err, "no error when get all users when the table is empty")
	suite.Equal(len(*users), 3, "insert 3 records before the all data, so it should contain three users")
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
	cfg := config.GetConfig(config.Test)
	suite.Run(t, &userRepositorySuite{cfg: *cfg})
}
