package repository

import (
	"context"

	"github.com/alexandre-pinon/epic-road-trip/model"
	"github.com/alexandre-pinon/epic-road-trip/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	db   *mongo.Database
	coll *mongo.Collection
}

type UserRepository interface {
	GetAllUsers() (*[]model.User, error)
	GetUserByID(id primitive.ObjectID, populate bool) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	CreateUser(user *model.User) (*mongo.InsertOneResult, error)
	UpdateUser(id primitive.ObjectID, user *model.User) (*mongo.UpdateResult, error)
	DeleteUser(id primitive.ObjectID) (*mongo.DeleteResult, error)
}

func NewUserRepository(db *mongo.Database) UserRepository {
	coll := db.Collection("user")
	indexes := []mongo.IndexModel{{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}, {
		Keys:    bson.D{{Key: "phone", Value: 1}},
		Options: options.Index().SetUnique(true),
	}}

	utils.AddIndexes(coll, indexes)

	return &userRepository{db, coll}
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

	return &results, nil
}

func (repo *userRepository) GetUserByID(id primitive.ObjectID, populate bool) (*model.User, error) {
	var user model.User
	filter := bson.D{{Key: "_id", Value: id}}

	if !populate {
		result := repo.coll.FindOne(context.Background(), filter)

		if err := result.Decode(&user); err != nil {
			return nil, err
		}
		return &user, nil
	}

	aggSearch := bson.M{"$match": filter}
	aggUnwind1 := bson.M{"$unwind": bson.M{"path": "$trips", "preserveNullAndEmptyArrays": true}}
	aggPopulate := bson.M{"$lookup": bson.M{
		"from":         "tripStep",           // the collection name
		"localField":   "trips.tripSteps_id", // the field on the child struct
		"foreignField": "_id",                // the field on the parent struct
		"as":           "trips.tripSteps",    // the field to populate into
	}}
	aggGroup := bson.M{"$group": bson.M{
		"_id":            "$_id",
		"firstname":      bson.M{"$first": "$firstname"},
		"lastname":       bson.M{"$first": "$lastname"},
		"email":          bson.M{"$first": "$email"},
		"hashedpassword": bson.M{"$first": "$hashedpassword"},
		"phone":          bson.M{"$first": "$phone"},
		"trips":          bson.M{"$push": "$trips"},
	}}

	agg := []bson.M{aggSearch, aggUnwind1, aggPopulate, aggGroup}
	cursor, err := repo.coll.Aggregate(context.Background(), agg)
	if err != nil {
		return nil, err
	}

	if cursor.Next(context.Background()) {
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}

		roadtrips := user.Trips
		if len(roadtrips) == 1 {
			if len(*roadtrips[0].TripSteps) == 0 {
				user.Trips = nil
			}
		}
	}

	return &user, nil
}

func (repo *userRepository) GetUserByEmail(email string) (*model.User, error) {
	filter := bson.D{{Key: "email", Value: email}}
	result := repo.coll.FindOne(context.Background(), filter)

	var user model.User
	if err := result.Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *userRepository) CreateUser(user *model.User) (*mongo.InsertOneResult, error) {
	return repo.coll.InsertOne(context.Background(), user)
}

func (repo *userRepository) UpdateUser(id primitive.ObjectID, user *model.User) (*mongo.UpdateResult, error) {
	update := bson.D{{Key: "$set", Value: user}}

	return repo.coll.UpdateByID(context.Background(), id, update)
}

func (repo *userRepository) DeleteUser(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.D{{Key: "_id", Value: id}}

	return repo.coll.DeleteOne(context.Background(), filter)
}
