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

func (rec *User) toDomain() users.Domain {
	return users.Domain{
		Id:        rec.Id,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
		Name:      rec.Name,
		Email:     rec.Email,
		Alamat:    rec.Alamat,
		Password:  rec.Password,
	}
}

func fromDomain(user users.Domain) *User {
	return &User{
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
