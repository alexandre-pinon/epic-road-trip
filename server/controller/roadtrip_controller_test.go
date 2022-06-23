package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alexandre-pinon/epic-road-trip/config"
	"github.com/alexandre-pinon/epic-road-trip/mocks"
	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/alexandre-pinon/epic-road-trip/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type roadtripControllerSuite struct {
	suite.Suite
	cfg           *config.Config
	googleService *mocks.GoogleService
	crtl          RoadTripController
	testServer    *httptest.Server
}

func (suite *roadtripControllerSuite) SetupTest() {
	googleService := new(mocks.GoogleService)
	crtl := NewRoadTripController(suite.cfg, googleService)

	router := gin.New()
	apiRoutes := router.Group("/api/v1")
	{
		roadtripRoutes := apiRoutes.Group("/roadtrip")
		{
			roadtripRoutes.POST("/enjoy", utils.ServeHTTP(crtl.Enjoy))
		}
	}
	server := httptest.NewServer(router)

	suite.testServer = server
	suite.googleService = googleService
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

	suite.googleService.On("GeoCoding", suite.cfg.BaseUrlGoogle, request.City).Return(&location, nil)
	suite.googleService.On("Enjoy", suite.cfg.BaseUrlGoogle, location).Return(&activities, nil)

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

	suite.googleService.On("GeoCoding", suite.cfg.BaseUrlGoogle, request.City).Return(&location, nil)
	suite.googleService.On("Enjoy", suite.cfg.BaseUrlGoogle, location).Return(nil, &noResult)

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

func TestGoogleController(t *testing.T) {
	gin.SetMode(gin.TestMode)
	cfg := config.GetConfig()
	cfg.Env = config.Test
	suite.Run(t, &roadtripControllerSuite{cfg: cfg})
}
