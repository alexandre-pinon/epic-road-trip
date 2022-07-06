package repository

import (
	"testing"
	"time"

	"github.com/alexandre-pinon/epic-road-trip/config"
	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/alexandre-pinon/epic-road-trip/utils"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepositorySuite struct {
	suite.Suite
	cfg             *config.Config
	userRepo        UserRepository
	tripStepRepo    TripStepRepository
	cleanupExecutor utils.DropCollectionExecutor
}

func (suite *userRepositorySuite) SetupTest() {
	db := config.ConnectDB(suite.cfg)
	userRepo := NewUserRepository(db)
	tripStepRepo := NewTripStepRepository(db)

	suite.userRepo = userRepo
	suite.tripStepRepo = tripStepRepo

	suite.cleanupExecutor = utils.NewDropCollectionExecutor(db)
}

func (suite *userRepositorySuite) TearDownTest() {
	defer config.DisconnectDB(suite.cfg, suite.cleanupExecutor.DB.Client())
	defer suite.cleanupExecutor.DropCollection([]string{"user"})
}

func (suite *userRepositorySuite) TestGetAllUsers_EmptySlice_Positive() {
	users, err := suite.userRepo.GetAllUsers()
	suite.NoError(err, "no error when get all users when the table is empty")
	suite.Equal(0, len(*users), "length of users should be 0, since it is empty slice")
	suite.Equal([]model.User(nil), *users, "users is an empty slice")
}

func (suite *userRepositorySuite) TestGetAllUsers_FilledRecords_Positive() {
	insertUsers := []model.User{
		{
			Firstname:      "yoimiya",
			Lastname:       "naganohara",
			Email:          "yoimiya.naganohara@gmail.com",
			HashedPassword: "12345678",
			Phone:          "+33612345678",
			Trips:          []*model.Roadtrip{},
		},
		{
			Firstname:      "hu",
			Lastname:       "tao",
			Email:          "hu.tao@gmail.com",
			HashedPassword: "23456789",
			Phone:          "+33623456789",
			Trips:          []*model.Roadtrip{},
		},
		{
			Firstname:      "kokomi",
			Lastname:       "sangonomiya",
			Email:          "kokomi.sangonomiya@gmail.com",
			HashedPassword: "87654321",
			Phone:          "+33687654321",
			Trips:          []*model.Roadtrip{},
		},
	}

	for _, user := range insertUsers {
		_, err := suite.userRepo.CreateUser(&user)
		suite.NoError(err, "no error when create user with valid input")
	}

	users, err := suite.userRepo.GetAllUsers()
	suite.NoError(err, "no error when get all users when the table is empty")
	suite.Equal(3, len(*users), "insert 3 records before the all data, so it should contain three users")
}

func (suite *userRepositorySuite) TestGetUserByID_NotFound_Negative() {
	id := primitive.NewObjectID()

	_, err := suite.userRepo.GetUserByID(id, false)
	suite.Error(err, "error not found")
	suite.Equal(mongo.ErrNoDocuments, err)
}

func (suite *userRepositorySuite) TestGetUserByID_Exists_Positive() {
	user := model.User{
		Firstname:      "yoimiya",
		Lastname:       "naganohara",
		Email:          "yoimiya.naganohara@gmail.com",
		HashedPassword: "12345678",
		Phone:          "+33612345678",
		Trips:          []*model.Roadtrip{},
	}

	id, err := suite.userRepo.CreateUser(&user)
	suite.NoError(err, "no error when create user with valid input")

	result, err := suite.userRepo.GetUserByID(id.InsertedID.(primitive.ObjectID), false)
	suite.NoError(err, "no error because user is found")
	suite.Equal(user.Firstname, (*result).Firstname, "should be equal between result and user")
	suite.Equal(user.Email, (*result).Email, "should be equal between result and user")
}

