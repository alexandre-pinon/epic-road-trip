package repository

import (
	"testing"

	"github.com/alexandre-pinon/epic-road-trip/config"
	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/alexandre-pinon/epic-road-trip/utils"
	"github.com/stretchr/testify/suite"
)

type activityRepositorySuite struct {
	suite.Suite
	cfg             config.Config
	repo            ActivityRepository
	cleanupExecutor utils.DropCollectionExecutor
}

func (suite *activityRepositorySuite) EnjoyTest() {
	insertActivities := []model.Activity{
		{
			Name:      "yoimiya",
		},
		{
			Name:      "hu",
		},
		{
			Name:      "kokomi",
		}
	}

	activities, err := suite.repo.Enjoy("test")
	suite.NoError(err, "no error when get activities")
	suite.Equal(0, len(*users), "insert 0 records before the all data, so it should contain nothing")
}
