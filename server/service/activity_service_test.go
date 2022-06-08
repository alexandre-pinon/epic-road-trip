package service

import (
	"github.com/alexandre-pinon/epic-road-trip/mocks"
	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/stretchr/testify/suite"
)

type activityServiceSuite struct {
	suite.Suite
	repo *mocks.ActivityRepository
	svc ActivityService
}

func (suite *activityServiceSuite) TestEnjoy() {
	activities := []model.Activity{
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
	suite.repo.On("Enjoy" , "test" ).Return(&activities, nil)
	result, err := suite.svc.Enjoy("test")
	suite.NoError(err, "no error when get all activities")
	suite.Equal(len(activities), len(*result), "activities and result should have the same length")
	suite.Equal(activities, *result, "result and activities are the same")
	suite.repo.AssertExpectations(suite.T())
}
