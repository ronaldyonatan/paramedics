package paramedics

import (
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
)

func TestInsertParamedic(t *testing.T) {
	// Connect to the test database
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
	repoInstance := &paramedicRepo{db: db}

	if err != nil {
		panic("fail connect db")
	}

	data := paramedic{
		FirstName:  "John",
		LastName:   "Doe",
		Email:      "john@example.com",
		UserCreate: "admin",
	}

	result, err := repoInstance.InsertParamedic(data)
	if err != nil {
		panic("error insert data")
	}
	if result.Paramedicid == "" {
		panic("no data result found")
	}
}
