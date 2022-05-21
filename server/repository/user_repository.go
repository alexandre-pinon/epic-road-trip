package repository

import (
	"errors"

	"github.com/alexandre-pinon/epic-road-trip/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	db *mongo.Database
}

type UserRepository interface {
	CreateUser(user *model.User) error
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{db}
}

func (repository *userRepository) CreateUser(user *model.User) error {
	return errors.New("TODO: implement create user")
}
