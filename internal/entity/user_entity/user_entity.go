package user_entity

import (
	"context"

	"github.com/google/uuid"
	"github.com/marcofilho/go-auction-concurrency/internal/internal_error"
)

type User struct {
	ID   string
	Name string
}

type UserRepositoryInterface interface {
	FindUserById(ctx context.Context, id string) (*User, *internal_error.InternalError)
	CreateUser(ctx context.Context, userEntity *User) *internal_error.InternalError
}

func CreateUser(name string) (*User, *internal_error.InternalError) {
	user := &User{
		ID:   uuid.NewString(),
		Name: name,
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	return user, nil
}

func (a *User) Validate() *internal_error.InternalError {
	if len(a.Name) <= 3 {
		return internal_error.NewBadRequestError("Invalid user data")
	}
	return nil
}
