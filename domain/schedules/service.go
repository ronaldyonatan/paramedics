package schedules

import (
	context "context"
	"fmt"
)

type repo interface {
	writescheduleRepo
	readScheduleRepo
}

type service struct {
	repo repo
}

type scheuleService struct {
	repo repo
}

func NewService(repo repo) service {
	return service{
		repo: repo,
	}
}

type writescheduleRepo interface {
	InsertSchedule(data paramedicSchedule) (result *ScheduleProto, err error)
	UpdateSchedule(data int32) (result *ScheduledetailProto, err error)
}

type readScheduleRepo interface {
	FindAll(ctx context.Context, req FindAll) (data *ListscheduledetailProto, err error)
	FindBooked(ctx context.Context, req FindBooked) (data *ListscheduledetailProto, err error)
	FindAvailable(ctx context.Context, req FindAvailable) (data *ListscheduledetailProto, err error)
}

func (s service) CreateSchedule(ctx context.Context, req scheduleCreateRequest) (result *ScheduleProto, err error) {
	fmt.Println("inside service")

	result, err = s.repo.InsertSchedule(req.ConvertToSchedule())
	if err != nil {
		return nil, err
	}

	return
}

func (s service) UpdateSchedule(ctx context.Context, req scheduleSetRequest) (result *ScheduledetailProto, err error) {
	fmt.Println("inside service")

	result, err = s.repo.UpdateSchedule(req.ScheduleSlotId)
	if err != nil {
		return nil, err
	}

	return
}

func (s service) FindAll(ctx context.Context, req FindAll) (data *ListscheduledetailProto, err error) {
	data, err = s.repo.FindAll(ctx, req)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, err
	}

	return
}

func (s service) FindAvailable(ctx context.Context, req FindAvailable) (data *ListscheduledetailProto, err error) {
	data, err = s.repo.FindAvailable(ctx, req)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, err
	}

	return
}

func (s service) FindBooked(ctx context.Context, req FindBooked) (data *ListscheduledetailProto, err error) {
	data, err = s.repo.FindBooked(ctx, req)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, err
	}

	return
}
