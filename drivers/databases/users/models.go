package users

import (
	"pembayaran/business/users"
	"time"
)

type User struct {
	Id        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `gorm:"index" json:"deletedAt"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Alamat    string    `json:"alamat"`
	Password  string    `json:"-"`
}

func fromDomain(domain *users.Domain) *User {
	return &User{
		Id:        domain.Id,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		Name:      domain.Name,
		Email:     domain.Email,
		Alamat:    domain.Alamat,
		Password:  domain.Password,
	}
}

func (user *User) toDomain() users.Domain {
	return users.Domain{
		Id:        user.Id,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
		Name:      user.Name,
		Email:     user.Email,
		Alamat:    user.Alamat,
		Password:  user.Password,
	}
}
