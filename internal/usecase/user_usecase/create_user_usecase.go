package user_usecase

import (
	"context"

	"github.com/marcofilho/go-auction-concurrency/internal/entity/user_entity"
	"github.com/marcofilho/go-auction-concurrency/internal/internal_error"
)

func (u *UserUseCase) CreateUser(ctx context.Context, name string) *internal_error.InternalError {
	user, err := user_entity.CreateUser(name)
	if err != nil {
		return err
	}
	if err := u.UserRepositoryInterface.CreateUser(ctx, user); err != nil {
		return err
	}

	return nil
}
