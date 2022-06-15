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
	cfg config.Config
}

type GoogleService interface {
	Enjoy(url string, position []float64) (*[]model.Activity, error)
	GeoCoding(url, position string) (*model.Location, error)
}

func NewGoogleService(cfg config.Config) GoogleService {
	return &googleService{cfg}
}

func (svc *googleService) Enjoy(url string, position []float64) (*[]model.Activity, error) {

	response, err := http.Get(fmt.Sprintf("%s/place/nearbysearch/json?location=-33.8670522,151.1957362&type=restaurant&key=%s", url, svc.cfg.GoogleKey))
	if err == nil {
		return nil, err
	}
	defer response.Body.Close()
	responseBody := []model.Activity{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	return &responseBody, nil
}

func (svc *googleService) GeoCoding(url, city string) (*model.Location, error) {
	response, err := http.Get(
		fmt.Sprintf("%s/geocode/json?address=%s&key=%s", url, city, svc.cfg.GoogleKey),
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
