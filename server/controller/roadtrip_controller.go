package controller

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/alexandre-pinon/epic-road-trip/config"
	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/alexandre-pinon/epic-road-trip/repository"
	"github.com/alexandre-pinon/epic-road-trip/service"
	"github.com/alexandre-pinon/epic-road-trip/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type roadtripController struct {
	cfg                *config.Config
	userService        service.UserService
	googleService      service.GoogleService
	amadeusService     service.AmadeusService
	amadeusAccessToken model.AccessToken
	tripStepRepository repository.TripStepRepository
}

type RoadtripController interface {
	CreateRoadtrip(ctx *gin.Context) (*model.AppResult, *model.AppError)
	DeleteRoadtrip(ctx *gin.Context) (*model.AppResult, *model.AppError)
	Travel(ctx *gin.Context) (*model.AppResult, *model.AppError)
	TravelAir(ctx *gin.Context) (*model.AppResult, *model.AppError)
	TravelGround(ctx *gin.Context) (*model.AppResult, *model.AppError)
	Enjoy(ctx *gin.Context) (*model.AppResult, *model.AppError)
	Sleep(ctx *gin.Context) (*model.AppResult, *model.AppError)
	Eat(c *gin.Context) (*model.AppResult, *model.AppError)
	Drink(c *gin.Context) (*model.AppResult, *model.AppError)
}

func NewRoadtripController(cfg *config.Config, userService service.UserService, googleService service.GoogleService, amadeusService service.AmadeusService, tripStepRepository repository.TripStepRepository) RoadtripController {
	amadeusAccessToken := model.AccessToken{}
	return &roadtripController{cfg, userService, googleService, amadeusService, amadeusAccessToken, tripStepRepository}
}

// Create roadtrip godoc
// @Summary Create roadtrip
// @Description Add roadtrip to the user given userID & trip steps
// @Tags Roadtrip
// @Accept json
// @Produce json
// @Param userID query string true "User ID"
// @Param tripSteps body []model.TripStep true "city & start/end date & activities"
// @Success 200 {object} model.CreateRoadtripSuccess "Success"
// @Failure 400 {object} model.InvalidJsonBody "Invalid request params/body"
// @Failure 401 {object} model.Unauthorized "Missing/Expired token"
// @Failure 500 {object} model.InternalServerError "Internal server error"
// @Router /roadtrip [post]
func (ctrl *roadtripController) CreateRoadtrip(ctx *gin.Context) (*model.AppResult, *model.AppError) {
	var tripSteps []model.TripStep

	userIDParam, exists := ctx.GetQuery("userID")
	if !exists {
		return nil, &model.AppError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("invalid request parameters: missing userID"),
		}
	}

	userID, err := primitive.ObjectIDFromHex(userIDParam)
	if err != nil {
		return nil, &model.AppError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("invalid request parameters: invalid userID"),
		}
	}

	if err := ctx.ShouldBindJSON(&tripSteps); err != nil {
		return nil, &model.AppError{
			StatusCode: http.StatusBadRequest,
			Err:        err,
		}
	}

	user, err := ctrl.userService.GetUserByID(userID, false)
	if err != nil {
		return nil, err.(*model.AppError)
	}

	var insertedIDs []primitive.ObjectID
	for _, tripStep := range tripSteps {
		res, err := ctrl.tripStepRepository.CreateTripStep(&tripStep)
		if err != nil {
			return nil, &model.AppError{
				StatusCode: http.StatusInternalServerError,
				Err:        err,
			}
		}

		insertedIDs = append(insertedIDs, res.InsertedID.(primitive.ObjectID))
	}

	user.Trips = append(user.Trips, &model.Roadtrip{
		ID:          primitive.NewObjectID(),
		Startdate:   tripSteps[0].Startdate,
		Enddate:     tripSteps[len(tripSteps)-1].Enddate,
		TripStepsID: insertedIDs,
	})

	if err := ctrl.userService.UpdateUser(userID, user); err != nil {
		return nil, err.(*model.AppError)
	}

	return &model.AppResult{
		StatusCode: http.StatusOK,
		Message:    fmt.Sprintf("Added roadtrip to user %s successfully", userID.Hex()),
		Data:       struct{}{},
	}, nil
}

