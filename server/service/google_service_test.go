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
	cfg  *config.Config
	repo *mocks.GoogleRepository
	svc  GoogleService
}

func (suite *googleServiceSuite) SetupTest() {
	repo := new(mocks.GoogleRepository)
	svc := NewGoogleService(suite.cfg)

	suite.repo = repo
	suite.svc = svc
}

func (suite *googleServiceSuite) TestEnjoyWithZeroResult() {
	noResult := model.Activity{
		HTMLAttributions: []interface{}{},
		Results:          []model.ActivityResult{},
		Status:           "ZERO_RESULTS",
	}

	router := gin.Default()
	router.GET("/place/nearbysearch/json", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, &noResult)
	})
	server := httptest.NewServer(router)

	result, err := suite.svc.Enjoy(server.URL, model.Location{Lat: 48.856614, Lng: 2.3522219})
	suite.Error(err, "error: no results")
	suite.Equal(http.StatusNotFound, err.(*model.AppError).StatusCode)
	suite.Equal(noResult.Status, err.Error())
	suite.Nil(result)
}

func (suite *googleServiceSuite) TestEnjoyWithGoodAnswer() {
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
	withResults := model.Activity{
		HTMLAttributions: []interface{}{},
		Results:          activities,
		Status:           "no error",
	}
	router := gin.Default()
	router.GET("/place/nearbysearch/json", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, &withResults)
	})
	server := httptest.NewServer(router)
	result, err := suite.svc.Enjoy(server.URL, location)
	suite.NoError(err, "no crashed")
	suite.Equal(activities, *result, "result and error are the same")
}

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

	result, err := suite.svc.GeoCoding(server.URL, "paris")
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
	suite.Run(t, &googleServiceSuite{cfg: cfg})
}
