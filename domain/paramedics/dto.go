package paramedics

type paramedicCreateRequest struct {
	FirstName  string `json:"first_name" validate:"required"`
	LastName   string `json:"last_name" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	UserCreate string `json:"user_create" validate:"required"`
}

type FindByHospitalRequest struct {
	HospitalId string `json:"hospitalid" validate:"required"`
}

func (a paramedicCreateRequest) ConvertToParamedic() paramedic {
	return paramedic{
		Email:      a.Email,
		FirstName:  a.FirstName,
		LastName:   a.LastName,
		UserCreate: a.UserCreate,
		// CreatedAt:  time.Now(),
		// IsActive:   true,
	}
}
