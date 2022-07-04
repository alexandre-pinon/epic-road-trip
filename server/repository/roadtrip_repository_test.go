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

type roadtripRepositorySuite struct {
	suite.Suite
	cfg             *config.Config
	repo            RoadtripRepository
	cleanupExecutor utils.DropCollectionExecutor
}

func (suite *roadtripRepositorySuite) SetupTest() {
	db := config.ConnectDB(suite.cfg)
	repo := NewRoadtripRepository(db)

	suite.repo = repo

	suite.cleanupExecutor = utils.NewDropCollectionExecutor(db)
}

func (suite *roadtripRepositorySuite) TearDownTest() {
	defer config.DisconnectDB(suite.cfg, suite.cleanupExecutor.DB.Client())
	defer suite.cleanupExecutor.DropCollection([]string{"roadtrip"})
}

func (suite *roadtripRepositorySuite) TestGetAllRoadtrips_EmptySlice_Positive() {
	roadtrips, err := suite.repo.GetAllRoadtrips()
	suite.NoError(err, "no error when get all roadtrips when the table is empty")
	suite.Equal(0, len(*roadtrips), "length of roadtrips should be 0, since it is empty slice")
	suite.Equal([]model.Roadtrip(nil), *roadtrips, "roadtrips is an empty slice")
}

func (suite *roadtripRepositorySuite) TestGetAllRoadtrips_FilledRecords_Positive() {
	insertRoadtrips := []model.Roadtrip{
		{
			Startdate:   time.Date(2022, 8, 5, 0, 0, 0, 0, time.UTC),
			Enddate:     time.Date(2022, 8, 21, 0, 0, 0, 0, time.UTC),
			TripStepsID: []primitive.ObjectID{primitive.NewObjectID(), primitive.NewObjectID()},
		},
		{
			Startdate:   time.Date(2022, 8, 5, 0, 0, 0, 0, time.UTC),
			Enddate:     time.Date(2022, 8, 21, 0, 0, 0, 0, time.UTC),
			TripStepsID: []primitive.ObjectID{primitive.NewObjectID(), primitive.NewObjectID()},
		},
		{
			Startdate:   time.Date(2022, 8, 5, 0, 0, 0, 0, time.UTC),
			Enddate:     time.Date(2022, 8, 21, 0, 0, 0, 0, time.UTC),
			TripStepsID: []primitive.ObjectID{primitive.NewObjectID(), primitive.NewObjectID()},
		},
	}

	for _, roadtrip := range insertRoadtrips {
		_, err := suite.repo.CreateRoadtrip(&roadtrip)
		suite.NoError(err, "no error when create roadtrip with valid input")
	}

	roadtrips, err := suite.repo.GetAllRoadtrips()
	suite.NoError(err, "no error when get all roadtrips when the table is empty")
	suite.Equal(3, len(*roadtrips), "insert 3 records before the all data, so it should contain three roadtrips")
}

func (suite *roadtripRepositorySuite) TestGetRoadtripByID_NotFound_Negative() {
	id := primitive.NewObjectID()

	_, err := suite.repo.GetRoadtripByID(id)
	suite.Error(err, "error not found")
	suite.Equal(mongo.ErrNoDocuments, err)
}

func (suite *roadtripRepositorySuite) TestGetRoadtripByID_Exists_Positive() {
	roadtrip := model.Roadtrip{
		Startdate:   time.Date(2022, 8, 5, 0, 0, 0, 0, time.UTC),
		Enddate:     time.Date(2022, 8, 21, 0, 0, 0, 0, time.UTC),
		TripStepsID: []primitive.ObjectID{primitive.NewObjectID(), primitive.NewObjectID()},
	}

	id, err := suite.repo.CreateRoadtrip(&roadtrip)
	suite.NoError(err, "no error when create roadtrip with valid input")

	result, err := suite.repo.GetRoadtripByID(id.InsertedID.(primitive.ObjectID))
	suite.NoError(err, "no error because roadtrip is found")
	suite.Equal(roadtrip.Startdate, (*result).Startdate, "should be equal between result and roadtrip")
	suite.Equal(roadtrip.TripStepsID, (*result).TripStepsID, "should be equal between result and roadtrip")
}

func (suite *roadtripRepositorySuite) TestCreateRoadtrip_Positive() {
	roadtrip := model.Roadtrip{
		Startdate:   time.Date(2022, 8, 5, 0, 0, 0, 0, time.UTC),
		Enddate:     time.Date(2022, 8, 21, 0, 0, 0, 0, time.UTC),
		TripStepsID: []primitive.ObjectID{primitive.NewObjectID(), primitive.NewObjectID()},
	}

	_, err := suite.repo.CreateRoadtrip(&roadtrip)
	suite.NoError(err, "no error when create roadtrip with valid input")
}

func (suite *roadtripRepositorySuite) TestCreateRoadtrip_EmptyFields_Positive() {
	var roadtrip model.Roadtrip
	_, err := suite.repo.CreateRoadtrip(&roadtrip)
	suite.NoError(err, "no error when create roadtrip with empty fields")
}

func (suite *roadtripRepositorySuite) TestDeleteRoadtrip_Positive() {
	roadtrip := model.Roadtrip{
		Startdate:   time.Date(2022, 8, 5, 0, 0, 0, 0, time.UTC),
		Enddate:     time.Date(2022, 8, 21, 0, 0, 0, 0, time.UTC),
		TripStepsID: []primitive.ObjectID{primitive.NewObjectID(), primitive.NewObjectID()},
	}

	id, err := suite.repo.CreateRoadtrip(&roadtrip)
	suite.NoError(err, "no error when create roadtrip with valid input")

	result, err := suite.repo.DeleteRoadtrip(id.InsertedID.(primitive.ObjectID))
	suite.NoError(err, "no error because roadtrip is found")
	suite.Require().NotNil(result)
	suite.Equal(int64(1), result.DeletedCount)
}

func TestRoadtripRepository(t *testing.T) {
	cfg := config.GetConfig()
	cfg.App.Env = config.Test
	suite.Run(t, &roadtripRepositorySuite{cfg: cfg})
}
