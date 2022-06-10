package repository

import (
	"github.com/alexandre-pinon/epic-road-trip/cache"
	"github.com/stretchr/testify/suite"
	
)


type googleRepositorySuite struct {
	suite.Suite 
	cfg	cache.ActivityCache
	repo GoogleRepository
	
}