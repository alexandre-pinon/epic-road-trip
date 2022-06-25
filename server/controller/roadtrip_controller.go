package controller

import (
	"errors"
	"net/http"
	"time"

	"github.com/alexandre-pinon/epic-road-trip/config"
	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/alexandre-pinon/epic-road-trip/service"
	"github.com/alexandre-pinon/epic-road-trip/utils"
	"github.com/gin-gonic/gin"
)

type roadtripController struct {
	cfg                *config.Config
	googleService      service.GoogleService
	amadeusService     service.AmadeusService
	amadeusAccessToken model.AccessToken
}

type RoadTripController interface {
	Enjoy(ctx *gin.Context) (*model.AppResult, *model.AppError)
	Travel(ctx *gin.Context) (*model.AppResult, *model.AppError)
	TravelAir(ctx *gin.Context) (*model.AppResult, *model.AppError)
	TravelGround(ctx *gin.Context) (*model.AppResult, *model.AppError)
}

func NewRoadTripController(cfg *config.Config, googleService service.GoogleService, amadeusService service.AmadeusService) RoadTripController {
	amadeusAccessToken := model.AccessToken{}
	return &roadtripController{cfg, googleService, amadeusService, amadeusAccessToken}
}

func (ctrl *roadtripController) Enjoy(c *gin.Context) (*model.AppResult, *model.AppError) {
	var position model.CityFormData

	if err := c.ShouldBindJSON(&position); err != nil {
		return nil, &model.AppError{
			StatusCode: http.StatusBadRequest,
			Err:        err,
		}
	}

	location, err := ctrl.googleService.GeoCoding(ctrl.cfg.Google.BaseUrl, position.City)
	if err != nil {
		return nil, err.(*model.AppError)
	}

	activities, err := ctrl.googleService.Enjoy(ctrl.cfg.Google.BaseUrl, *location)
	if err != nil {
		return nil, err.(*model.AppError)
	}

	return &model.AppResult{
		StatusCode: http.StatusOK,
		Message:    "Activities retrieved successfuly",
		Data:       &activities,
	}, nil
}

func (ctrl *roadtripController) Travel(ctx *gin.Context) (*model.AppResult, *model.AppError) {
	mode, _ := ctx.Get("mode")
	switch mode {
	case model.Air:
		return ctrl.TravelAir(ctx)
	case model.Ground:
		return ctrl.TravelGround(ctx)
	default:
		return nil, &model.AppError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("invalid travel mode"),
		}
	}
}

func (ctrl *roadtripController) TravelAir(ctx *gin.Context) (*model.AppResult, *model.AppError) {
	var flightFormData model.FlightFormData

	if err := ctx.ShouldBindJSON(&flightFormData); err != nil {
		return nil, &model.AppError{
			StatusCode: http.StatusBadRequest,
			Err:        err,
		}
	}

	flightFormData.OriginLocationCode = utils.CityToIata(flightFormData.OriginLocation)
	flightFormData.DestinationLocationCode = utils.CityToIata(flightFormData.DestinationLocation)
	if flightFormData.OriginLocationCode == "" || flightFormData.DestinationLocationCode == "" {
		return nil, &model.AppError{
			StatusCode: http.StatusNotFound,
			Err:        errors.New("no airport found for origin/destination cities"),
		}
	}

	if time.Now().Unix() > int64(ctrl.amadeusAccessToken.Exp) {
		accessToken, err := ctrl.amadeusService.GetAccessToken(ctrl.cfg.Amadeus.BaseUrl)
		if err != nil {
			return nil, err.(*model.AppError)
		}
		ctrl.amadeusAccessToken = *accessToken
	}

	itineraries, err := ctrl.amadeusService.GetFlightOffers(ctrl.cfg.Amadeus.BaseUrl, ctrl.amadeusAccessToken.Value, &flightFormData)
	if err != nil {
		return nil, err.(*model.AppError)
	}

	return &model.AppResult{
		StatusCode: http.StatusOK,
		Message:    "Itineraries retrieved successfully",
		Data:       &itineraries,
	}, nil
}

func (ctrl *roadtripController) TravelGround(ctx *gin.Context) (*model.AppResult, *model.AppError) {
	var directionsFormData model.DirectionsFormData

	if err := ctx.ShouldBindJSON(&directionsFormData); err != nil {
		return nil, &model.AppError{
			StatusCode: http.StatusBadRequest,
			Err:        err,
		}
	}

	itineraries, err := ctrl.googleService.GetDirections(ctrl.cfg.Google.BaseUrl, &directionsFormData)
	if err != nil {
		return nil, err.(*model.AppError)
	}

	return &model.AppResult{
		StatusCode: http.StatusOK,
		Message:    "Itineraries retrieved successfully",
		Data:       &itineraries,
	}, nil
}
