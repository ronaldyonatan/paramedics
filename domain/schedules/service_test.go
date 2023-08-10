package schedules

import (
	context "context"
	"fmt"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func TestCreateSchedule(t *testing.T) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Jakarta",
		"172.16.20.40",
		"5432",
		"sa",
		"P@ssw0rd",
		"HParamedic",
		"disable",
	)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("successfull connect db")

	start := "2023-08-01 10:00"
	layout := "2006-01-02 15:04"
	parsedTimeStart, err := time.Parse(layout, start)

	end := "2023-08-01 12:00"
	layout2 := "2006-01-02 15:04"
	parsedTimeEnd, err := time.Parse(layout2, end)

	fmt.Println(parsedTimeStart)
	fmt.Println(parsedTimeEnd)

	req := scheduleCreateRequest{
		HealthcareId:  "001",
		ParamedicId:   "D00001",
		ScheduleStart: parsedTimeStart,
		ScheduleEnd:   parsedTimeEnd,
		Duration:      6,
		UserCreate:    "ronald",
	}

	serviceInstance := &service{repo: NewRepo(db)}

	t.Run("Successful Insertion", func(t *testing.T) {
		fmt.Println("Start Test")
		result, err := serviceInstance.CreateSchedule(context.Background(), req)
		if err != nil {
			panic(err.Error())
		}
		if result == nil {
			panic("no result")
		}
		fmt.Println(result)
	})
}

func TestFindAll(t *testing.T) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Jakarta",
		"172.16.20.40",
		"5432",
		"sa",
		"P@ssw0rd",
		"HParamedic",
		"disable",
	)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("successfull connect db")

	tgl := "2023-08-01 00:00"
	layout := "2006-01-02 15:04"

	parsedTime, err := time.Parse(layout, tgl)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	req := FindAll{
		ParamedicId:  "D00001",
		HealthcareId: "001",
		Scheduledate: parsedTime,
	}

	serviceInstance := &service{repo: NewRepo(db)}

	t.Run("Successful Insertion", func(t *testing.T) {
		result, err := serviceInstance.FindAll(context.Background(), req)
		if err != nil {
			panic(err.Error())
		}
		if result == nil {
			panic("no result")
		}
		fmt.Println(result)

	})
}
