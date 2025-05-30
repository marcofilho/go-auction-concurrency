package user

import (
	"context"

	"github.com/marcofilho/go-auction-concurrency/configuration/logger"
	"github.com/marcofilho/go-auction-concurrency/internal/entity/user_entity"
	"github.com/marcofilho/go-auction-concurrency/internal/internal_error"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserEntityMongo struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
}

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		Collection: db.Collection("users"),
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, userEntity *user_entity.User) *internal_error.InternalError {
	userEntityMongo := UserEntityMongo{
		ID:   userEntity.ID,
		Name: userEntity.Name,
	}

	_, err := r.Collection.InsertOne(ctx, userEntityMongo)
	if err != nil {
		logger.Error("Error creating user", err)
		return internal_error.NewInternalServerError("Error creating user")
	}

	return nil
}
