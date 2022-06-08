package service

import (
	"github.com/alexandre-pinon/epic-road-trip/model"
)

type ActivityService interface {
	Enjoy(position string) (*[]model.Activity , error)
}


