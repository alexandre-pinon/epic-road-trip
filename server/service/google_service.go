package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/alexandre-pinon/epic-road-trip/config"
	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/alexandre-pinon/epic-road-trip/utils"
)

type googleService struct {
	cfg *config.Config
}

type GoogleService interface {
	Enjoy(url string, position model.Location, constraint model.Constraints) (*[]model.ActivityResult, error)
	Sleep(url string, position model.Location, constraint model.Constraints) (*[]model.ActivityResult, error)
	Eat(url string, position model.Location, constraint model.Constraints) (*[]model.ActivityResult, error)
	Drink(url string, position model.Location, constraint model.Constraints) (*[]model.ActivityResult, error)
	GeoCoding(url, position string) (*model.Location, error)
	GetDirections(url string, directionsFormData *model.DirectionsFormData) (*[]model.Itinerary, error)
}

func NewGoogleService(cfg *config.Config) GoogleService {
	return &googleService{cfg}
}

func (svc *googleService) Enjoy(googleBaseUrl string, position model.Location, constraint model.Constraints) (*[]model.ActivityResult, error) {
	query := fmt.Sprintf("location=%f,%f", position.Lat, position.Lng)
	query += fmt.Sprintf("&key=%s", svc.cfg.Google.Key)
	query += utils.ConstraintStringify(constraint)
	query += "&type=tourist_attraction"
	uri := fmt.Sprintf("%s/place/nearbysearch/json?%s", googleBaseUrl, url.PathEscape(query))
	response, err := http.Get(uri)

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

func (svc *googleService) Sleep(url string, position model.Location, constraint model.Constraints) (*[]model.ActivityResult, error) {

	params := utils.ConstraintStringify(constraint)

	response, err := http.Get(fmt.Sprintf("%s/place/nearbysearch/json?location=%f,%f&type=lodging%s&key=%s", url, position.Lat, position.Lng, params, svc.cfg.Google.Key))
	if err != nil {
		return nil, &model.AppError{
			StatusCode: http.StatusBadRequest,
			Err:        err,
		}
	}
	defer response.Body.Close()
	responseBody := model.Hotel{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	if len(responseBody.Results) == 0 {
		return nil, &model.AppError{
			StatusCode: http.StatusNotFound,
			Err:        errors.New(responseBody.Status),
		}
	}

	return &responseBody.Results, nil
}

func (svc *googleService) Eat(url string, position model.Location, constraint model.Constraints) (*[]model.ActivityResult, error) {

	params := utils.ConstraintStringify(constraint)

	response, err := http.Get(fmt.Sprintf("%s/place/nearbysearch/json?location=%f,%f&type=restaurant%s&key=%s", url, position.Lat, position.Lng, params, svc.cfg.Google.Key))
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

func (svc *googleService) Drink(url string, position model.Location, constraint model.Constraints) (*[]model.ActivityResult, error) {

	params := utils.ConstraintStringify(constraint)

	response, err := http.Get(fmt.Sprintf("%s/place/nearbysearch/json?location=%f,%f&type=bar%s&key=%s", url, position.Lat, position.Lng, params, svc.cfg.Google.Key))
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

func (svc *googleService) GeoCoding(googleBaseUrl, city string) (*model.Location, error) {
	query := fmt.Sprintf("address=%s", city)
	query += fmt.Sprintf("&key=%s", svc.cfg.Google.Key)
	uri := fmt.Sprintf("%s/geocode/json?%s", googleBaseUrl, url.PathEscape(query))
	response, err := http.Get(uri)
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

func (svc *googleService) GetDirections(googleBaseUrl string, directionsFormData *model.DirectionsFormData) (*[]model.Itinerary, error) {
	query := fmt.Sprintf("origin=%s", directionsFormData.Origin)
	query += fmt.Sprintf("&destination=%s", directionsFormData.Destination)
	query += fmt.Sprintf("&departure_time=%d", directionsFormData.DepartureTime.Unix())
	query += fmt.Sprintf("&key=%s", svc.cfg.Google.Key)
	query += "&mode=transit"
	uri := fmt.Sprintf("%s/directions/json?%s", googleBaseUrl, url.PathEscape(query))

	response, err := http.Get(uri)
	if err != nil {
		return nil, &model.AppError{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	defer response.Body.Close()

	responseBody := model.GoogleDirectionsResponse{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	if len(responseBody.Routes) == 0 {
		return nil, &model.AppError{
			StatusCode: http.StatusNotFound,
			Err:        errors.New(responseBody.Status),
		}
	}

	var itineraries []model.Itinerary
	route := responseBody.Routes[0].Legs[0]
	var googleDestinationSteps []model.GoogleDestinationStep

	for _, step := range route.Steps {
		if step.TravelMode == "TRANSIT" {
			googleDestinationSteps = append(googleDestinationSteps, step) // remove walking steps
		}
	}

	departureSplit := strings.Split(route.StartAddress, ", ")
	stationDeparture := model.Station{
		Name:    route.StartAddress,
		City:    departureSplit[0],
		Country: departureSplit[len(departureSplit)-1],
	}
	arrivalSplit := strings.Split(route.EndAddress, ", ")
	stationArrival := model.Station{
		Name:    route.StartAddress,
		City:    arrivalSplit[0],
		Country: arrivalSplit[len(arrivalSplit)-1],
	}

	startdate := utils.ExtractGoogleDate(route.DepartureTime.Value, route.DepartureTime.TimeZone)
	enddate := utils.ExtractGoogleDate(route.ArrivalTime.Value, route.ArrivalTime.TimeZone)
	duration := time.Duration(route.Duration.Value * int(time.Second))

	var steps []model.ItineraryStep
	for _, gStep := range googleDestinationSteps {
		details := gStep.TransitDetails
		step := model.ItineraryStep{
			Type:           details.Line.Vehicle.Name,
			Departure:      details.DepartureStop.Name,
			Arrival:        details.ArrivalStop.Name,
			Duration:       time.Duration(gStep.Duration.Value * int(time.Second)),
			DurationString: time.Duration(gStep.Duration.Value * int(time.Second)).String(),
			Startdate:      utils.ExtractGoogleDate(details.DepartureTime.Value, details.DepartureTime.TimeZone),
			Enddate:        utils.ExtractGoogleDate(details.ArrivalTime.Value, details.ArrivalTime.TimeZone),
		}
		steps = append(steps, step)
	}

	itinerary := model.Itinerary{
		Type:           model.Ground,
		Departure:      stationDeparture,
		Arrival:        stationArrival,
		Duration:       duration,
		DurationString: duration.String(),
		Startdate:      startdate,
		Enddate:        enddate,
		Steps:          steps,
	}
	itineraries = append(itineraries, itinerary)

	return &itineraries, nil
}