func (suite *userRepositorySuite) TestGetUserByID_Populate_Exists_Positive() {
	tripStep := model.TripStep{
		Startdate: time.Date(2022, 8, 5, 0, 0, 0, 0, time.UTC),
		Enddate:   time.Date(2022, 8, 21, 0, 0, 0, 0, time.UTC),
		City:      "Paris",
		Enjoy: &[]model.Enjoy{{
			Name:     "Hôtel de Ville",
			Rating:   4.4,
			Vicinity: "Place de l'Hôtel de Ville, Paris",
		}},
		Sleep: &[]model.Sleep{{
			Name:     "Britannique Hotel - Paris Centre",
			Rating:   4.7,
			Vicinity: "20 Avenue Victoria, Paris",
		}},
		Eat: &[]model.Eat{{
			Name:     "L'Art Brut Bistrot",
			Rating:   4.6,
			Vicinity: "78 Rue Quincampoix, Paris",
		}},
		Drink: &[]model.Drink{{
			Name:     "Hôtel Duo",
			Rating:   4.2,
			Vicinity: "11 Rue du Temple, Paris",
		}},
	}
	res, err := suite.tripStepRepo.CreateTripStep(&tripStep)
	suite.Require().NoError(err, "no error when create tripstep with valid input")

	user := model.User{
		Firstname:      "yoimiya",
		Lastname:       "naganohara",
		Email:          "yoimiya.naganohara@gmail.com",
		HashedPassword: "12345678",
		Phone:          "+33612345678",
		Trips: []*model.Roadtrip{{
			Startdate:   time.Date(2022, 12, 12, 0, 0, 0, 0, time.UTC),
			Enddate:     time.Date(2022, 12, 22, 0, 0, 0, 0, time.UTC),
			TripStepsID: []primitive.ObjectID{res.InsertedID.(primitive.ObjectID)},
		}},
	}

	id, err := suite.userRepo.CreateUser(&user)
	suite.Require().NoError(err, "no error when create user with valid input")

	result, err := suite.userRepo.GetUserByID(id.InsertedID.(primitive.ObjectID), true)
	suite.NoError(err, "no error because user is found")
	suite.Equal(user.Firstname, (*result).Firstname, "should be equal between result and user")
	suite.Equal(user.Email, (*result).Email, "should be equal between result and user")
	suite.NotEmpty((*result).Trips[0].TripSteps)
}

func (suite *userRepositorySuite) TestGetUserByEmail_NotFound_Negative() {
	_, err := suite.userRepo.GetUserByEmail("")
	suite.Error(err, "error not found")
	suite.Equal(mongo.ErrNoDocuments, err)
}

func (suite *userRepositorySuite) TestGetUserByEmail_Exists_Positive() {
	email := "yoimiya.naganohara@gmail.com"
	user := model.User{
		Firstname:      "yoimiya",
		Lastname:       "naganohara",
		Email:          email,
		HashedPassword: "12345678",
		Phone:          "+33612345678",
		Trips:          []*model.Roadtrip{},
	}

	_, err := suite.userRepo.CreateUser(&user)
	suite.NoError(err, "no error when create user with valid input")

	result, err := suite.userRepo.GetUserByEmail(email)
	suite.NoError(err, "no error because user is found")
	suite.Equal(user.Phone, (*result).Phone, "should be equal between result and user")
	suite.Equal(user.Lastname, (*result).Lastname, "should be equal between result and user")
}

func (suite *userRepositorySuite) TestCreateUser_Positive() {
	user := model.User{
		Firstname:      "yoimiya",
		Lastname:       "naganohara",
		Email:          "yoimiya.naganohara@gmail.com",
		HashedPassword: "12345678",
		Phone:          "+33612345678",
		Trips:          []*model.Roadtrip{},
	}

	_, err := suite.userRepo.CreateUser(&user)
	suite.NoError(err, "no error when create user with valid input")
}

func (suite *userRepositorySuite) TestCreateUser_EmptyFields_Positive() {
	var user model.User
	_, err := suite.userRepo.CreateUser(&user)
	suite.NoError(err, "no error when create user with empty fields")
}

func (suite *userRepositorySuite) TestUpdateUser_Positive() {
	user := model.User{
		Firstname:      "yoimiya",
		Lastname:       "naganohara",
		Email:          "yoimiya.naganohara@gmail.com",
		HashedPassword: "12345678",
		Phone:          "+33612345678",
		Trips:          []*model.Roadtrip{},
	}

	createResult, err := suite.userRepo.CreateUser(&user)
	suite.NoError(err, "no error when create user with valid input")

	id := createResult.InsertedID.(primitive.ObjectID)
	user.Firstname = "yoiyoiyoimiya"
	user.Phone = "+33712345678"

	updateResult, err := suite.userRepo.UpdateUser(id, &user)
	suite.Require().NoError(err, "no error when update user with valid input")

	suite.Equal(int64(1), updateResult.ModifiedCount)

	userResult, err := suite.userRepo.GetUserByID(id, false)
	suite.NoError(err, "no error because user is found")
	suite.Equal("yoiyoiyoimiya", (*userResult).Firstname, "should be equal between result and user")
	suite.Equal("+33712345678", (*userResult).Phone, "should be equal between result and user")
}

func (suite *userRepositorySuite) TestDeleteUser_Positive() {
	user := model.User{
		Firstname:      "yoimiya",
		Lastname:       "naganohara",
		Email:          "yoimiya.naganohara@gmail.com",
		HashedPassword: "12345678",
		Phone:          "+33612345678",
		Trips:          []*model.Roadtrip{},
	}

	id, err := suite.userRepo.CreateUser(&user)
	suite.NoError(err, "no error when create user with valid input")

	result, err := suite.userRepo.DeleteUser(id.InsertedID.(primitive.ObjectID))
	suite.NoError(err, "no error because user is found")
	suite.Require().NotNil(result)
	suite.Equal(int64(1), result.DeletedCount)
}

func TestUserRepository(t *testing.T) {
	cfg := config.GetConfig(string(config.Test))
	suite.Run(t, &userRepositorySuite{cfg: cfg})
}
