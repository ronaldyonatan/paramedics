package paramedics

import (
	"context"
	"fmt"
)

type Repository interface {
	writeParamedicRepo
	readParamedicRepo
}

type service struct {
	repo Repository
}

func NewService(repo Repository) service {
	return service{
		repo: repo,
	}
}

type writeParamedicRepo interface {
	InsertParamedic(data paramedic) (result *ParamedicProto, err error)
}

type readParamedicRepo interface {
	GetParamedicByHospitalID(hospitalid string) (data *ListparamedicProto, err error)
}

func (s service) CreateParamedic(ctx context.Context, req paramedicCreateRequest) (data *ParamedicProto, err error) {
	fmt.Println("create paramedic")

	tampung := req.ConvertToParamedic()

	fmt.Println("Convert to paramedic from request")
	fmt.Println(tampung)

	result, err := s.repo.InsertParamedic(req.ConvertToParamedic())

	data = result

	//fmt.Println(result)

	// data = &ParamedicProto{ // Initialize the result variable
	// 	Paramedicid: result.Paramedicid,
	// 	Email:       result.Email,
	// 	Firstname:   result.Firstname,
	// 	Lastname:    result.Lastname,
	// }

	// fmt.Println(data)

	if err != nil {
		fmt.Println("error insert")
		return nil, err
	}
	if data.Paramedicid == "" {
		fmt.Println("gagal insert")
		return nil, err
	}
	return data, nil
}

func (s service) FindByHospital(ctx context.Context, req FindByHospitalRequest) (data *ListparamedicProto, err error) {
	data, err = s.repo.GetParamedicByHospitalID(req.HospitalId)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, err
	}

	return data, nil
}
