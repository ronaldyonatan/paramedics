package users

import "github.com/gofiber/fiber/v2"

type serviceWeb interface {
	createAuthService
	activateAuthService
	signInService
}

type createAuthServiceWeb interface {
	CreateAuth(req authCreateRequest) (err error)
}
type activateAuthServiceWeb interface {
	ActivateAuth(token string) (err error)
}
type signInServiceWeb interface {
	SignInAuth(req authSignInRequest) (data authSignInResponse, err error)
}
type handlerWeb struct {
	svcAuth serviceWeb
}

func NewHandlerWeb(svcauth serviceWeb) handlerWeb {
	return handlerWeb{svcAuth: svcauth}
}

func (h *handlerWeb) Home(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Home",
	})
}

func (h *handlerWeb) SignUp(c *fiber.Ctx) error {
	return c.Render("sign-up", fiber.Map{
		"Title": "Sign Up",
	})
}

func (h *handlerWeb) SignIn(c *fiber.Ctx) error {
	return c.Render("sign-in", fiber.Map{
		"Title": "Sign In",
	})
}
