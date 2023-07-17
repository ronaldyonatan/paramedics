package users

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/fernandojec/assignment-2/pkg/smtp"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	repo repo
}

type repo interface {
	writeAuthRepo
	readAuthRepo
	writeVerifyAuthRepo
	readVerifyAuthRepo
}

type writeAuthRepo interface {
	InsertAuth(data auth) (id uint, err error)
	UpdateAuthIsActive(data auth) (err error)
}

type readAuthRepo interface {
	GetAuthByEmail(email string) (data auth, err error)
	GetAuthByID(id uint) (data auth, err error)
}

type writeVerifyAuthRepo interface {
	InsertVerifyAuth(data verifyAuth) (err error)
	UpdateVerifyAuth(data verifyAuth) (err error)
}

type readVerifyAuthRepo interface {
	GetVerifyAuthByToken(token string) (data verifyAuth, err error)
}

func NewService(repo repo) authService {
	return authService{repo: repo}
}

func (s authService) CreateAuth(req authCreateRequest, baseVerifyEmail string) (err error) {
	dataAuth, err := s.repo.GetAuthByEmail(req.Email)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if err == nil && dataAuth.Id != 0 {
		return errors.New("this email has been registered")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(pass)
	id, err := s.repo.InsertAuth(req.ConvertToAuth())
	if err != nil {
		return err
	}
	dataVerify := verifyAuth{
		AuthId:    id,
		Token:     uuid.New().String(),
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().Add(10 * time.Minute),
	}
	err = s.repo.InsertVerifyAuth(dataVerify)
	if err != nil {
		return err
	}

	err = smtp.SendMail(smtp.Mail{
		From:    "My APP Registration Confirmation",
		To:      []string{req.Email},
		Subject: "Registration Activation",
		Body: fmt.Sprintf(
			`<a href='%s%s'>Click here to activate your account</a>`,
			baseVerifyEmail,
			dataVerify.Token,
		),
	})

	return
}

func (s authService) ActivateAuth(token string) (err error) {
	verifyAuth, err := s.repo.GetVerifyAuthByToken(token)
	if err != nil {
		return err
	}
	if verifyAuth.ActivatedAt.Valid {
		return //errors.New("you account has been activated")
	}
	if verifyAuth.ExpiredAt.Before(time.Now()) {
		return errors.New("your activation link has been expired. please send new activation link")
	}
	verifyAuth.ActivatedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	err = s.repo.UpdateVerifyAuth(verifyAuth)
	if err != nil {
		return err
	}
	dataAuth := auth{
		Id:          verifyAuth.AuthId,
		IsActive:    true,
		ActivatedAt: verifyAuth.ActivatedAt,
	}
	err = s.repo.UpdateAuthIsActive(dataAuth)
	if err != nil {
		return err
	}
	return
}

func (s authService) SignInAuth(req authSignInRequest) (data authSignInResponse, err error) {
	// reqauth := req.ConvertToAuth()
	dataAuth, err := s.repo.GetAuthByEmail(req.Email)
	if err != nil && err == sql.ErrNoRows {
		return authSignInResponse{}, errors.New("user not found")
	}
	if err != nil {
		return authSignInResponse{}, err
	}
	if dataAuth.Id == 0 {
		return authSignInResponse{}, errors.New("user not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(dataAuth.Password), []byte(req.Password)); err != nil {
		return authSignInResponse{}, errors.New("invalid password")
	}
	if !dataAuth.ActivatedAt.Valid {
		return authSignInResponse{}, errors.New("user is not activated")
	}
	data = dataAuth.ConvertToAuthSignInResponse()
	return
}

func (s authService) SendNewActivationLink(token string, baseVerifyEmail string) (err error) {
	verifyAuthData, err := s.repo.GetVerifyAuthByToken(token)
	if err != nil {
		return err
	}

	dataAuth, err := s.repo.GetAuthByID(verifyAuthData.AuthId)

	if err != nil {
		return
	}

	dataVerify := verifyAuth{
		AuthId:    verifyAuthData.AuthId,
		Token:     uuid.New().String(),
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().Add(10 * time.Minute),
	}
	err = s.repo.InsertVerifyAuth(dataVerify)
	if err != nil {
		return err
	}

	err = smtp.SendMail(smtp.Mail{
		From:    "My APP Registration Confirmation",
		To:      []string{dataAuth.Email},
		Subject: "Registration Activation",
		Body: fmt.Sprintf(
			`<a href='%s%s'>Click here to activate your account</a>`,
			baseVerifyEmail,
			dataVerify.Token,
		),
	})

	return
}
