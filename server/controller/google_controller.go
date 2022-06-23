package controller

import (
	"net/http"

	"github.com/alexandre-pinon/epic-road-trip/config"
	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/alexandre-pinon/epic-road-trip/service"
	"github.com/gin-gonic/gin"
)

type roadtripController struct {
	googleService service.GoogleService
}

type RoadTripController interface {
	Enjoy(ctx *gin.Context) (*model.AppResult, *model.AppError)
}

func NewRoadTripController(svc service.GoogleService) RoadTripController {
	return &roadtripController{svc}
}

func (crtl *roadtripController) Enjoy(c *gin.Context) (*model.AppResult, *model.AppError) {

	var position model.CityFormData

	if err := c.ShouldBindJSON(&position); err != nil {
		return nil, &model.AppError{
			StatusCode: http.StatusBadRequest,
			Err:        err,
		}
	}

	location, err := crtl.googleService.GeoCoding(config.GetConfig().BaseUrlGoogle, position.City)
	if err != nil {
		return nil, &model.AppError{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	activities, err := crtl.googleService.Enjoy(config.GetConfig().BaseUrlGoogle, *location)
	if err != nil {
		return nil, &model.AppError{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	if len(*activities) == 0 {
		return &model.AppResult{
			StatusCode: http.StatusOK,
			Message:    "Activities retrieved successfuly but empty",
			Data:       []model.ActivityResult{},
		}, nil
	}

	return &model.AppResult{
		StatusCode: http.StatusOK,
		Message:    "Activities retrieved successfuly",
		Data:       &activities,
	}, nil
}
