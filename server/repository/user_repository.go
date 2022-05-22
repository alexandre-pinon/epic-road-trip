package repository

import (
	"context"

	"github.com/alexandre-pinon/epic-road-trip/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	db   *mongo.Database
	coll *mongo.Collection
}

type UserRepository interface {
	GetAllUsers() (*[]model.User, error)
	GetUserByID(id primitive.ObjectID) (*model.User, error)
	CreateUser(user *model.User) (*mongo.InsertOneResult, error)
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{db, db.Collection("user")}
}

func (repo *userRepository) GetAllUsers() (*[]model.User, error) {
	ctx := context.Background()

	cursor, err := repo.coll.Find(ctx, struct{}{})
	if err != nil {
		return nil, err
	}

	var results []model.User
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return &results, err
}

func (repo *userRepository) GetUserByID(id primitive.ObjectID) (*model.User, error) {
	filter := bson.D{{Key: "_id", Value: id}}
	result := repo.coll.FindOne(context.Background(), filter)

	var user *model.User
	err := result.Decode(&user)

	return user, err
}

func (repo *userRepository) CreateUser(user *model.User) (*mongo.InsertOneResult, error) {
	return repo.coll.InsertOne(context.Background(), user)
}
