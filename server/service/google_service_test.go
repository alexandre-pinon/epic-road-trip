package service

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alexandre-pinon/epic-road-trip/config"
	"github.com/alexandre-pinon/epic-road-trip/mocks"
	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type googleServiceSuite struct {
	suite.Suite
	cfg  config.Config
	repo *mocks.GoogleRepository
	svc  GoogleService
}

func (suite *googleServiceSuite) SetupTest() {
	repo := new(mocks.GoogleRepository)
	svc := NewGoogleService(suite.cfg)

	suite.repo = repo
	suite.svc = svc
}

// func (suite *googleServiceSuite) TestEnjoyWithZeroResult() {
// 	activities := &model.GoogleZeroResult{
// 			Result: []interface{}{},
// 			Status: errors.New("ZERO RESULTS"),
// 	}

// 	result, err := suite.svc.Enjoy("test")
// 	suite.NoError(err, "no crashed")
// 	suite.Equal(activities, *result, "result and error are the same")
// 	suite.repo.AssertExpectations(suite.T())
// }

// func (suite *googleServiceSuite) TestEnjoyWithGoodAnswer() {

// 	activities := []model.Activity{
// 		{
// 			Place_id: "CHIJyoimiya",
// 			Name:     "yoimiya",
// 			Opening_hours: []*model.Opening_hours{
// 				{
// 					Open_now: true,
// 					Periods: []*model.Periods{
// 						{
// 							Open_periods: []*model.Open{
// 								{
// 									Open_day:  0,
// 									Open_time: "2020",
// 								},
// 								{
// 									Open_day:  0,
// 									Open_time: "2020",
// 								},
// 							},
// 							Close_periods: []*model.Close{
// 								{
// 									Close_day:  0,
// 									Close_time: "2020",
// 								},
// 								{
// 									Close_day:  0,
// 									Close_time: "2020",
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 			Address:            "8 rue de ta mère c'est la catin du coin",
// 			Types:              "solide froid boite",
// 			User_ratings_total: 2,
// 			Price_level:        2,
// 			Rating:             4.2,
// 		},
// 		{
// 			Place_id: "CHIJyoimiya",
// 			Name:     "yoimiya",
// 			Opening_hours: []*model.Opening_hours{
// 				{
// 					Open_now: true,
// 					Periods: []*model.Periods{
// 						{
// 							Open_periods: []*model.Open{
// 								{
// 									Open_day:  0,
// 									Open_time: "2020",
// 								},
// 								{
// 									Open_day:  0,
// 									Open_time: "2020",
// 								},
// 							},
// 							Close_periods: []*model.Close{
// 								{
// 									Close_day:  0,
// 									Close_time: "2020",
// 								},
// 								{
// 									Close_day:  0,
// 									Close_time: "2020",
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 			Address:            "8 rue de ta mère c'est la catin du coin",
// 			Types:              "solide froid boite",
// 			User_ratings_total: 2,
// 			Price_level:        2,
// 			Rating:             4.2,
// 		},
// 		{
// 			Place_id: "CHIJyoimiya",
// 			Name:     "yoimiya",
// 			Opening_hours: []*model.Opening_hours{
// 				{
// 					Open_now: true,
// 					Periods: []*model.Periods{
// 						{
// 							Open_periods: []*model.Open{
// 								{
// 									Open_day:  0,
// 									Open_time: "2020",
// 								},
// 								{
// 									Open_day:  0,
// 									Open_time: "2020",
// 								},
// 							},
// 							Close_periods: []*model.Close{
// 								{
// 									Close_day:  0,
// 									Close_time: "2020",
// 								},
// 								{
// 									Close_day:  0,
// 									Close_time: "2020",
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 			Types:              "solide froid boite",
// 			User_ratings_total: 2,
// 			Price_level:        2,
// 			Rating:             4.2,
// 		},
// 	}
// 	suite.repo.On("Enjoy", "test").Return(&activities, nil)
// 	result, err := suite.svc.Enjoy("test")
// 	suite.NoError(err, "no error when get all activities")
// 	suite.Equal(len(activities), len(*result), "activities and result should have the same length")
// 	suite.Equal(activities, *result, "result and activities are the same")
// 	suite.repo.AssertExpectations(suite.T())
// }

func (suite *googleServiceSuite) TestGeocoding_NoResult_Negative() {
	noResults := model.GoogleGeocodingResponse{
		Results: []model.GoogleGeocodingResult{},
		Status:  "ZERO RESULTS",
	}
	router := gin.Default()
	router.GET("/geocode/json", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, &noResults)
	})
	server := httptest.NewServer(router)

	result, err := suite.svc.GeoCoding(server.URL, "pancake")
	suite.Error(err, "error: no results")
	suite.Equal(http.StatusNotFound, err.(*model.AppError).StatusCode)
	suite.Equal(noResults.Status, err.Error())
	suite.Nil(result)
}

func (suite *googleServiceSuite) TestGeocoding_Results_Positive() {
	location := model.Location{Lat: 48.856614, Lng: 2.3522219}
	withResults := model.GoogleGeocodingResponse{
		Results: []model.GoogleGeocodingResult{{
			AddressComponents: []model.AddressComponent{{
				LongName:  "Paris",
				ShortName: "Paris",
				Types: []string{
					"locality",
					"political",
				}},
			},
			FormattedAddress: "Paris, France",
			Geometry: model.Geometry{
				Bounds: model.Bounds{
					Northeast: model.Location{Lat: 48.9021475, Lng: 2.4698509},
					Southwest: model.Location{Lat: 48.8155622, Lng: 2.2242191},
				},
				Location:     location,
				LocationType: "APPROXIMATE",
				Viewport: model.Bounds{
					Northeast: model.Location{Lat: 48.9021475, Lng: 2.4698509},
					Southwest: model.Location{Lat: 48.8155622, Lng: 2.2242191},
				},
			},
			PlaceID: "ChIJD7fiBh9u5kcRYJSMaMOCCwQ",
			Types: []string{
				"locality",
				"political",
			}},
		},
		Status: "OK",
	}
	router := gin.Default()
	router.GET("/geocode/json", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, &withResults)
	})
	server := httptest.NewServer(router)

	result, err := suite.svc.GeoCoding(server.URL, "Paris")
	suite.NoError(err, "no error if result")
	suite.Equal(location, *result)
}

func TestGoogleService(t *testing.T) {
	cfg := config.GetConfig()
	cfg.Env = config.Test
	suite.Run(t, new(googleServiceSuite))
}
