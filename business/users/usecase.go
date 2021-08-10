package users

import (
	"context"
	"pembayaran/helpers/middlewares"
	"time"
)

type UserUseCase struct {
	Repository     Repository
	ContextTimeout time.Duration
}

func NewUserUsecase(repository Repository, timeout time.Duration) UseCase {
	return &UserUseCase{
		Repository:     repository,
		ContextTimeout: timeout,
	}
}

func (usecase *UserUseCase) Login(ctx context.Context, email, password string) (Domain, error) {
	result, err := usecase.Repository.Login(ctx, email, password)
	if err != nil {
		return Domain{}, err
	}
	result.Token, _ = middlewares.GenerateTokenJWT(int32(result.Id))
	return result, nil
}
