package repository

import (
	"errors"

	"github.com/alexandre-pinon/epic-road-trip/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type tripStepRepository struct {
	db   *mongo.Database
	coll *mongo.Collection
}

type TripStepRepository interface {
	GetAllTripSteps() (*[]model.TripStep, error)
	GetTripStepByID(id primitive.ObjectID) (*model.TripStep, error)
	CreateTripStep(tripStep *model.TripStep) (*mongo.InsertOneResult, error)
	DeleteTripStep(id primitive.ObjectID) (*mongo.DeleteResult, error)
}

func NewTripStepRepository(db *mongo.Database) TripStepRepository {
	coll := db.Collection("tripStep")

	return &tripStepRepository{db, coll}
}

func (repo *tripStepRepository) GetAllTripSteps() (*[]model.TripStep, error) {
	return &[]model.TripStep{}, errors.New("TODO: implement GetAllTripSteps")
}
func (repo *tripStepRepository) GetTripStepByID(id primitive.ObjectID) (*model.TripStep, error) {
	return &model.TripStep{}, errors.New("TODO: implement GetTripStepByID")
}
func (repo *tripStepRepository) CreateTripStep(tripStep *model.TripStep) (*mongo.InsertOneResult, error) {
	return &mongo.InsertOneResult{}, errors.New("TODO: implement CreateTripStep")
}
func (repo *tripStepRepository) DeleteTripStep(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	return &mongo.DeleteResult{}, errors.New("TODO: implement DeleteTripStep")
}
