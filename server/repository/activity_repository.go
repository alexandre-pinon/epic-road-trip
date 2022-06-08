package repository

import "github.com/alexandre-pinon/epic-road-trip/model"

type ActivityRepository interface {
	Enjoy(position string) ([]*model.Activity , error)
}