package schedules

import (
	"time"
)

type paramedicSchedule struct {
	HealthcareID  string    `db:"healthcare_id"`
	ParamedicID   string    `db:"paramedic_id"`
	ScheduleStart time.Time `db:"schedule_start"`
	ScheduleEnd   time.Time `db:"schedule_end"`
	Duration      int       `db:"duration"`
	IsActive      bool      `db:"is_active"`
	CreatedAt     time.Time `db:"create_at"`
	UserCreate    string    `db:"user_create"`
}

type paramedicScheduleDetail struct {
	ScheduleDetailID int       `db:"schedule_slot_id"`
	SlotTime         time.Time `db:"slot_time"`
	IsBooked         bool      `db:"is_booked"`
	CreatedAt        time.Time `db:"create_at"`
	UserCreate       string    `db:"user_create"`
}
