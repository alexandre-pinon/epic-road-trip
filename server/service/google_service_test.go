package service

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

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

	params := model.Constraints{
		Radius:   0,
		MaxPrice: 500,
		MinPrice: 0,
		OpenNow:  false,
	}

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

	result, err := suite.svc.Enjoy(server.URL, model.Location{Lat: 48.856614, Lng: 2.3522219}, params)
	suite.Error(err, "error: no results")
	suite.Equal(http.StatusNotFound, err.(*model.AppError).StatusCode)
	suite.Equal(noResult.Status, err.Error())
	suite.Nil(result)
}

func (suite *googleServiceSuite) TestEnjoyWithGoodAnswer() {
	params := model.Constraints{
		Radius:   0,
		MaxPrice: 500,
		MinPrice: 0,
		OpenNow:  false,
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
	result, err := suite.svc.Enjoy(server.URL, location, params)
	suite.NoError(err, "no crashed")
	suite.Equal(activities, *result, "result and error are the same")
}

func (suite *googleServiceSuite) TestSleepWithGoodAnswer() {
	params := model.Constraints{
		Radius:   0,
		MaxPrice: 500,
		MinPrice: 0,
		OpenNow:  false,
	}
	location := model.Location{Lat: 48.856614, Lng: 2.3522219}
	hotels := []model.ActivityResult{
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
			Types:            []string{"lodging"},
			UserRatingsTotal: 5,
			Vicinity:         "",
		},
	}
	withResults := model.Hotel{
		HTMLAttributions: []interface{}{},
		NextPageToken:    "nvoinrvo",
		Results:          hotels,
		Status:           "no error",
	}
	router := gin.Default()
	router.GET("/place/nearbysearch/json", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, &withResults)
	})
	server := httptest.NewServer(router)
	result, err := suite.svc.Sleep(server.URL, location, params)
	suite.NoError(err, "no crashed")
	suite.Equal(hotels, *result, "result and error are the same")
}

func (suite *googleServiceSuite) TestSleepWithZeroResult() {
	params := model.Constraints{
		Radius:   0,
		MaxPrice: 500,
		MinPrice: 0,
		OpenNow:  false,
	}
	noResult := model.Hotel{
		HTMLAttributions: []interface{}{},
		NextPageToken:    "",
		Results:          []model.ActivityResult{},
		Status:           "ZERO_RESULTS",
	}

	router := gin.Default()
	router.GET("/place/nearbysearch/json", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, &noResult)
	})
	server := httptest.NewServer(router)

	result, err := suite.svc.Sleep(server.URL, model.Location{Lat: 48.856614, Lng: 2.3522219}, params)
	suite.Error(err, "error: no results")
	suite.Equal(http.StatusNotFound, err.(*model.AppError).StatusCode)
	suite.Equal(noResult.Status, err.Error())
	suite.Nil(result)
}

func (suite *googleServiceSuite) TestEatWithGoodAnswer() {
	params := model.Constraints{
		Radius:   0,
		MaxPrice: 500,
		MinPrice: 0,
		OpenNow:  false,
	}
	location := model.Location{Lat: 48.856614, Lng: 2.3522219}
	restaurant := []model.ActivityResult{
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
			Types:            []string{"restaurant"},
			UserRatingsTotal: 5,
			Vicinity:         "",
		},
	}
	withResults := model.Activity{
		HTMLAttributions: []interface{}{},
		Results:          restaurant,
		Status:           "no error",
	}
	router := gin.Default()
	router.GET("/place/nearbysearch/json", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, &withResults)
	})
	server := httptest.NewServer(router)
	result, err := suite.svc.Eat(server.URL, location, params)
	suite.NoError(err, "no crashed")
	suite.Equal(restaurant, *result, "result and error are the same")
}

func (suite *googleServiceSuite) TestEatWithZeroResult() {
	params := model.Constraints{
		Radius:   0,
		MaxPrice: 500,
		MinPrice: 0,
		OpenNow:  false,
	}
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

	result, err := suite.svc.Eat(server.URL, model.Location{Lat: 48.856614, Lng: 2.3522219}, params)
	suite.Error(err, "error: no results")
	suite.Equal(http.StatusNotFound, err.(*model.AppError).StatusCode)
	suite.Equal(noResult.Status, err.Error())
	suite.Nil(result)
}

