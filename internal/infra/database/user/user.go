package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/marcofilho/go-auction-concurrency/configuration/logger"
	"github.com/marcofilho/go-auction-concurrency/internal/entity/user_entity"
	"github.com/marcofilho/go-auction-concurrency/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
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

func (r *UserRepository) CreateUser(ctx context.Context, user *user_entity.User) (*user_entity.User, *internal_error.InternalError) {
	userEntity := UserEntityMongo{
		ID:   user.ID,
		Name: user.Name,
	}

	_, err := r.Collection.InsertOne(ctx, userEntity)
	if err != nil {
		logger.Error("Error creating user", err)
		return nil, internal_error.NewInternalServerError("Error creating user")
	}

	return user, nil
}

func (r *UserRepository) FindUserById(ctx context.Context, id string) (*user_entity.User, *internal_error.InternalError) {
	filter := bson.M{"_id": id}

	var userEntity UserEntityMongo
	err := r.Collection.FindOne(ctx, filter).Decode(&userEntity)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error(fmt.Sprintf("User not found with this id = %s", id), err)
			return nil, internal_error.NewNotFoundError(
				fmt.Sprintf("User not found with this id = %s", id))
		}

		logger.Error(fmt.Sprintf("Error finding user with id = %s", id), err)
		return nil, internal_error.NewInternalServerError("Error finding user")
	}

	user := &user_entity.User{
		ID:   userEntity.ID,
		Name: userEntity.Name,
	}
	return user, nil

}
