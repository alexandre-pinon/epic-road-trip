package service

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alexandre-pinon/epic-road-trip/config"
	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type amadeusServiceSuite struct {
	suite.Suite
	amadeusService AmadeusService
	cfg            *config.Config
}

func (suite *amadeusServiceSuite) SetupTest() {
	amadeusService := NewAmadeusService(suite.cfg)
	suite.amadeusService = amadeusService
}

func (suite *amadeusServiceSuite) TestGetAccessToken_Positive() {
	accessTokenResponse := model.AccessTokenResponse{
		Type:            "amadeusOAuth2Token",
		Username:        "alexandre@yakow.io",
		ApplicationName: "epic-road-tirp",
		ClientID:        "amcU3qA6gA7D68dlgR85BC4KNIt1j4vG",
		TokenType:       "Bearer",
		AccessToken:     "kzijAxGphgIbhSm3WGIhmWvhep2G",
		ExpiresIn:       1799,
		State:           "approved",
		Scope:           "",
	}
	router := gin.New()
	router.POST("/v1/security/oauth2/token", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, &accessTokenResponse)
	})
	server := httptest.NewServer(router)

	accessToken, err := suite.amadeusService.GetAccessToken(server.URL)
	suite.NoError(err, "no error if result")
	suite.Equal(accessTokenResponse.AccessToken, accessToken)
}

func (suite *amadeusServiceSuite) TestGetAccessToken_InvalidCredentials_Negative() {
	accessTokenResponse := model.AccessTokenError{
		Error:            "invalid_client",
		ErrorDescription: "Client credentials are invalid",
		Code:             38187,
		Title:            "Invalid parameters",
	}
	router := gin.New()
	router.POST("/v1/security/oauth2/token", func(ctx *gin.Context) {
		ctx.JSON(http.StatusUnauthorized, &accessTokenResponse)
	})
	server := httptest.NewServer(router)

	accessToken, err := suite.amadeusService.GetAccessToken(server.URL)
	suite.Error(err, "error when bad credentials")
	suite.Equal(http.StatusUnauthorized, err.(*model.AppError).StatusCode)
	suite.Equal("invalid client credentials", err.Error())
	suite.Empty(accessToken)
}

