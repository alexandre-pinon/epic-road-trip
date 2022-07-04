package repository

import (
	"errors"

	"github.com/alexandre-pinon/epic-road-trip/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type roadtripRepository struct {
	db   *mongo.Database
	coll *mongo.Collection
}

type RoadtripRepository interface {
	GetAllRoadtrips() (*[]model.Roadtrip, error)
	GetRoadtripByID(id primitive.ObjectID) (*model.Roadtrip, error)
	CreateRoadtrip(roadtrip *model.Roadtrip) (*mongo.InsertOneResult, error)
	DeleteRoadtrip(id primitive.ObjectID) (*mongo.DeleteResult, error)
}

func NewRoadtripRepository(db *mongo.Database) RoadtripRepository {
	coll := db.Collection("roadtrip")

	return &roadtripRepository{db, coll}
}

func (repo *roadtripRepository) GetAllRoadtrips() (*[]model.Roadtrip, error) {
	return &[]model.Roadtrip{}, errors.New("TODO: implement GetAllRoadtrips")
}
func (repo *roadtripRepository) GetRoadtripByID(id primitive.ObjectID) (*model.Roadtrip, error) {
	return &model.Roadtrip{}, errors.New("TODO: implement GetRoadtripByID")
}
func (repo *roadtripRepository) CreateRoadtrip(roadtrip *model.Roadtrip) (*mongo.InsertOneResult, error) {
	return &mongo.InsertOneResult{}, errors.New("TODO: implement CreateRoadtrip")
}
func (repo *roadtripRepository) DeleteRoadtrip(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	return &mongo.DeleteResult{}, errors.New("TODO: implement DeleteRoadtrip")
}
