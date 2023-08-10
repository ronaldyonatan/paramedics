package paramedics

import (
	context "context"
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func TestCreateParamedic(t *testing.T) {
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

	req := paramedicCreateRequest{
		FirstName:  "donald",
		LastName:   "duck",
		Email:      "mickey@mouse.com",
		UserCreate: "ronald",
	}

	serviceInstance := &service{repo: NewRepo(db)}

	t.Run("Successful Insertion", func(t *testing.T) {
		fmt.Println("Start Test")
		result, err := serviceInstance.CreateParamedic(context.Background(), req)
		if err != nil {
			panic(err.Error())
		}
		if result == nil {
			panic("no result")
		}

	})
}

func TestFindByHospital(t *testing.T) {
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

	req := FindByHospitalRequest{
		HospitalId: "001",
	}

	serviceInstance := &service{repo: NewRepo(db)}

	t.Run("Successful Insertion", func(t *testing.T) {
		fmt.Println("Start Test")
		result, err := serviceInstance.FindByHospital(context.Background(), req)
		if err != nil {
			panic(err.Error())
		}
		if result == nil {
			panic("no result")
		}

	})
}
