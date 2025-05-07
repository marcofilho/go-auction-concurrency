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

	return &user_entity.User{
		ID:   userEntity.ID,
		Name: userEntity.Name,
	}, nil

}