// Delete roadtrip godoc
// @Summary Delete roadtrip
// @Description Remove roadtrip from the user given userID & roadtrip ID
// @Tags Roadtrip
// @Accept json
// @Produce json
// @Param userID query string true "User ID"
// @Param id path string true "Roadtrip ID"
// @Success 200 {object} model.DeleteRoadtripSuccess "Success"
// @Failure 400 {object} model.InvalidID "Invalid ID/params"
// @Failure 401 {object} model.Unauthorized "Missing/Expired token"
// @Failure 404 {object} model.UserNotFound "User not found"
// @Failure 500 {object} model.InternalServerError "Internal server error"
// @Router /roadtrip/:id [delete]
func (ctrl *roadtripController) DeleteRoadtrip(ctx *gin.Context) (*model.AppResult, *model.AppError) {
	id, _ := ctx.Get("id")

	userIDParam, exists := ctx.GetQuery("userID")
	if !exists {
		return nil, &model.AppError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("invalid query parameters: missing userID"),
		}
	}

	userID, err := primitive.ObjectIDFromHex(userIDParam)
	if err != nil {
		return nil, &model.AppError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("invalid query parameters: invalid userID"),
		}
	}

	user, err := ctrl.userService.GetUserByID(userID, false)
	if err != nil {
		return nil, err.(*model.AppError)
	}

	for idx, roadtrip := range user.Trips {
		if roadtrip.ID == id.(primitive.ObjectID) {
			for _, tripStepID := range roadtrip.TripStepsID {
				_, err := ctrl.tripStepRepository.DeleteTripStep(tripStepID)
				if err != nil {
					return nil, &model.AppError{
						StatusCode: http.StatusInternalServerError,
						Err:        err,
					}
				}
			}
			user.Trips = append(user.Trips[:idx], user.Trips[idx+1:]...)
		}
	}

	if err := ctrl.userService.UpdateUser(userID, user); err != nil {
		return nil, err.(*model.AppError)
	}

	return &model.AppResult{
		StatusCode: http.StatusOK,
		Message:    fmt.Sprintf("Removed roadtrip from user %s successfully", userID.Hex()),
		Data:       struct{}{},
	}, nil
}

// Enjoy godoc
// @Summary Enjoy
// @Description Search for tourist attraction around the given city & constraints
// @Tags Roadtrip
// @Accept json
// @Produce json
// @Param cityFormData body model.CityFormData true "city & constraints"
// @Success 200 {object} model.EnjoySuccess "Success"
// @Failure 400 {object} model.InvalidJsonBody "Invalid request body"
// @Failure 401 {object} model.Unauthorized "Missing/Expired token"
// @Failure 404 {object} model.GoogleNotFound "Zero results"
// @Failure 500 {object} model.InternalServerError "Internal server error"
// @Router /roadtrip/enjoy [post]
func (ctrl *roadtripController) Enjoy(c *gin.Context) (*model.AppResult, *model.AppError) {
	var position model.CityFormData

	if err := c.ShouldBindJSON(&position); err != nil {
		return nil, &model.AppError{
			StatusCode: http.StatusBadRequest,
			Err:        err,
		}
	}

	location, err := ctrl.googleService.GeoCoding(ctrl.cfg.Google.BaseUrl, position.City)
	if err != nil {
		return nil, err.(*model.AppError)
	}

	activities, err := ctrl.googleService.Enjoy(ctrl.cfg.Google.BaseUrl, *location, position.Constraints)
	if err != nil {
		return nil, err.(*model.AppError)
	}

	return &model.AppResult{
		StatusCode: http.StatusOK,
		Message:    "Activities retrieved successfuly",
		Data:       &activities,
	}, nil
}

