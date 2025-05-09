package user_usecase

import (
	"context"

	"github.com/marcofilho/go-auction-concurrency/internal/internal_error"
)

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
