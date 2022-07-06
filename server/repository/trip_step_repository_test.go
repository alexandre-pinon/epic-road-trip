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

type tripStepRepositorySuite struct {
	suite.Suite
	cfg             *config.Config
	repo            TripStepRepository
	cleanupExecutor utils.DropCollectionExecutor
}

func (suite *tripStepRepositorySuite) SetupTest() {
	db := config.ConnectDB(suite.cfg)
	repo := NewTripStepRepository(db)

	suite.repo = repo

	suite.cleanupExecutor = utils.NewDropCollectionExecutor(db)
}

func (suite *tripStepRepositorySuite) TearDownTest() {
	defer config.DisconnectDB(suite.cfg, suite.cleanupExecutor.DB.Client())
	defer suite.cleanupExecutor.DropCollection([]string{"tripStep"})
}

func (suite *tripStepRepositorySuite) TestGetAllTripSteps_EmptySlice_Positive() {
	tripSteps, err := suite.repo.GetAllTripSteps()
	suite.NoError(err, "no error when get all tripSteps when the table is empty")
	suite.Equal(0, len(*tripSteps), "length of tripSteps should be 0, since it is empty slice")
	suite.Equal([]model.TripStep(nil), *tripSteps, "tripSteps is an empty slice")
}

func (suite *tripStepRepositorySuite) TestGetAllTripSteps_FilledRecords_Positive() {
	insertTripSteps := []model.TripStep{
		{
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
		},
		{
			Startdate: time.Date(2022, 8, 5, 0, 0, 0, 0, time.UTC),
			Enddate:   time.Date(2022, 8, 21, 0, 0, 0, 0, time.UTC),
			City:      "London",
			Travel: &model.Itinerary{
				Type: model.Ground,
				Departure: model.Station{
					Name:    "Paris, France",
					City:    "Paris",
					Country: "France",
				},
				Arrival: model.Station{
					Name:    "London, England",
					City:    "London",
					Country: "England",
				},
				Duration:       9 * time.Hour,
				DurationString: (9 * time.Hour).String(),
				Startdate:      time.Date(2022, 12, 12, 12, 0, 0, 0, time.UTC),
				Enddate:        time.Date(2022, 12, 13, 2, 0, 0, 0, time.UTC),
				Steps: []model.ItineraryStep{{
					Type:           "Train",
					Departure:      "Montparnasse",
					Arrival:        "Gare de Hendaye",
					Duration:       4*time.Hour + 36*time.Hour,
					DurationString: (4*time.Hour + 36*time.Hour).String(),
					Startdate:      time.Date(2022, 12, 12, 12, 23, 0, 0, time.UTC),
					Enddate:        time.Date(2022, 12, 12, 16, 59, 0, 0, time.UTC),
				}, {
					Type:           "Bus",
					Departure:      "Hendaye",
					Arrival:        "Bilbao (Bus Station)",
					Duration:       2 * time.Hour,
					DurationString: (2 * time.Hour).String(),
					Startdate:      time.Date(2022, 12, 12, 17, 14, 0, 0, time.UTC),
					Enddate:        time.Date(2022, 12, 12, 19, 14, 0, 0, time.UTC),
				}},
			},
			Enjoy: &[]model.Enjoy{},
			Sleep: &[]model.Sleep{},
			Eat:   &[]model.Eat{},
			Drink: &[]model.Drink{},
		},
		{
			Startdate: time.Date(2022, 8, 5, 0, 0, 0, 0, time.UTC),
			Enddate:   time.Date(2022, 8, 21, 0, 0, 0, 0, time.UTC),
			City:      "Tokyo",
			Travel: &model.Itinerary{
				Type: model.Air,
				Departure: model.Station{
					Name:    "LHR",
					City:    "London",
					Country: "EN",
				},
				Arrival: model.Station{
					Name:    "HND",
					City:    "Tokyo",
					Country: "JP",
				},
				Duration:       10 * time.Hour,
				DurationString: (10 * time.Hour).String(),
				Startdate:      time.Date(2022, 12, 12, 14, 0, 0, 0, time.UTC),
				Enddate:        time.Date(2022, 12, 13, 8, 0, 0, 0, time.UTC),
				Price:          999.99,
			},
			Enjoy: &[]model.Enjoy{},
			Sleep: &[]model.Sleep{},
			Eat:   &[]model.Eat{},
			Drink: &[]model.Drink{},
		},
	}

	for _, tripStep := range insertTripSteps {
		_, err := suite.repo.CreateTripStep(&tripStep)
		suite.NoError(err, "no error when create tripStep with valid input")
	}

	tripSteps, err := suite.repo.GetAllTripSteps()
	suite.NoError(err, "no error when get all tripSteps when the table is empty")
	suite.Equal(3, len(*tripSteps), "insert 3 records before the all data, so it should contain three tripSteps")
}

func (suite *tripStepRepositorySuite) TestGetTripStepByID_NotFound_Negative() {
	id := primitive.NewObjectID()

	_, err := suite.repo.GetTripStepByID(id)
	suite.Error(err, "error not found")
	suite.Equal(mongo.ErrNoDocuments, err)
}

func (suite *tripStepRepositorySuite) TestGetTripStepByID_Exists_Positive() {
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

	id, err := suite.repo.CreateTripStep(&tripStep)
	suite.NoError(err, "no error when create tripStep with valid input")

	result, err := suite.repo.GetTripStepByID(id.InsertedID.(primitive.ObjectID))
	suite.NoError(err, "no error because tripStep is found")
	suite.Equal(tripStep.Startdate, (*result).Startdate, "should be equal between result and tripStep")
	suite.Equal(tripStep.Enjoy, (*result).Enjoy, "should be equal between result and tripStep")
	suite.Equal(tripStep.Drink, (*result).Drink, "should be equal between result and tripStep")
}

func (suite *tripStepRepositorySuite) TestCreateTripStep_Positive() {
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

	_, err := suite.repo.CreateTripStep(&tripStep)
	suite.NoError(err, "no error when create tripStep with valid input")
}

func (suite *tripStepRepositorySuite) TestCreateTripStep_EmptyFields_Positive() {
	var tripStep model.TripStep
	_, err := suite.repo.CreateTripStep(&tripStep)
	suite.NoError(err, "no error when create tripStep with empty fields")
}

func (suite *tripStepRepositorySuite) TestDeleteTripStep_Positive() {
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

	id, err := suite.repo.CreateTripStep(&tripStep)
	suite.NoError(err, "no error when create tripStep with valid input")

	result, err := suite.repo.DeleteTripStep(id.InsertedID.(primitive.ObjectID))
	suite.NoError(err, "no error because tripStep is found")
	suite.Require().NotNil(result)
	suite.Equal(int64(1), result.DeletedCount)
}

func TestTripStepRepository(t *testing.T) {
	cfg := config.GetConfig(string(config.Test))
	suite.Run(t, &tripStepRepositorySuite{cfg: cfg})
}
