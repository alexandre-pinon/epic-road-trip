package cache

import "github.com/alexandre-pinon/epic-road-trip/model"

type ActivityCache interface {
	Set(key string, value *model.Activity ) 
	Get(key string) *model.Activity
}