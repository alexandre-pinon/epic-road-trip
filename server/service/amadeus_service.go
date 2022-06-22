package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/alexandre-pinon/epic-road-trip/config"
	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/gin-gonic/gin"
)

type amadeusService struct {
	cfg config.Config
}

type AmadeusService interface {
	GetAccessToken(amadeusBaseUrl string) (string, error)
	GetFlightOffers(amadeusBaseUrl, accessToken string, flightFormData *model.FlightFormData) (*model.FlighOffersResponse, error)
}

func NewAmadeusService(cfg config.Config) AmadeusService {
	return &amadeusService{cfg}
}

func (svc *amadeusService) GetAccessToken(amadeusBaseUrl string) (string, error) {
	requestBody := url.Values{}
	requestBody.Set("grant_type", "client_credentials")
	requestBody.Set("client_id", svc.cfg.AmadeusKey)
	requestBody.Set("client_secret", svc.cfg.AmadeusSecret)

	response, err := http.Post(
		fmt.Sprintf("%s/v1/security/oauth2/token", amadeusBaseUrl),
		gin.MIMEPOSTForm,
		strings.NewReader(requestBody.Encode()),
	)
	if err != nil {
		return "", &model.AppError{
			StatusCode: response.StatusCode,
			Err:        err,
		}
	}

	if response.StatusCode == http.StatusUnauthorized {
		return "", &model.AppError{
			StatusCode: response.StatusCode,
			Err:        errors.New("invalid client credentials"),
		}
	}

	defer response.Body.Close()
	responseBody := model.AccessTokenResponse{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	return responseBody.AccessToken, nil
}

func (svc *amadeusService) GetFlightOffers(amadeusBaseUrl, accessToken string, flightFormData *model.FlightFormData) (*model.FlighOffersResponse, error) {
	query := fmt.Sprintf("originLocationCode=%s", flightFormData.OriginLocationCode)
	query += fmt.Sprintf("&destinationLocationCode=%s", flightFormData.DestinationLocationCode)
	query += fmt.Sprintf("&departureDate=%s", flightFormData.DepartureDate.Format("2000-01-01"))
	query += fmt.Sprintf("&adults=%d", flightFormData.Adults)
	query += "&nonStop=true&max=50"
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
	responseBody := model.FlighOffersResponse{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	if len(responseBody.Data) == 0 {
		return nil, &model.AppError{
			StatusCode: http.StatusNotFound,
			Err:        errors.New("flight offers not found"),
		}
	}

	return &responseBody, nil
}
