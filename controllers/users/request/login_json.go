package request

import "pembayaran/business/users"

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *LoginUser) ToDomain() *users.Domain {
	return &users.Domain{
		Email:    req.Email,
		Password: req.Password,
	}
}
