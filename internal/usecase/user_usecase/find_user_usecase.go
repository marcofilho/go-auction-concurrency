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
}

func (u *UserUseCase) FindUserById(ctx context.Context, id string) (*UserOutputDTO, *internal_error.InternalError) {
	userEntity, err := u.UserRepositoryInterface.FindUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &UserOutputDTO{
		ID:   userEntity.ID,
		Name: userEntity.Name,
	}, nil
}
