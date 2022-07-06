package repository

import (
	"context"
	"time"

	"github.com/alexandre-pinon/epic-road-trip/model"
	"go.mongodb.org/mongo-driver/bson"
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
	ctx := context.Background()

	cursor, err := repo.coll.Find(ctx, struct{}{})
	if err != nil {
		return nil, err
	}

	var results []model.TripStep
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return &results, nil
}
func (repo *tripStepRepository) GetTripStepByID(id primitive.ObjectID) (*model.TripStep, error) {
	filter := bson.D{{Key: "_id", Value: id}}
	result := repo.coll.FindOne(context.Background(), filter)

	var tripStep model.TripStep
	if err := result.Decode(&tripStep); err != nil {
		return nil, err
	}

	return &tripStep, nil
}
func (repo *tripStepRepository) CreateTripStep(tripStep *model.TripStep) (*mongo.InsertOneResult, error) {
	if tripStep.Travel != nil {
		duration, err := time.ParseDuration(tripStep.Travel.DurationString)
		if err != nil {
			return nil, err
		}

		tripStep.Travel.Duration = duration
	}

	return repo.coll.InsertOne(context.Background(), tripStep)
}
func (repo *tripStepRepository) DeleteTripStep(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.D{{Key: "_id", Value: id}}

	return repo.coll.DeleteOne(context.Background(), filter)
}
