package service

import (
	"errors"
	"testing"
	"net/http"

	"github.com/alexandre-pinon/epic-road-trip/mocks"
	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/stretchr/testify/suite"
)

type googleServiceSuite struct {
	suite.Suite
	repo *mocks.GoogleRepository
	svc  GoogleService
}

func (suite *googleServiceSuite) SetupTest() {
	repo := new(mocks.GoogleRepository)
	svc := NewGoogleService(repo)

	suite.repo = repo
	suite.svc = svc		
	
}

func (suite *googleServiceSuite) TestEnjoyWithZeroResult() {
	activities := &model.AppError{
			StatusCode: http.StatusInternalServerError,
			Err: errors.New("ZERO RESULTS"),
	}

	suite.repo.On("Enjoy", "test").Return(&activities, nil)
	result, err := suite.svc.Enjoy("test")
	suite.NoError(err, "no crashed")
	suite.Equal(activities, *result, "result and error are the same")
	suite.repo.AssertExpectations(suite.T())
}

func (suite *googleServiceSuite) TestEnjoyWithGoodAnswer() {

	activities := []model.Activity{
		{
			Place_id: "CHIJyoimiya",
			Name:     "yoimiya",
			Opening_hours: []*model.Opening_hours{
				{
					Open_now: true,
					Periods: []*model.Periods{
						{
							Open_periods: []*model.Open{
								{
									Open_day:  0,
									Open_time: "2020",
								},
								{
									Open_day:  0,
									Open_time: "2020",
								},
							},
							Close_periods: []*model.Close{
								{
									Close_day:  0,
									Close_time: "2020",
								},
								{
									Close_day:  0,
									Close_time: "2020",
								},
							},
						},
					},
				},
			},
			Address:            "8 rue de ta mère c'est la catin du coin",
			Types:              "solide froid boite",
			User_ratings_total: 2,
			Price_level:        2,
			Rating:             4.2,
		},
		{
			Place_id: "CHIJyoimiya",
			Name:     "yoimiya",
			Opening_hours: []*model.Opening_hours{
				{
					Open_now: true,
					Periods: []*model.Periods{
						{
							Open_periods: []*model.Open{
								{
									Open_day:  0,
									Open_time: "2020",
								},
								{
									Open_day:  0,
									Open_time: "2020",
								},
							},
							Close_periods: []*model.Close{
								{
									Close_day:  0,
									Close_time: "2020",
								},
								{
									Close_day:  0,
									Close_time: "2020",
								},
							},
						},
					},
				},
			},
			Address:            "8 rue de ta mère c'est la catin du coin",
			Types:              "solide froid boite",
			User_ratings_total: 2,
			Price_level:        2,
			Rating:             4.2,
		},
		{
			Place_id: "CHIJyoimiya",
			Name:     "yoimiya",
			Opening_hours: []*model.Opening_hours{
				{
					Open_now: true,
					Periods: []*model.Periods{
						{
							Open_periods: []*model.Open{
								{
									Open_day:  0,
									Open_time: "2020",
								},
								{
									Open_day:  0,
									Open_time: "2020",
								},
							},
							Close_periods: []*model.Close{
								{
									Close_day:  0,
									Close_time: "2020",
								},
								{
									Close_day:  0,
									Close_time: "2020",
								},
							},
						},
					},
				},
			},
			Types:              "solide froid boite",
			User_ratings_total: 2,
			Price_level:        2,
			Rating:             4.2,
		},
	}
	suite.repo.On("Enjoy", "test").Return(&activities, nil)
	result, err := suite.svc.Enjoy("test")
	suite.NoError(err, "no error when get all activities")
	suite.Equal(len(activities), len(*result), "activities and result should have the same length")
	suite.Equal(activities, *result, "result and activities are the same")
	suite.repo.AssertExpectations(suite.T())
}

func TestGoogleService(t *testing.T) {
	suite.Run(t, new(googleServiceSuite))
}
