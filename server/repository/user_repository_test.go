package repository

import (
	"testing"

	"github.com/alexandre-pinon/epic-road-trip/config"
	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/alexandre-pinon/epic-road-trip/utils"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepositorySuite struct {
	suite.Suite
	cfg             config.Config
	repo            UserRepository
	cleanupExecutor utils.DropCollectionExecutor
}

func (suite *userRepositorySuite) SetupTest() {
	db := config.ConnectDB(&suite.cfg)
	repo := NewUserRepository(db)

	suite.repo = repo

	suite.cleanupExecutor = utils.NewDropCollectionExecutor(db)
}

func (suite *userRepositorySuite) TearDownTest() {
	defer config.DisconnectDB(&suite.cfg, suite.cleanupExecutor.DB.Client())
	defer suite.cleanupExecutor.DropCollection([]string{"user"})
}

func (suite *userRepositorySuite) TestGetAllUsers_EmptySlice_Positive() {
	users, err := suite.repo.GetAllUsers()
	suite.NoError(err, "no error when get all users when the table is empty")
	suite.Equal(0, len(*users), "length of users should be 0, since it is empty slice")
	suite.Equal([]model.User(nil), *users, "users is an empty slice")
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

	_, err := suite.repo.CreateUser(&user)
	suite.NoError(err, "no error when create user with valid input")
	_, err = suite.repo.CreateUser(&user)
	suite.NoError(err, "no error when create user with valid input")
	_, err = suite.repo.CreateUser(&user)
	suite.NoError(err, "no error when create user with valid input")

	users, err := suite.repo.GetAllUsers()
	suite.NoError(err, "no error when get all users when the table is empty")
	suite.Equal(3, len(*users), "insert 3 records before the all data, so it should contain three users")
}

func (suite *userRepositorySuite) TestGetUserByID_NotFound_Negative() {
	id := primitive.NewObjectID()

	_, err := suite.repo.GetUserByID(id)
	suite.Error(err, "error not found")
	suite.Equal(mongo.ErrNoDocuments, err)
}

func (suite *userRepositorySuite) TestGetUserByID_Exists_Positive() {
	user := model.User{
		Firstname: "yoimiya",
		Lastname:  "naganohara",
		Email:     "yoimiya.naganohara@gmail.com",
		Password:  "12345678",
		Phone:     "+33612345678",
		Trips:     []*model.RoadTrip{},
	}

	id, err := suite.repo.CreateUser(&user)
	suite.NoError(err, "no error when create user with valid input")

	result, err := suite.repo.GetUserByID(id.InsertedID.(primitive.ObjectID))
	suite.NoError(err, "no error because user is found")
	suite.Equal(user.Firstname, (*result).Firstname, "should be equal between result and user")
	suite.Equal(user.Email, (*result).Email, "should be equal between result and user")
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

	_, err := suite.repo.CreateUser(&user)
	suite.NoError(err, "no error when create user with valid input")
}

func (suite *userRepositorySuite) TestCreateUser_NilPointer_Negative() {
	_, err := suite.repo.CreateUser(nil)
	suite.Error(err, "create error with nil input returns error")
}

func (suite *userRepositorySuite) TestCreateUser_EmptyFields_Positive() {
	var user model.User
	_, err := suite.repo.CreateUser(&user)
	suite.NoError(err, "no error when create user with empty fields")
}

func TestUserRepository(t *testing.T) {
	cfg := config.GetConfig(config.Test)
	suite.Run(t, &userRepositorySuite{cfg: *cfg})
}
