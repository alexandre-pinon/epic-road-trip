package repository

import (
	"github.com/alexandre-pinon/epic-road-trip/cache"
	"github.com/stretchr/testify/suite"
	
)

type activityRepositorySuite struct {
	suite.Suite 
	repo ActivityRepository
}