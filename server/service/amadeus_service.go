package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/alexandre-pinon/epic-road-trip/config"
	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/alexandre-pinon/epic-road-trip/utils"
	"github.com/gin-gonic/gin"
)

type amadeusService struct {
	cfg *config.Config
}

type AmadeusService interface {
	GetAccessToken(amadeusBaseUrl string) (*model.AccessToken, error)
	GetFlightOffers(amadeusBaseUrl, accessToken string, flightFormData *model.FlightFormData) (*[]model.Itinerary, error)
}

func NewAmadeusService(cfg *config.Config) AmadeusService {
	return &amadeusService{cfg}
}

func (svc *amadeusService) GetAccessToken(amadeusBaseUrl string) (*model.AccessToken, error) {
	requestBody := url.Values{}
	requestBody.Set("grant_type", "client_credentials")
	requestBody.Set("client_id", svc.cfg.Amadeus.Key)
	requestBody.Set("client_secret", svc.cfg.Amadeus.Secret)

	response, err := http.Post(
		fmt.Sprintf("%s/v1/security/oauth2/token", amadeusBaseUrl),
		gin.MIMEPOSTForm,
		strings.NewReader(requestBody.Encode()),
	)
	if err != nil {
		return nil, &model.AppError{
			StatusCode: response.StatusCode,
			Err:        err,
		}
	}

	if response.StatusCode == http.StatusUnauthorized {
		return nil, &model.AppError{
			StatusCode: response.StatusCode,
			Err:        errors.New("invalid client credentials"),
		}
	}

	defer response.Body.Close()
	responseBody := model.AccessTokenResponse{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	now := time.Now().Unix()
	accessToken := model.AccessToken{
		Value: responseBody.AccessToken,
		Iat:   now,
		Exp:   int(now) + responseBody.ExpiresIn,
	}

	return &accessToken, nil
}

func (svc *amadeusService) GetFlightOffers(amadeusBaseUrl, accessToken string, flightFormData *model.FlightFormData) (*[]model.Itinerary, error) {
	query := fmt.Sprintf("originLocationCode=%s", flightFormData.OriginLocationCode)
	query += fmt.Sprintf("&destinationLocationCode=%s", flightFormData.DestinationLocationCode)
	query += fmt.Sprintf("&departureDate=%s", flightFormData.DepartureDate.Format("2006-01-02"))
	query += fmt.Sprintf("&adults=%d", flightFormData.Adults)
	query += "&nonStop=true&max=20"
	url := fmt.Sprintf("%s/v2/shopping/flight-offers?%s", amadeusBaseUrl, query)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, &model.AppError{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, &model.AppError{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		responseBody := model.FlightOffersResponseError{}
		json.NewDecoder(response.Body).Decode(&responseBody)

		return nil, &model.AppError{
			StatusCode: responseBody.Errors.Code,
			Err:        errors.New(responseBody.Errors.Title),
		}
	}

	responseBody := model.FlightOffersResponse{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	if len(responseBody.Data) == 0 {
		return nil, &model.AppError{
			StatusCode: http.StatusNotFound,
			Err:        errors.New("flight offers not found"),
		}
	}

	var itineraries []model.Itinerary
	for _, flightOffer := range responseBody.Data {
		departure := flightOffer.Itineraries[0].Segments[0].Departure
		arrival := flightOffer.Itineraries[0].Segments[0].Arrival

		startdate, _ := time.Parse(time.RFC3339, departure.At+"Z")
		enddate, _ := time.Parse(time.RFC3339, arrival.At+"Z")
		duration := utils.ExtractAmadeusTime(flightOffer.Itineraries[0].Duration)

		stationDeparture := model.Station{
			Name:    departure.IataCode,
			City:    flightFormData.OriginLocation,
			Country: responseBody.Dictionaries.Locations[departure.IataCode].CountryCode,
		}
		stationArrival := model.Station{
			Name:    arrival.IataCode,
			City:    flightFormData.DestinationLocation,
			Country: responseBody.Dictionaries.Locations[arrival.IataCode].CountryCode,
		}

		price, _ := strconv.ParseFloat(flightOffer.Price.GrandTotal, 64)

		itinerary := model.Itinerary{
			Type:           model.Airplane,
			Departure:      stationDeparture,
			Arrival:        stationArrival,
			Duration:       duration,
			DurationString: duration.String(),
			Startdate:      startdate,
			Enddate:        enddate,
			Price:          price,
		}
		itineraries = append(itineraries, itinerary)
	}

	return &itineraries, nil
}
