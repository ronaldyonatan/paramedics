package paramedics

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type paramedicRepo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) paramedicRepo {
	return paramedicRepo{db}
}

func (r paramedicRepo) InsertParamedic(data paramedic) (result *ParamedicProto, err error) {
	fmt.Println("insert Paramedic Repo")

	query := `insert into paramedics (paramedic_id, first_name, last_name, email, is_active, user_create, create_at) 
				select concat('D',right(cast(1000000 + coalesce(cast(max(right(paramedic_ID,5)) as int),0) + 1 as varchar(10)),5)),
					$1, $2, $3, true, $4, $5
				from paramedics returning paramedic_id;`
	var insertedID string
	err = r.db.QueryRow(query, data.FirstName, data.LastName, data.Email, data.UserCreate, time.Now()).Scan(&insertedID)
	if err != nil {
		return
	}

	fmt.Println(insertedID)

	result = &ParamedicProto{ // Initialize the result variable
		Paramedicid: insertedID,
		Email:       data.Email,
		Firstname:   data.FirstName,
		Lastname:    data.LastName,
	}

	return
}

func (r paramedicRepo) GetParamedicByHospitalID(hospitalid string) (data *ListparamedicProto, err error) {

	dataProto := &ListparamedicProto{
		Paramedics: make([]*ParamedicProto, 0), // Initialize the slice
	}

	var paramedics []paramedic

	query := `select b.paramedic_id, b.first_name, b.last_name,  b.email
			  from paramedicschedules a 
			  left join paramedics b
					on b.paramedic_id = a.paramedic_id
			  where healthcare_id=$1`
	fmt.Println("tampung data list")
	fmt.Println(query)

	err = r.db.Select(&paramedics, query, hospitalid)

	for _, paramedic := range paramedics {
		fmt.Printf("ID: %v", paramedic.FirstName)

		result := &ParamedicProto{ // Initialize the result variable
			Paramedicid: paramedic.ParamedicID,
			Email:       paramedic.Email,
			Firstname:   paramedic.FirstName,
			Lastname:    paramedic.LastName,
		}
		fmt.Println("insert data")
		dataProto.Paramedics = append(dataProto.Paramedics, result)
		fmt.Println("data inserted")
	}

	if err != nil {
		fmt.Println("errror oiii")
		return nil, err
	}

	return dataProto, err
}
