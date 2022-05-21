package repository

import (
	"context"
	"errors"

	"github.com/alexandre-pinon/epic-road-trip/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	db   *mongo.Database
	coll *mongo.Collection
}

type UserRepository interface {
	GetAllUsers() (*[]model.User, error)
	CreateUser(user *model.User) error
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{db, db.Collection("user")}
}

func (repo *userRepository) GetAllUsers() (*[]model.User, error) {
	return &[]model.User{}, errors.New("TODO: implement GetAllUsers")
}

func (repo *userRepository) CreateUser(user *model.User) error {
	_, err := repo.coll.InsertOne(context.Background(), user)
	return err
}
