package cache

import "github.com/alexandre-pinon/epic-road-trip/model"

type UserCache interface {
	Set(key string, value model.User ) 
	Get(key string) *model.User
}