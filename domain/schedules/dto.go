package schedules

import "time"

type scheduleCreateRequest struct {
	HealthcareId  string    `json:"healthcare_id" validate:"required"`
	ParamedicId   string    `json:"paramedic_id" validate:"required"`
	ScheduleStart time.Time `json:"schedule_start" validate:"required"`
	ScheduleEnd   time.Time `json:"schedule_end" validate:"required"`
	Duration      int       `json:"duration" validate:"required"`
	UserCreate    string    `json:"user_create" validate:"required"`
}

type scheduleSetRequest struct {
	ScheduleSlotId int32 `json:"schedule_slot_id" validate:"required"`
}

type FindAll struct {
	HealthcareId string    `json:"healthcare_id" validate:"required"`
	ParamedicId  string    `json:"paramedic_id" validate:"required"`
	Scheduledate time.Time `json:"schedule_date" validate:"required"`
}

type FindAvailable struct {
	HealthcareId string    `json:"healthcare_id" validate:"required"`
	ParamedicId  string    `json:"paramedic_id" validate:"required"`
	Scheduledate time.Time `json:"schedule_date" validate:"required"`
}

type FindBooked struct {
	HealthcareId string    `json:"healthcare_id" validate:"required"`
	ParamedicId  string    `json:"paramedic_id" validate:"required"`
	Scheduledate time.Time `json:"schedule_date" validate:"required"`
}

func (a scheduleCreateRequest) ConvertToSchedule() paramedicSchedule {
	return paramedicSchedule{
		HealthcareID:  a.HealthcareId,
		ParamedicID:   a.ParamedicId,
		ScheduleStart: a.ScheduleStart,
		ScheduleEnd:   a.ScheduleEnd,
		Duration:      a.Duration,
		UserCreate:    a.UserCreate,
	}
}
