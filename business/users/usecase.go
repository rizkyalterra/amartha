package users

import (
	"context"
	"errors"
	"time"
)

type UserUsecase struct {
	userRepository Repository
	contextTimeout time.Duration
}

func (userUseCase *UserUsecase) Login(ctx context.Context, email, password string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, userUseCase.contextTimeout)
	defer cancel()

	user, err := userUseCase.userRepository.Login(ctx, email, password)
	if err != nil {
		return Domain{}, errors.New("Gagal login")
	}
	return user, nil
}

func NewUserUsecase(repository Repository, timeout time.Duration) Usecase {
	return &UserUsecase{
		userRepository: repository,
		contextTimeout: timeout,
	}
}