func (suite *googleServiceSuite) TestDrinkWithGoodAnswer() {
	params := model.Constraints{
		Radius:   0,
		MaxPrice: 500,
		MinPrice: 0,
		OpenNow:  false,
	}
	location := model.Location{Lat: 48.856614, Lng: 2.3522219}
	restaurant := []model.ActivityResult{
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
			Types:            []string{"restaurant"},
			UserRatingsTotal: 5,
			Vicinity:         "",
		},
	}
	withResults := model.Activity{
		HTMLAttributions: []interface{}{},
		Results:          restaurant,
		Status:           "no error",
	}
	router := gin.Default()
	router.GET("/place/nearbysearch/json", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, &withResults)
	})
	server := httptest.NewServer(router)
	result, err := suite.svc.Drink(server.URL, location, params)
	suite.NoError(err, "no crashed")
	suite.Equal(restaurant, *result, "result and error are the same")
}

func (suite *googleServiceSuite) TestDrinkWithZeroResult() {
	params := model.Constraints{
		Radius:   0,
		MaxPrice: 500,
		MinPrice: 0,
		OpenNow:  false,
	}
	noResult := model.Hotel{
		HTMLAttributions: []interface{}{},
		NextPageToken:    "",
		Results:          []model.ActivityResult{},
		Status:           "ZERO_RESULTS",
	}

	router := gin.Default()
	router.GET("/place/nearbysearch/json", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, &noResult)
	})
	server := httptest.NewServer(router)

	result, err := suite.svc.Drink(server.URL, model.Location{Lat: 48.856614, Lng: 2.3522219}, params)
	suite.Error(err, "error: no results")
	suite.Equal(http.StatusNotFound, err.(*model.AppError).StatusCode)
	suite.Equal(noResult.Status, err.Error())
	suite.Nil(result)
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