func (suite *amadeusServiceSuite) TestGetFlightOffers_Positive() {
	flightOfferResponse := model.FlighOffersResponse{
		Meta: model.FlightOfferMeta{
			Count: 1,
			Links: model.FlightOfferMetaLinks{
				Self: "https://test.api.amadeus.com/v2/shopping/flight-offers?originLocationCode=LON&destinationLocationCode=TYO&departureDate=2022-12-05&adults=1&nonStop=true&max=50",
			},
		},
		Data: []model.FlightOffer{{
			Type:                     "flight-offer",
			ID:                       "1",
			Source:                   "GDS",
			InstantTicketingRequired: false,
			NonHomogeneous:           false,
			OneWay:                   false,
			LastTicketingDate:        "2022-11-07",
			NumberOfBookableSeats:    9,
			Itineraries: []model.FlightOfferItinary{{
				Duration: "PT11H50M",
				Segments: []model.Segment{{
					Departure: model.Departure{
						IataCode: "LHR",
						Terminal: "2",
						At:       "2022-12-05T19:00:00",
					},
					Arrival: model.Arrival{
						IataCode: "HND",
						At:       "2022-12-06T15:50:00",
					},
					CarrierCode: "NH",
					Number:      "212",
					Aircraft: model.Aircraft{
						Code: "77W",
					},
					Operating: model.Operating{
						CarrierCode: "NH",
					},
					Duration:        "PT11H50M",
					ID:              "2",
					NumberOfStops:   0,
					BlacklistedInEU: false,
				}},
			}},
			Price: model.Price{
				BasePrice: model.BasePrice{
					Currency: "EUR",
					Total:    "1229.85",
					Base:     "776.00",
				},
				Fees: []model.Fee{{
					Amount: "0.00",
					Type:   "SUPPLIER",
				}, {
					Amount: "0.00",
					Type:   "TICKETING",
				}},
				GrandTotal: "1229.85",
			},
			PricingOptions: model.PricingOptions{
				FareType:                []string{"PUBLISHED"},
				IncludedCheckedBagsOnly: true,
			},
			ValidatingAirlineCodes: []string{"NH"},
			TravelerPricings: []model.TravelerPricing{{
				TravelerID:   "1",
				FareOption:   "STANDARD",
				TravelerType: "ADULT",
				Price: model.BasePrice{
					Currency: "EUR",
					Total:    "1229.85",
					Base:     "776.00",
				},
				FareDetailsBySegment: []model.FareDetails{{
					SegmentID: "2",
					Cabin:     "ECONOMY",
					FareBasis: "HLRCOJGB",
					Class:     "H",
					IncludedCheckedBags: model.IncludedCheckedBags{
						Quantity: 2,
					},
				}},
			}},
		}},
		Dictionaries: model.FlightOfferDictionary{
			Locations: map[string]model.FlightOfferLocations{
				"LHR": {
					CityCode:    "LON",
					CountryCode: "GB",
				},
				"HND": {
					CityCode:    "TYO",
					CountryCode: "JP",
				},
			},
			Aircraft: map[string]string{
				"773": "BOEING 777-300",
				"77W": "BOEING 777-300ER",
				"788": "BOEING 787-8",
				"789": "BOEING 787-9",
			},
			Currencies: map[string]string{
				"EUR": "EURO",
			},
			Carriers: map[string]string{
				"JL": "JAPAN AIRLINES",
				"NH": "ALL NIPPON AIRWAYS",
				"BA": "BRITISH AIRWAYS",
			},
		},
	}
	router := gin.New()
	router.GET("/v2/shopping/flight-offers", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, &flightOfferResponse)
	})
	server := httptest.NewServer(router)

	flightFormData := model.FlightFormData{
		OriginLocation:          "London",
		DestinationLocation:     "Tokyo",
		OriginLocationCode:      "LON",
		DestinationLocationCode: "TYO",
		DepartureDate:           time.Date(2022, 12, 5, 0, 0, 0, 0, time.UTC),
		Adults:                  1,
	}
	accessToken := "kzijAxGphgIbhSm3WGIhmWvhep2G"

	flightOffers, err := suite.amadeusService.GetFlightOffers(server.URL, accessToken, &flightFormData)
	suite.NoError(err, "no error if result")
	suite.NotEmpty(flightOffers)
	suite.Equal(1, len(*flightOffers))
	suite.Equal(1229.85, (*flightOffers)[0].Price)

	departure := model.Station{
		Name:    "LHR",
		City:    "London",
		Country: "GB",
	}
	suite.Equal(departure, (*flightOffers)[0].Departure)
	suite.Equal(time.Date(2022, 12, 6, 15, 50, 0, 0, time.UTC), (*flightOffers)[0].Enddate)
}

func (suite *amadeusServiceSuite) TestGetFlightOffers_NotFound_Negative() {
	flightOfferResponse := model.FlighOffersResponse{
		Meta: model.FlightOfferMeta{
			Count: 0,
			Links: model.FlightOfferMetaLinks{
				Self: "https://test.api.amadeus.com/v2/shopping/flight-offers?originLocationCode=ORY&destinationLocationCode=TYO&departureDate=2022-12-05&adults=1&nonStop=true&max=50",
			},
		},
		Data: []model.FlightOffer{},
	}
	router := gin.New()
	router.GET("/v2/shopping/flight-offers", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, &flightOfferResponse)
	})
	server := httptest.NewServer(router)

	flightFormData := model.FlightFormData{
		OriginLocationCode:      "ORY",
		DestinationLocationCode: "TYO",
		DepartureDate:           time.Date(2022, 12, 05, 0, 0, 0, 0, time.UTC),
		Adults:                  1,
	}
	accessToken := "kzijAxGphgIbhSm3WGIhmWvhep2G"

	flightOffers, err := suite.amadeusService.GetFlightOffers(server.URL, accessToken, &flightFormData)
	suite.Error(err, "error: not found")
	suite.Equal(http.StatusNotFound, err.(*model.AppError).StatusCode)
	suite.Equal("flight offers not found", err.Error())
	suite.Nil(flightOffers)
}

func TestAmadeusService(t *testing.T) {
	gin.SetMode(gin.TestMode)
	cfg := config.GetConfig()
	cfg.App.Env = config.Test
	suite.Run(t, &amadeusServiceSuite{cfg: cfg})
}
