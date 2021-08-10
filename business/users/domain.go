package users

import (
	"context"
	"time"
)

type Domain struct {
	Id        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Name      string
	Email     string
	Alamat    string
	Password  string
}

type Usecase interface {
	Login(ctx context.Context, email, password string) (Domain, error)
}

type Repository interface {
	Login(ctx context.Context, email, password string) (Domain, error)
}
