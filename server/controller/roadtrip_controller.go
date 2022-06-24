package controller

import (
	"errors"
	"net/http"

	"github.com/alexandre-pinon/epic-road-trip/config"
	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/alexandre-pinon/epic-road-trip/service"
	"github.com/gin-gonic/gin"
)

type roadtripController struct {
	cfg            *config.Config
	googleService  service.GoogleService
	amadeusService service.AmadeusService
}

type RoadTripController interface {
	Enjoy(ctx *gin.Context) (*model.AppResult, *model.AppError)
	Travel(ctx *gin.Context) (*model.AppResult, *model.AppError)
}

func NewRoadTripController(cfg *config.Config, googleService service.GoogleService, amadeusService service.AmadeusService) RoadTripController {
	return &roadtripController{cfg, googleService, amadeusService}
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
	return &model.AppResult{}, &model.AppError{
		StatusCode: http.StatusNotImplemented,
		Err:        errors.New("TODO: implement Travel"),
	}
}
