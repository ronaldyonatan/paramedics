package paramedics

import (
	context "context"
)

type paramedicHandlerGrpc struct {
	svc service
}

func NewParamedicHandlerGPRC(svc service) paramedicHandlerGrpc {
	return paramedicHandlerGrpc{
		svc: svc,
	}
}

func (a paramedicHandlerGrpc) CreateParamedic(ctx context.Context, req *ParamedicCreateProto) (data *ParamedicProto, err error) {
	reqPayload := paramedicCreateRequest{
		Email:      req.Email,
		FirstName:  req.Firstname,
		LastName:   req.Lastname,
		UserCreate: req.Usercreate,
	}

	data, err = a.svc.CreateParamedic(ctx, reqPayload)

	return
}

func (a paramedicHandlerGrpc) FindByHospital(ctx context.Context, req *ParamedicFindByHospitalProto) (data *ListparamedicProto, err error) {
	reqPayload := FindByHospitalRequest{
		HospitalId: req.Hospitalid,
	}

	result, err := a.svc.FindByHospital(ctx, reqPayload)

	data = result
	// for _, v := range result {
	// 	data.Paramedics = append(data.Paramedics, &v)

	// }

	// emp = new(emptypb.Empty)
	return data, err
}

func NewParamedichHandlerGPRC(svc service) paramedicHandlerGrpc {
	return paramedicHandlerGrpc{
		svc: svc,
	}
}

func (a paramedicHandlerGrpc) mustEmbedUnimplementedParamedicServer() {

}
