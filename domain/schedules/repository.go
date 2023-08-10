package schedules

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type scheduleRepo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) scheduleRepo {
	return scheduleRepo{db}
}

func (r scheduleRepo) InsertSchedule(data paramedicSchedule) (result *ScheduleProto, err error) {
	fmt.Println("inside repo")
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback() // Rollback if an error occurs
	query := `insert into paramedicschedules (healthcare_id, paramedic_id, schedule_start, schedule_end, 
						duration,is_active,user_create, create_at)
				values ($1, $2, $3,$4,$5, true,$6,$7) returning schedule_id;`

	var insertedID string
	fmt.Println("insert data pertama")
	err = r.db.QueryRow(query, data.HealthcareID, data.ParamedicID, data.ScheduleStart, data.ScheduleEnd, data.Duration, data.UserCreate, time.Now()).Scan(&insertedID)
	if err != nil {
		return
	}
	fmt.Println(insertedID)

	//durationInterval := time.Minute * time.Duration(data.Duration)
	durationInterval := strconv.Itoa(data.Duration) + " minutes"

	fmt.Println(durationInterval)

	query2 := `insert into paramedicscheduleslots(
				schedule_id, slot_time, is_booked, is_active, user_create, create_at	
			)
			select $1, generate_series(
				$2::timestamp,
				$3::timestamp,
				$4::interval 
			) AS time_slot, false, true, $5, $6`

	fmt.Println("insert data kedua")

	fmt.Println(insertedID)
	fmt.Println(data.ScheduleStart)
	fmt.Println(data.ScheduleEnd)
	fmt.Println(durationInterval)
	fmt.Println(data.UserCreate)

	_, err = r.db.Exec(query2, insertedID, data.ScheduleStart, data.ScheduleEnd, durationInterval, data.UserCreate, time.Now())
	if err != nil {
		return
	}
	fmt.Println("berhasil insert data kedua")

	err = tx.Commit()
	if err != nil {
		return
	}

	result = &ScheduleProto{ // Initialize the result variable
		Scheduleid:    insertedID,
		Healthcareid:  data.HealthcareID,
		Paramedicid:   data.ParamedicID,
		Schedulestart: data.ScheduleStart.String(),
		Scheduleend:   data.ScheduleEnd.String(),
		Duration:      int32(data.Duration),
	}

	return
}

func (r scheduleRepo) UpdateSchedule(slotid int32) (result *ScheduledetailProto, err error) {
	fmt.Println("inside repo")
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback() // Rollback if an error occurs
	query := `update paramedicscheduleslots
		set is_booked = true 
		where schedule_slot_id =  $1`

	_, err = r.db.Exec(query, slotid)
	if err != nil {
		return
	}

	query2 := `select slot_time 
		from  paramedicscheduleslots
		where schedule_slot_id =  $1`

	var SlotTime time.Time

	err = r.db.QueryRow(query2, slotid).Scan(&SlotTime)
	if err != nil {
		return
	}

	fmt.Println(SlotTime)

	result = &ScheduledetailProto{ // Initialize the result variable
		Scheduleslotid: slotid,
		Slottime:       timestamppb.New(SlotTime),
		Isbooked:       true,
	}

	return
}

func (r scheduleRepo) FindAll(ctx context.Context, req FindAll) (data *ListscheduledetailProto, err error) {
	dataProto := &ListscheduledetailProto{
		Schedules: make([]*ScheduledetailProto, 0), // Initialize the slice
	}

	var schedules []paramedicScheduleDetail

	query := `select c.schedule_slot_id, slot_time, is_booked
			  from paramedicschedules a 
			  left join paramedics b
					on b.paramedic_id = a.paramedic_id
			  inner join paramedicscheduleslots c
			  	on c.schedule_id = a.schedule_id
			  where healthcare_id=$1
			 	and a.paramedic_id = $2
				and date(a.schedule_start) = $3 
			  `
	fmt.Println(query)

	err = r.db.Select(&schedules, query, req.HealthcareId, req.ParamedicId, req.Scheduledate)

	for _, schedule := range schedules {
		fmt.Printf("ID: %v", schedule.ScheduleDetailID)

		result := &ScheduledetailProto{ // Initialize the result variable
			Scheduleslotid: int32(schedule.ScheduleDetailID),
			Slottime:       timestamppb.New(schedule.SlotTime),
			Isbooked:       schedule.IsBooked,
		}
		dataProto.Schedules = append(dataProto.Schedules, result)
	}

	if err != nil {
		return nil, err
	}

	return dataProto, err
}

func (r scheduleRepo) FindAvailable(ctx context.Context, req FindAvailable) (data *ListscheduledetailProto, err error) {
	dataProto := &ListscheduledetailProto{
		Schedules: make([]*ScheduledetailProto, 0), // Initialize the slice
	}

	var schedules []paramedicScheduleDetail

	query := `select c.schedule_slot_id, slot_time, is_booked
			  from paramedicschedules a 
			  left join paramedics b
					on b.paramedic_id = a.paramedic_id
			  inner join paramedicscheduleslots c
			  	on c.schedule_id = a.schedule_id
			  where healthcare_id=$1
			 	and a.paramedic_id = $2
				and date(a.schedule_start) = $3 
				and c.is_booked = false
			  `
	fmt.Println(query)

	err = r.db.Select(&schedules, query, req.HealthcareId, req.ParamedicId, req.Scheduledate)

	for _, schedule := range schedules {
		fmt.Printf("ID: %v", schedule.ScheduleDetailID)

		result := &ScheduledetailProto{ // Initialize the result variable
			Scheduleslotid: int32(schedule.ScheduleDetailID),
			Slottime:       timestamppb.New(schedule.SlotTime),
			Isbooked:       schedule.IsBooked,
		}
		dataProto.Schedules = append(dataProto.Schedules, result)
	}

	if err != nil {
		return nil, err
	}

	return dataProto, err
}

func (r scheduleRepo) FindBooked(ctx context.Context, req FindBooked) (data *ListscheduledetailProto, err error) {
	dataProto := &ListscheduledetailProto{
		Schedules: make([]*ScheduledetailProto, 0), // Initialize the slice
	}

	var schedules []paramedicScheduleDetail

	query := `select c.schedule_slot_id, slot_time, is_booked
			  from paramedicschedules a 
			  left join paramedics b
					on b.paramedic_id = a.paramedic_id
			  inner join paramedicscheduleslots c
			  	on c.schedule_id = a.schedule_id
			  where healthcare_id=$1
			 	and a.paramedic_id = $2
				and date(a.schedule_start) = $3 
				and c.is_booked = true
			  `
	fmt.Println(query)

	err = r.db.Select(&schedules, query, req.HealthcareId, req.ParamedicId, req.Scheduledate)

	for _, schedule := range schedules {
		fmt.Printf("ID: %v", schedule.ScheduleDetailID)

		result := &ScheduledetailProto{ // Initialize the result variable
			Scheduleslotid: int32(schedule.ScheduleDetailID),
			Slottime:       timestamppb.New(schedule.SlotTime),
			Isbooked:       schedule.IsBooked,
		}
		dataProto.Schedules = append(dataProto.Schedules, result)
	}

	if err != nil {
		return nil, err
	}

	return dataProto, err
}
