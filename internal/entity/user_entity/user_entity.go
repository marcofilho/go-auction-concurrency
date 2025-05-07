package user_entity

import (
	"context"

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