// Sleep godoc
// @Summary Sleep
// @Description Search for hotels around the given city & constraints
// @Tags Roadtrip
// @Accept json
// @Produce json
// @Param cityFormData body model.CityFormData true "city & constraints"
// @Success 200 {object} model.SleepSuccess "Success"
// @Failure 400 {object} model.InvalidJsonBody "Invalid request body"
// @Failure 401 {object} model.Unauthorized "Missing/Expired token"
// @Failure 404 {object} model.GoogleNotFound "Zero results"
// @Failure 500 {object} model.InternalServerError "Internal server error"
// @Router /roadtrip/sleep [post]
func (ctrl *roadtripController) Sleep(c *gin.Context) (*model.AppResult, *model.AppError) {
	var position model.CityFormData

	if err := c.ShouldBindJSON(&position); err != nil {
		return nil, &model.AppError{
			StatusCode: http.StatusBadRequest,
			Err:        err,
		}
	}

	location, err := ctrl.googleService.GeoCoding(ctrl.cfg.Google.BaseUrl, position.City)
	if err != nil {
		return nil, err.(*model.AppError)
	}

	hotels, err := ctrl.googleService.Sleep(ctrl.cfg.Google.BaseUrl, *location, position.Constraints)
	if err != nil {
		return nil, err.(*model.AppError)
	}

	return &model.AppResult{
		StatusCode: http.StatusOK,
		Message:    "Hotels retrieved successfuly",
		Data:       &hotels,
	}, nil
}

// Eat godoc
// @Summary Eat
// @Description Search for restaurants around the given city & constraints
// @Tags Roadtrip
// @Accept json
// @Produce json
// @Param cityFormData body model.CityFormData true "city & constraints"
// @Success 200 {object} model.EatSuccess "Success"
// @Failure 400 {object} model.InvalidJsonBody "Invalid request body"
// @Failure 401 {object} model.Unauthorized "Missing/Expired token"
// @Failure 404 {object} model.GoogleNotFound "Zero results"
// @Failure 500 {object} model.InternalServerError "Internal server error"
// @Router /roadtrip/eat [post]
func (ctrl *roadtripController) Eat(c *gin.Context) (*model.AppResult, *model.AppError) {
	var position model.CityFormData

	if err := c.ShouldBindJSON(&position); err != nil {
		return nil, &model.AppError{
			StatusCode: http.StatusBadRequest,
			Err:        err,
		}
	}

	location, err := ctrl.googleService.GeoCoding(ctrl.cfg.Google.BaseUrl, position.City)
	if err != nil {
		return nil, err.(*model.AppError)
	}

	restaurants, err := ctrl.googleService.Eat(ctrl.cfg.Google.BaseUrl, *location, position.Constraints)
	if err != nil {
		return nil, err.(*model.AppError)
	}

	return &model.AppResult{
		StatusCode: http.StatusOK,
		Message:    "Restaurants retrieved successfuly",
		Data:       &restaurants,
	}, nil
}

// Drink godoc
// @Summary Drink
// @Description Search for bars around the given city & constraints
// @Tags Roadtrip
// @Accept json
// @Produce json
// @Param cityFormData body model.CityFormData true "city & constraints"
// @Success 200 {object} model.DrinkSuccess "Success"
// @Failure 400 {object} model.InvalidJsonBody "Invalid request body"
// @Failure 401 {object} model.Unauthorized "Missing/Expired token"
// @Failure 404 {object} model.GoogleNotFound "Zero results"
// @Failure 500 {object} model.InternalServerError "Internal server error"
// @Router /roadtrip/drink [post]
func (ctrl *roadtripController) Drink(c *gin.Context) (*model.AppResult, *model.AppError) {
	var position model.CityFormData

	if err := c.ShouldBindJSON(&position); err != nil {
		return nil, &model.AppError{
			StatusCode: http.StatusBadRequest,
			Err:        err,
		}
	}

	location, err := ctrl.googleService.GeoCoding(ctrl.cfg.Google.BaseUrl, position.City)
	if err != nil {
		return nil, err.(*model.AppError)
	}

	bars, err := ctrl.googleService.Drink(ctrl.cfg.Google.BaseUrl, *location, position.Constraints)
	if err != nil {
		return nil, err.(*model.AppError)
	}

	return &model.AppResult{
		StatusCode: http.StatusOK,
		Message:    "Bars retrieved successfuly",
		Data:       &bars,
	}, nil
}

