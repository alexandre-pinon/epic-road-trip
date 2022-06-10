package service

import (
	"errors"

	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/alexandre-pinon/epic-road-trip/repository"
)

type googleService struct {
	activityRepository repository.GoogleRepository
}

type GoogleService interface {
	Enjoy(position string) (*[]model.Activity, error)
}

func NewGoogleService(repo repository.GoogleRepository) GoogleService  {
	return &googleService{repo}
}


func (svc *googleService) Enjoy(position string) (*[]model.Activity , error) {
	return &[]model.Activity{}, errors.New("TODO")
}