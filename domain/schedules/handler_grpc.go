package schedules

import (
	context "context"
)

type scheduleHandlerGrpc struct {
	svc service
}

func NewScheduleHandlerGPRC(svc service) scheduleHandlerGrpc {
	return scheduleHandlerGrpc{
		svc: svc,
	}
}

func (a scheduleHandlerGrpc) CreateSchedule(ctx context.Context, req *ScheduleCreateProto) (result *ScheduleProto, err error) {
	reqPayload := scheduleCreateRequest{
		ScheduleStart: req.Schedulestart.AsTime(),
		ScheduleEnd:   req.Scheduleend.AsTime(),
		ParamedicId:   req.Paramedicid,
		HealthcareId:  req.Usercreate,
		Duration:      int(req.Duration),
		UserCreate:    req.Usercreate,
	}

	result, err = a.svc.CreateSchedule(ctx, reqPayload)

	return
}
func (a scheduleHandlerGrpc) FindScheduleAll(ctx context.Context, req *FindScheduleProto) (result *ListscheduledetailProto, err error) {
	reqPayload := FindAll{
		ParamedicId:  req.Paramedicid,
		HealthcareId: req.Healthcareid,
		Scheduledate: req.Scheduledate.AsTime(),
	}

	data, err := a.svc.FindAll(ctx, reqPayload)

	result = data
	// for _, v := range result {
	// 	data.Paramedics = append(data.Paramedics, &v)

	// }

	// emp = new(emptypb.Empty)
	return data, err
}
func (a scheduleHandlerGrpc) FindScheduleAvailable(ctx context.Context, req *FindScheduleProto) (result *ListscheduledetailProto, err error) {
	reqPayload := FindAvailable{
		ParamedicId:  req.Paramedicid,
		HealthcareId: req.Healthcareid,
		Scheduledate: req.Scheduledate.AsTime(),
	}

	data, err := a.svc.FindAvailable(ctx, reqPayload)

	result = data
	// for _, v := range result {
	// 	data.Paramedics = append(data.Paramedics, &v)

	// }

	// emp = new(emptypb.Empty)
	return data, err
}
func (a scheduleHandlerGrpc) FindScheduleBooked(ctx context.Context, req *FindScheduleProto) (result *ListscheduledetailProto, err error) {
	reqPayload := FindBooked{
		ParamedicId:  req.Paramedicid,
		HealthcareId: req.Healthcareid,
		Scheduledate: req.Scheduledate.AsTime(),
	}

	data, err := a.svc.FindBooked(ctx, reqPayload)

	result = data
	// for _, v := range result {
	// 	data.Paramedics = append(data.Paramedics, &v)

	// }

	// emp = new(emptypb.Empty)
	return data, err
}
func (a scheduleHandlerGrpc) BookSlot(ctx context.Context, req *SetScheduleProto) (result *ScheduledetailProto, err error) {
	reqPayload := scheduleSetRequest{
		ScheduleSlotId: req.Scheduleslotid,
	}

	result, err = a.svc.UpdateSchedule(ctx, reqPayload)

	return
}
func (a scheduleHandlerGrpc) mustEmbedUnimplementedScheduleServer() {}
