package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alexandre-pinon/epic-road-trip/config"
	"github.com/alexandre-pinon/epic-road-trip/mocks"
	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/alexandre-pinon/epic-road-trip/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type roadtripControllerSuite struct {
	suite.Suite
	cfg            *config.Config
	googleService  *mocks.GoogleService
	amadeusService *mocks.AmadeusService
	crtl           RoadTripController
	testServer     *httptest.Server
}

func (suite *roadtripControllerSuite) SetupTest() {
	googleService := new(mocks.GoogleService)
	amadeusService := new(mocks.AmadeusService)
	crtl := NewRoadTripController(suite.cfg, googleService, amadeusService)

	router := gin.New()
	apiRoutes := router.Group("/api/v1")
	{
		roadtripRoutes := apiRoutes.Group("/roadtrip")
		{
			roadtripRoutes.POST("/enjoy", utils.ServeHTTP(crtl.Enjoy))
			roadtripRoutes.POST("/travel", utils.ServeHTTP(crtl.Travel))
		}
	}
	server := httptest.NewServer(router)

	suite.testServer = server
	suite.googleService = googleService
	suite.amadeusService = amadeusService
	suite.crtl = crtl
}

func (suite *roadtripControllerSuite) TestEnjoyWithGoodAnswer() {
	request := model.CityFormData{
		City: "Paris",
	}
	location := model.Location{Lat: 48.856614, Lng: 2.3522219}
	activities := []model.ActivityResult{
		{
			BusinessStatus: "OPEN",
			Geometry: model.GeometryActivity{
				Location: location,
				Viewport: model.Bounds{
					Northeast: model.Location{Lat: 48.9021475, Lng: 2.4698509},
					Southwest: model.Location{Lat: 48.8155622, Lng: 2.2242191},
				},
			},
			Icon:                "ucfytc",
			IconBackgroundColor: "cbdosucb",
			IconMaskBaseURI:     "ubcs",
			Name:                "eonvfe",
			OpeningHours: model.OpeningHours{
				OpenNow: false,
			},
			Photos: []model.Photos{{
				Height:           800,
				HTMLAttributions: []string{"iubvd", "givuefbv"},
				PhotoReference:   "dfvdvfd",
				Width:            800,
			}},
			PlaceID: "CHIJoiubcfmigf1",
			PlusCode: model.PlusCode{
				CompoundCode: "",
				GlobalCode:   "",
			},
			PriceLevel:       2,
			Rating:           4.7695,
			Reference:        "",
			Scope:            "",
			Types:            []string{"amusement_park"},
			UserRatingsTotal: 5,
			Vicinity:         "",
		},
		{
			BusinessStatus: "OPEN",
			Geometry: model.GeometryActivity{
				Location: location,
				Viewport: model.Bounds{
					Northeast: model.Location{Lat: 48.9021475, Lng: 2.4698509},
					Southwest: model.Location{Lat: 48.8155622, Lng: 2.2242191},
				},
			},
			Icon:                "ucfytc",
			IconBackgroundColor: "cbdosucb",
			IconMaskBaseURI:     "ubcs",
			Name:                "eonvfe",
			OpeningHours: model.OpeningHours{
				OpenNow: false,
			},
			Photos: []model.Photos{{
				Height:           800,
				HTMLAttributions: []string{"iubvd", "givuefbv"},
				PhotoReference:   "dfvdvfd",
				Width:            800,
			}},
			PlaceID: "CHIJoiubcfmigf1",
			PlusCode: model.PlusCode{
				CompoundCode: "",
				GlobalCode:   "",
			},
			PriceLevel:       2,
			Rating:           4.7695,
			Reference:        "",
			Scope:            "",
			Types:            []string{"amusement_park"},
			UserRatingsTotal: 5,
			Vicinity:         "",
		},
	}

	suite.googleService.On("GeoCoding", suite.cfg.Google.BaseUrl, request.City).Return(&location, nil)
	suite.googleService.On("Enjoy", suite.cfg.Google.BaseUrl, location).Return(&activities, nil)

	requestBody, err := json.Marshal(request)
	if err != nil {
		suite.Nil(err, "json is empty")
	}

	response, err := http.Post(
		fmt.Sprintf("%s/api/v1/roadtrip/enjoy", suite.testServer.URL),
		gin.MIMEJSON,
		bytes.NewBuffer(requestBody),
	)

	suite.NoError(err, "no error when calling this endpoint")
	defer response.Body.Close()

	responseBody := model.AppResponse{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	suite.Equal(http.StatusOK, response.StatusCode)
	suite.Equal("Activities retrieved successfuly", responseBody.Message)
	suite.NotEmpty(responseBody.Data, "activities should be retrieved")
	suite.googleService.AssertExpectations(suite.T())
}

func (suite *roadtripControllerSuite) TestEnjoyWithZeroResult() {
	request := model.CityFormData{
		City: "Paris",
	}
	location := model.Location{Lat: 48.856614, Lng: 2.3522219}
	noResult := model.AppError{
		StatusCode: http.StatusNotFound,
		Err:        errors.New("ZERO_RESULTS"),
	}

	suite.googleService.On("GeoCoding", suite.cfg.Google.BaseUrl, request.City).Return(&location, nil)
	suite.googleService.On("Enjoy", suite.cfg.Google.BaseUrl, location).Return(nil, &noResult)

	requestBody, err := json.Marshal(request)
	if err != nil {
		suite.Nil(err, "json is emty")
	}
	response, err := http.Post(
		fmt.Sprintf("%s/api/v1/roadtrip/enjoy", suite.testServer.URL),
		gin.MIMEJSON,
		bytes.NewBuffer(requestBody),
	)

	suite.NoError(err, "no error when calling this endpoint")
	defer response.Body.Close()

	responseBody := model.AppResponse{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	suite.Equal(http.StatusNotFound, response.StatusCode)
	suite.Equal("ZERO_RESULTS", responseBody.Message)
	suite.Empty(responseBody.Data, "activities should not be retrieved")
	suite.googleService.AssertExpectations(suite.T())

}

func (suite *roadtripControllerSuite) TestTravel_Positive() {
	flighFormData := model.FlightFormData{
		OriginLocation:          "Paris",
		DestinationLocation:     "Tokyo",
		OriginLocationCode:      "PAR",
		DestinationLocationCode: "TYO",
		DepartureDate:           time.Date(2022, 12, 12, 12, 12, 12, 12, time.UTC),
		Adults:                  2,
	}
	now := time.Now().Unix()
	accessToken := model.AccessToken{
		Value: "kzijAxGphgIbhSm3WGIhmWvhep2G",
		Iat:   now,
		Exp:   int(now) + 30*60,
	}
	itineraries := []model.Itinerary{{
		Type: model.Air,
		Departure: model.Station{
			Name:    "CDG",
			City:    "Paris",
			Country: "FR",
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
	}}

	suite.amadeusService.On("GetAccessToken", suite.cfg.Amadeus.BaseUrl).Return(&accessToken, nil)
	suite.amadeusService.On("GetFlightOffers", suite.cfg.Amadeus.BaseUrl, accessToken.Value, &flighFormData).Return(&itineraries, nil)

	requestBody, err := json.Marshal(&flighFormData)
	suite.NoError(err, "can not marshal struct to json")

	response, err := http.Post(
		fmt.Sprintf("%s/api/v1/roadtrip/travel", suite.testServer.URL),
		gin.MIMEJSON,
		bytes.NewBuffer(requestBody),
	)
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	responseBody := model.AppResponse{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	suite.Equal(http.StatusOK, response.StatusCode)
	suite.Equal("Itineraries retrieved successfully", responseBody.Message)
	suite.NotEmpty(&responseBody.Data, "itineraries should be retrieved")
	suite.amadeusService.AssertExpectations(suite.T())
}

func (suite *roadtripControllerSuite) TestTravel_NoResults_Negative() {
	flighFormData := model.FlightFormData{
		OriginLocation:          "Paris",
		DestinationLocation:     "Sydney",
		OriginLocationCode:      "PAR",
		DestinationLocationCode: "SYD",
		DepartureDate:           time.Date(2022, 12, 12, 12, 12, 12, 12, time.UTC),
		Adults:                  2,
	}
	now := time.Now().Unix()
	accessToken := model.AccessToken{
		Value: "kzijAxGphgIbhSm3WGIhmWvhep2G",
		Iat:   now,
		Exp:   int(now) + 30*60,
	}
	noResults := model.AppError{
		StatusCode: http.StatusNotFound,
		Err:        errors.New("flight offers not found"),
	}

	suite.amadeusService.On("GetAccessToken", suite.cfg.Amadeus.BaseUrl).Return(&accessToken, nil)
	suite.amadeusService.On("GetFlightOffers", suite.cfg.Amadeus.BaseUrl, accessToken.Value, &flighFormData).Return(nil, &noResults)

	requestBody, err := json.Marshal(&flighFormData)
	suite.NoError(err, "can not marshal struct to json")

	response, err := http.Post(
		fmt.Sprintf("%s/api/v1/roadtrip/travel", suite.testServer.URL),
		gin.MIMEJSON,
		bytes.NewBuffer(requestBody),
	)
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	responseBody := model.AppResponse{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	suite.Equal(http.StatusNotFound, response.StatusCode)
	suite.Equal("flight offers not found", responseBody.Message)
	suite.Empty(&responseBody.Data, "itineraries should not be retrieved")
	suite.amadeusService.AssertExpectations(suite.T())
}

func (suite *roadtripControllerSuite) TestTravel_InvalidOriginDestination_Negative() {
	flighFormData := model.FlightFormData{
		OriginLocation:      "gerkj",
		DestinationLocation: "zofijzr",
		DepartureDate:       time.Date(2022, 12, 12, 12, 12, 12, 12, time.UTC),
		Adults:              2,
	}

	requestBody, err := json.Marshal(&flighFormData)
	suite.NoError(err, "can not marshal struct to json")

	response, err := http.Post(
		fmt.Sprintf("%s/api/v1/roadtrip/travel", suite.testServer.URL),
		gin.MIMEJSON,
		bytes.NewBuffer(requestBody),
	)
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	responseBody := model.AppResponse{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	suite.Equal(http.StatusNotFound, response.StatusCode)
	suite.Equal("no airport found for origin/destination cities", responseBody.Message)
	suite.Empty(&responseBody.Data, "itineraries should not be retrieved")
}

func TestRoadtripController(t *testing.T) {
	gin.SetMode(gin.TestMode)
	cfg := config.GetConfig()
	cfg.App.Env = config.Test
	suite.Run(t, &roadtripControllerSuite{cfg: cfg})
}
