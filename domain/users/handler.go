package users

import (
	"errors"

	"github.com/fernandojec/assignment-2/config"
	"github.com/gofiber/fiber/v2"
)

type service interface {
	createAuthService
	activateAuthService
	signInService
}

type createAuthService interface {
	CreateAuth(req authCreateRequest, baseVerifyEmail string) (err error)
}
type activateAuthService interface {
	ActivateAuth(token string) (err error)
}
type signInService interface {
	SignInAuth(req authSignInRequest) (data authSignInResponse, err error)
}
type handler struct {
	svcAuth service
}

func NewHandler(svcauth service) handler {
	return handler{svcAuth: svcauth}
}

func (h *handler) CreateAuth(c *fiber.Ctx) error {

	req := new(authCreateRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	if err := ValidateStruct(req); err != nil {
		// return writeErrorResponse(c, echo.ErrBadRequest.Code, err.Error())
		return c.JSON(err)
	}

	err := h.svcAuth.CreateAuth(
		*req,
		config.AppConfig.App.BaseUrl+config.AppConfig.App.BasePort+"/v1/auth/verify-email/",
	)

	if err != nil {
		// return writeErrorResponse(c, echo.ErrInternalServerError.Code, err.Error())
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(map[string]interface{}{
		"status":  fiber.StatusCreated,
		"message": "Check your email to activate your account.",
	})
}

func (h *handler) ActivateAuth(c *fiber.Ctx) error {
	token := c.Params("token")
	if token == "" {
		return errors.New("token is required")
	}

	err := h.svcAuth.ActivateAuth(token)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(map[string]interface{}{
		"status":  fiber.StatusCreated,
		"message": "Your account has been activated",
	})
}
func (h *handler) SignInAuth(c *fiber.Ctx) error {
	req := new(authSignInRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}
	if err := ValidateStruct(req); err != nil {
		// return writeErrorResponse(c, echo.ErrBadRequest.Code, err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	resp, err := h.svcAuth.SignInAuth(*req)

	if err != nil {
		// return writeErrorResponse(c, echo.ErrInternalServerError.Code, err.Error())
		return err
	}

	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"status": fiber.StatusOK,
		"data":   resp,
	})
}
