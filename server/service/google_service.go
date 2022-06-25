package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/alexandre-pinon/epic-road-trip/config"
	"github.com/alexandre-pinon/epic-road-trip/model"
)

type googleService struct {
	cfg *config.Config
}

type GoogleService interface {
	Enjoy(url string, position model.Location) (*[]model.ActivityResult, error)
	GeoCoding(url, position string) (*model.Location, error)
	GetDirections(url string, directionsFormData *model.DirectionsFormData) (*[]model.Itinerary, error)
}

func NewGoogleService(cfg *config.Config) GoogleService {
	return &googleService{cfg}
}

func (svc *googleService) Enjoy(url string, position model.Location) (*[]model.ActivityResult, error) {

	response, err := http.Get(fmt.Sprintf("%s/place/nearbysearch/json?location=%f,%f&radius=5000&type=tourist_attraction&key=%s", url, position.Lat, position.Lng, svc.cfg.Google.Key))
	if err != nil {
		return nil, &model.AppError{
			StatusCode: http.StatusBadRequest,
			Err:        err,
		}
	}
	defer response.Body.Close()
	responseBody := model.Activity{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	if len(responseBody.Results) == 0 {
		return nil, &model.AppError{
			StatusCode: http.StatusNotFound,
			Err:        errors.New(responseBody.Status),
		}
	}

	return &responseBody.Results, nil
}

func (svc *googleService) GeoCoding(url, city string) (*model.Location, error) {
	response, err := http.Get(
		fmt.Sprintf("%s/geocode/json?address=%s&key=%s", url, city, svc.cfg.Google.Key),
	)
	if err != nil {
		return nil, &model.AppError{
			StatusCode: http.StatusBadRequest,
			Err:        err,
		}
	}

	defer response.Body.Close()
	responseBody := model.GoogleGeocodingResponse{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	if len(responseBody.Results) == 0 {
		return nil, &model.AppError{
			StatusCode: http.StatusNotFound,
			Err:        errors.New(responseBody.Status),
		}
	}

	location := responseBody.Results[0].Geometry.Location

	return &location, nil
}

func (svc *googleService) GetDirections(url string, directionsFormData *model.DirectionsFormData) (*[]model.Itinerary, error) {
	return &[]model.Itinerary{}, &model.AppError{
		StatusCode: http.StatusNotImplemented,
		Err:        errors.New("TODO: implement GetDirections"),
	}
}
