package paramedics

import (
	"time"
)

type paramedic struct {
	ParamedicID string    `db:"paramedic_id"`
	FirstName   string    `db:"first_name"`
	LastName    string    `db:"last_name"`
	Email       string    `db:"email"`
	IsActive    bool      `db:"is_active"`
	CreatedAt   time.Time `db:"create_at"`
	UserCreate  string    `db:"user_create"`
}
