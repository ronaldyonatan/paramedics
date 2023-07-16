package users

import (
	"database/sql"
	"time"
)

type auth struct {
	Id          uint         `db:"id"`
	Email       string       `db:"email"`
	Password    string       `db:"password"`
	Username    string       `db:"username"`
	FirstName   string       `db:"first_name"`
	LastName    string       `db:"last_name"`
	IsActive    bool         `db:"is_active"`
	CreatedAt   time.Time    `db:"created_at"`
	ActivatedAt sql.NullTime `db:"activated_at"`
}

func (a auth) ConvertToAuthSignInResponse() authSignInResponse {
	return authSignInResponse{
		Email:     a.Email,
		Username:  a.Username,
		FirstName: a.FirstName,
		LastName:  a.LastName,
	}
}
