package user_usecase

import (
	"context"

	"github.com/marcofilho/go-auction-concurrency/internal/entity/user_entity"
	"github.com/marcofilho/go-auction-concurrency/internal/internal_error"
)

type UserUseCase struct {
	UserRepositoryInterface user_entity.UserRepositoryInterface
}

type UserOutputDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserUseCaseInterface interface {
	FindUserById(ctx context.Context, id string) (*UserOutputDTO, *internal_error.InternalError)
	CreateUser(ctx context.Context, name string) *internal_error.InternalError
}

func NewUserUseCase(userRepository user_entity.UserRepositoryInterface) UserUseCaseInterface {
	return &UserUseCase{
		UserRepositoryInterface: userRepository,
	}
}

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
