package users

import (
	"database/sql"
	"time"
)

type authSignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (a authSignInRequest) ConvertToAuth() auth {
	return auth{
		Email:    a.Email,
		Password: a.Password,
	}
}

type authSignInResponse struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
type verifyAuth struct {
	Id          uint         `db:"id"`
	AuthId      uint         `db:"auth_id"`
	Token       string       `db:"token"`
	CreatedAt   time.Time    `db:"created_at"`
	ExpiredAt   time.Time    `db:"expired_at"`
	ActivatedAt sql.NullTime `db:"activated_at"`
}

type authCreateRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	Username  string `json:"username" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}

func (a authCreateRequest) ConvertToAuth() auth {
	return auth{
		Email:     a.Email,
		Password:  a.Password,
		Username:  a.Username,
		FirstName: a.FirstName,
		LastName:  a.LastName,
	}
}
