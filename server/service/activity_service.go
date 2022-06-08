package service

import (
	"errors"
	"net/http"

	"github.com/alexandre-pinon/epic-road-trip/model"
)

type ActivityService interface {
	Enjoy(position string) (*[]model.Activity , error)
}

func  Enjoy(position string) (*[]model.Activity, error) {

	result := []model.Activity{
		{
			Name:      "yoimiya",
		},
		{
			Name:      "hu",
		},
		{
			Name:      "kokomi",
		},
	}

	if position == "" {
		return nil , &model.AppError{
			StatusCode: http.StatusNotFound,
			Err:        errors.New("user not found"),
		}
	} else {
		return &result , &model.AppError{}
	}

}

