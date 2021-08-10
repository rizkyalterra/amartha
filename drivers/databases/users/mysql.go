package users

import (
	"context"
	"pembayaran/business/users"

	"gorm.io/gorm"
)

type MysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(conn *gorm.DB) users.Repository {
	return &MysqlUserRepository{
		Conn: conn,
	}
}

func (repository *MysqlUserRepository) Login(ctx context.Context, email, password string) (users.Domain, error) {
	var userLogin = User{}
	result := repository.Conn.Where("email = ? AND password = ?", email, password).First(&userLogin)
	if result.Error != nil {
		return users.Domain{}, result.Error
	}
	return userLogin.toDomain(), nil
}
