package service

import (
	"errors"
	"net/http"

	"github.com/alexandre-pinon/epic-road-trip/config"
	"github.com/alexandre-pinon/epic-road-trip/model"
)

type amadeusService struct {
	cfg config.Config
}

type AmadeusService interface {
	GetAccessToken(amadeusBaseUrl string) (string, error)
}

func NewAmadeusService(cfg config.Config) AmadeusService {
	return &amadeusService{cfg}
}

func (svc *amadeusService) GetAccessToken(amadeusBaseUrl string) (string, error) {
	return "", &model.AppError{
		StatusCode: http.StatusNotImplemented,
		Err:        errors.New("TODO: implement GetAccessToken"),
	}
}
