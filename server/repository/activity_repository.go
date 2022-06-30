package repository

import "github.com/alexandre-pinon/epic-road-trip/model"

type GoogleRepository interface {
	Enjoy(position string) ([]*model.Activity , error)
}