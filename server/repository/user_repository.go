package repository

import (
	"context"

	"github.com/alexandre-pinon/epic-road-trip/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	db   *mongo.Database
	coll *mongo.Collection
}

type UserRepository interface {
	CreateUser(user *model.User) error
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{db, db.Collection("user")}
}

func (repo *userRepository) CreateUser(user *model.User) error {
	_, err := repo.coll.InsertOne(context.Background(), user)
	return err
}
