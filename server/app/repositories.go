package app

import (
	"github.com/alexandre-pinon/epic-road-trip/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repositories struct {
	UserRepository repository.UserRepository
}

func SetupRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		UserRepository: repository.NewUserRepository(db),
	}
}