func (suite *googleServiceSuite) TestGetDirections_Positive() {
	googleDirectionResponse := model.GoogleDirectionsResponse{
		GeocodedWaypoints: []model.GeocodedWaypoints{{
			GeocoderStatus: "OK",
			PlaceID:        "ChIJDbdkHFQayUwR7-8fITgxTmU",
			Types:          []string{"locality", "political"},
		}, {
			GeocoderStatus: "OK",
			PlaceID:        "ChIJpTvG15DL1IkRd8S0KlBVNTI",
			Types:          []string{"locality", "political"},
		}},
		Routes: []model.Routes{{
			Bounds: model.Bounds{
				Northeast: model.Location{
					Lat: 45.5001121,
					Lng: -73.5559308,
				},
				Southwest: model.Location{
					Lat: 43.6452417,
					Lng: -79.3806321,
				},
			},
			Copyrights: "Map data ©2022 Google",
			Legs: []model.Legs{{
				ArrivalTime: model.GoogleTime{
					Text:     "12:03pm",
					TimeZone: "America/Toronto",
					Value:    1659974580,
				},
				DepartureTime: model.GoogleTime{
					Text:     "6:51am",
					TimeZone: "America/Toronto",
					Value:    1659955860,
				},
				Distance: model.Distance{
					Text:  "538 km",
					Value: 538294,
				},
				Duration: model.Duration{
					Text:  "5 hours 12 mins",
					Value: 18720,
				},
				EndAddress: "Toronto, ON, Canada",
				EndLocation: model.Location{
					Lat: 43.6452417,
					Lng: -79.3806321,
				},
				StartAddress: "Montreal, QC, Canada",
				StartLocation: model.Location{
					Lat: 45.5001121,
					Lng: -73.5665224,
				},
				Steps: []model.GoogleDestinationStep{{
					Distance: model.Distance{
						Text:  "538 km",
						Value: 538294,
					},
					Duration: model.Duration{
						Text:  "5 hours 12 mins",
						Value: 18720,
					},
					EndLocation: model.Location{
						Lat: 43.6452417,
						Lng: -79.3806321,
					},
					HTMLInstructions: "Train towards Toronto",
					Polyline:         model.Polyline{Points: ""},
					StartLocation: model.Location{
						Lat: 45.5001121,
						Lng: -73.5665224,
					},
					TransitDetails: model.TransitDetails{
						ArrivalStop: model.Stop{
							Location: model.Location{
								Lat: 43.6452417,
								Lng: -79.3806321,
							},
							Name: "Union",
						},
						ArrivalTime: model.GoogleTime{
							Text:     "12:03pm",
							TimeZone: "America/Toronto",
							Value:    1659974580,
						},
						DepartureStop: model.Stop{
							Location: model.Location{
								Lat: 45.5001121,
								Lng: -73.5665224,
							},
							Name: "Gare Centrale",
						},
						DepartureTime: model.GoogleTime{
							Text:     "6:51am",
							TimeZone: "America/Toronto",
							Value:    1659955860,
						},
						Headsign: "Toronto",
						Line: model.Line{
							Agencies: []model.Agencies{{
								Name:  "Via Rail Canada Inc",
								Phone: "1 888 VIA-RAIL",
								URL:   "http://www.viarail.ca/",
							}},
							Color:     "#f2c106",
							Name:      "Montréal - Toronto",
							ShortName: "VIA Rail",
							TextColor: "#000000",
							Vehicle: model.Vehicle{
								Icon: "//maps.gstatic.com/mapfiles/transit/iw2/6/rail2.png",
								Name: "Train",
								Type: "HEAVY_RAIL",
							},
						},
						NumStops: 6,
					},
					TravelMode: "TRANSIT",
				}},
				TrafficSpeedEntry: []interface{}{},
				ViaWaypoint:       []interface{}{},
			}},
			OverviewPolyline: model.Polyline{Points: ""},
			Summary:          "",
			Warnings:         []string{},
			WaypointOrder:    []interface{}{},
		}},
		Status: "OK",
	}

	router := gin.New()
	router.GET("/directions/json", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, &googleDirectionResponse)
	})
	server := httptest.NewServer(router)

	directionsFormData := model.DirectionsFormData{
		Origin:        "Montreal",
		Destination:   "Toronto",
		DepartureTime: time.Date(2022, 8, 8, 10, 0, 0, 0, time.UTC),
	}

	itineraries, err := suite.svc.GetDirections(server.URL, &directionsFormData)
	suite.NoError(err, "no error if result")
	suite.NotEmpty(itineraries)
	suite.Require().Equal(1, len(*itineraries))
	suite.Equal(1, len((*itineraries)[0].Steps))
	suite.Equal(model.Ground, (*itineraries)[0].Type)

	departure := model.Station{
		Name:    "Montreal, QC, Canada",
		City:    "Montreal",
		Country: "Canada",
	}
	suite.Equal(departure, (*itineraries)[0].Departure)
	suite.Equal(5*time.Hour+12*time.Minute, (*itineraries)[0].Duration)

	loc, _ := time.LoadLocation("America/Toronto")
	arrivalTime := time.Unix(1659974580, 0).In(loc)
	suite.Equal(arrivalTime, (*itineraries)[0].Enddate)
}

func (suite *googleServiceSuite) TestGetDirections_NotFound_Negative() {
	noResults := model.GoogleDirectionsResponse{
		GeocodedWaypoints: []model.GeocodedWaypoints{{
			GeocoderStatus: "OK",
			PlaceID:        "ChIJ51cu8IcbXWARiRtXIothAS4",
			Types:          []string{"administrative_area_level_1", "political"},
		}, {
			GeocoderStatus: "OK",
			PlaceID:        "ChIJ8cM8zdaoAWARPR27azYdlsA",
			Types:          []string{"locality", "political"},
		}},
		Routes: []model.Routes{},
		Status: "ZERO_RESULTS",
	}

	router := gin.New()
	router.GET("/directions/json", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, &noResults)
	})
	server := httptest.NewServer(router)

	directionsFormData := model.DirectionsFormData{
		Origin:        "Montreal",
		Destination:   "Toronto",
		DepartureTime: time.Date(2022, 8, 8, 10, 0, 0, 0, time.UTC),
	}

	itineraries, err := suite.svc.GetDirections(server.URL, &directionsFormData)
	suite.Error(err, "error: not found")
	suite.Equal(http.StatusNotFound, err.(*model.AppError).StatusCode)
	suite.Equal("ZERO_RESULTS", err.Error())
	suite.Nil(itineraries)
}

func TestGoogleService(t *testing.T) {
	cfg := config.GetConfig(string(config.Test))
	suite.Run(t, &googleServiceSuite{cfg: cfg})
}