func (ctrl *roadtripController) Travel(ctx *gin.Context) (*model.AppResult, *model.AppError) {
	mode, _ := ctx.Get("mode")
	switch mode {
	case model.Air:
		return ctrl.TravelAir(ctx)
	case model.Ground:
		return ctrl.TravelGround(ctx)
	default:
		return nil, &model.AppError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("invalid travel mode"),
		}
	}
}

// Travel air godoc
// @Summary Travel air
// @Description Search for flight offers given a valid origin & destination
// @Tags Roadtrip
// @Accept json
// @Produce json
// @Param flightFormData body model.FlightFormData true "origin, destination, departure date, adults, max price"
// @Success 200 {object} model.TravelSuccess "Success"
// @Failure 400 {object} model.InvalidJsonBody "Invalid request body"
// @Failure 401 {object} model.Unauthorized "Missing/Expired token"
// @Failure 404 {object} model.FlightOfferNotFound "Flight offers not found"
// @Failure 500 {object} model.InternalServerError "Internal server error"
// @Router /roadtrip/travel/air [post]
func (ctrl *roadtripController) TravelAir(ctx *gin.Context) (*model.AppResult, *model.AppError) {
	var flightFormData model.FlightFormData

	if err := ctx.ShouldBindJSON(&flightFormData); err != nil {
		return nil, &model.AppError{
			StatusCode: http.StatusBadRequest,
			Err:        err,
		}
	}

	flightFormData.OriginLocationCode = utils.CityToIata(flightFormData.OriginLocation)
	flightFormData.DestinationLocationCode = utils.CityToIata(flightFormData.DestinationLocation)
	if flightFormData.OriginLocationCode == "" || flightFormData.DestinationLocationCode == "" {
		return nil, &model.AppError{
			StatusCode: http.StatusNotFound,
			Err:        errors.New("no airport found for origin/destination cities"),
		}
	}

	if time.Now().Unix() > int64(ctrl.amadeusAccessToken.Exp) {
		accessToken, err := ctrl.amadeusService.GetAccessToken(ctrl.cfg.Amadeus.BaseUrl)
		if err != nil {
			return nil, err.(*model.AppError)
		}
		ctrl.amadeusAccessToken = *accessToken
	}

	itineraries, err := ctrl.amadeusService.GetFlightOffers(ctrl.cfg.Amadeus.BaseUrl, ctrl.amadeusAccessToken.Value, &flightFormData)
	if err != nil {
		return nil, err.(*model.AppError)
	}

	return &model.AppResult{
		StatusCode: http.StatusOK,
		Message:    "Itineraries retrieved successfully",
		Data:       &itineraries,
	}, nil
}

// Travel ground godoc
// @Summary Travel ground
// @Description Search for train/bus directions given a valid origin & destination
// @Tags Roadtrip
// @Accept json
// @Produce json
// @Param directionsFormData body model.DirectionsFormData true "origin, destination, departure date"
// @Success 200 {object} model.TravelSuccess "Success"
// @Failure 400 {object} model.InvalidJsonBody "Invalid request body"
// @Failure 401 {object} model.Unauthorized "Missing/Expired token"
// @Failure 404 {object} model.GoogleNotFound "Zero results"
// @Failure 500 {object} model.InternalServerError "Internal server error"
// @Router /roadtrip/travel/ground [post]
func (ctrl *roadtripController) TravelGround(ctx *gin.Context) (*model.AppResult, *model.AppError) {
	var directionsFormData model.DirectionsFormData

	if err := ctx.ShouldBindJSON(&directionsFormData); err != nil {
		return nil, &model.AppError{
			StatusCode: http.StatusBadRequest,
			Err:        err,
		}
	}

	itineraries, err := ctrl.googleService.GetDirections(ctrl.cfg.Google.BaseUrl, &directionsFormData)
	if err != nil {
		return nil, err.(*model.AppError)
	}

	return &model.AppResult{
		StatusCode: http.StatusOK,
		Message:    "Itineraries retrieved successfully",
		Data:       &itineraries,
	}, nil
}
