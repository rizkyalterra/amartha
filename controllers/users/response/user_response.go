package response

import (
	"pembayaran/business/users"
	"time"
)

type User struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Alamat    string    `json:"alamat"`
	Token     string    `json:"token"`
}

func FromDomain(domain users.Domain) User {
	return User{
		Id:        domain.Id,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		Name:      domain.Name,
		Email:     domain.Email,
		Alamat:    domain.Alamat,
		Token:     domain.Token,
	}
}
